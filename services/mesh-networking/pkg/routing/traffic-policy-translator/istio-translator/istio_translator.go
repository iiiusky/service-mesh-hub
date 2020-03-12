package istio_translator

import (
	"context"
	"sort"

	"github.com/rotisserie/eris"
	core_types "github.com/solo-io/mesh-projects/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	discovery_v1alpha1 "github.com/solo-io/mesh-projects/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	networking_v1alpha1 "github.com/solo-io/mesh-projects/pkg/api/networking.zephyr.solo.io/v1alpha1"
	networking_types "github.com/solo-io/mesh-projects/pkg/api/networking.zephyr.solo.io/v1alpha1/types"
	"github.com/solo-io/mesh-projects/pkg/clients"
	istio_networking "github.com/solo-io/mesh-projects/pkg/clients/istio/networking"
	zephyr_discovery "github.com/solo-io/mesh-projects/pkg/clients/zephyr/discovery"
	"github.com/solo-io/mesh-projects/services/common"
	mc_manager "github.com/solo-io/mesh-projects/services/common/multicluster/manager"
	traffic_policy_translator "github.com/solo-io/mesh-projects/services/mesh-networking/pkg/routing/traffic-policy-translator"
	"github.com/solo-io/mesh-projects/services/mesh-networking/pkg/routing/traffic-policy-translator/preprocess"
	api_v1alpha3 "istio.io/api/networking/v1alpha3"
	client_v1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	"k8s.io/apimachinery/pkg/labels"
)

const (
	TranslatorId = "istio-translator"
)

var (
	ClientNotFoundErr = func(clusterName string) error {
		return eris.Errorf("Client not found for cluster with name: %s", clusterName)
	}
)

type IstioTranslator traffic_policy_translator.TrafficPolicyMeshTranslator

func NewIstioTrafficPolicyTranslator(
	dynamicClientGetter mc_manager.DynamicClientGetter,
	meshClient zephyr_discovery.MeshClient,
	meshServiceClient zephyr_discovery.MeshServiceClient,
	meshServiceSelector preprocess.MeshServiceSelector,
	virtualServiceClientFactory istio_networking.VirtualServiceClientFactory,
	destinationRuleClientFactory istio_networking.DestinationRuleClientFactory,
) IstioTranslator {
	return &istioTrafficPolicyTranslator{
		dynamicClientGetter:          dynamicClientGetter,
		meshClient:                   meshClient,
		meshServiceClient:            meshServiceClient,
		meshServiceSelector:          meshServiceSelector,
		virtualServiceClientFactory:  virtualServiceClientFactory,
		destinationRuleClientFactory: destinationRuleClientFactory,
	}
}

type istioTrafficPolicyTranslator struct {
	dynamicClientGetter          mc_manager.DynamicClientGetter
	meshClient                   zephyr_discovery.MeshClient
	meshServiceClient            zephyr_discovery.MeshServiceClient
	virtualServiceClientFactory  istio_networking.VirtualServiceClientFactory
	destinationRuleClientFactory istio_networking.DestinationRuleClientFactory
	meshServiceSelector          preprocess.MeshServiceSelector
}

/*
	Translate a TrafficPolicy into the following Istio specific configuration:
	https://istio.io/docs/concepts/traffic-management/

	1. VirtualService - routing rules (e.g. retries, fault injection, traffic shifts)
	2. DestinationRule - post-routing rules (e.g. subset declaration)
*/
func (i *istioTrafficPolicyTranslator) TranslateTrafficPolicy(
	ctx context.Context,
	meshService *discovery_v1alpha1.MeshService,
	mesh *discovery_v1alpha1.Mesh,
	mergedTrafficPolicies []*networking_v1alpha1.TrafficPolicy,
) *networking_types.TrafficPolicyStatus_TranslatorError {
	if mesh.Spec.GetIstio() == nil {
		return nil
	}
	destinationRuleClient, virtualServiceClient, err := i.fetchClientsForMeshService(ctx, meshService)
	if err != nil {
		return i.errorToStatus(err)
	}
	translatorError := i.ensureDestinationRule(ctx, meshService, destinationRuleClient)
	if translatorError != nil {
		return translatorError
	}
	translatorError = i.ensureVirtualService(ctx, meshService, mergedTrafficPolicies, virtualServiceClient)
	if translatorError != nil {
		return translatorError
	}
	return nil
}

// get DestinationRule and VirtualService clients for MeshService's cluster
func (i *istioTrafficPolicyTranslator) fetchClientsForMeshService(
	ctx context.Context,
	meshService *discovery_v1alpha1.MeshService,
) (istio_networking.DestinationRuleClient, istio_networking.VirtualServiceClient, error) {
	clusterName, err := i.getClusterNameForMeshService(ctx, meshService)
	if err != nil {
		return nil, nil, err
	}
	dynamicClient, ok := i.dynamicClientGetter.GetClientForCluster(clusterName)
	if !ok {
		return nil, nil, ClientNotFoundErr(clusterName)
	}
	return i.destinationRuleClientFactory(dynamicClient), i.virtualServiceClientFactory(dynamicClient), nil
}

func (i *istioTrafficPolicyTranslator) ensureDestinationRule(
	ctx context.Context,
	meshService *discovery_v1alpha1.MeshService,
	destinationRuleClient istio_networking.DestinationRuleClient,
) *networking_types.TrafficPolicyStatus_TranslatorError {
	destinationRule := &client_v1alpha3.DestinationRule{
		ObjectMeta: clients.ResourceRefToObjectMeta(meshService.Spec.GetKubeService().GetRef()),
		Spec: api_v1alpha3.DestinationRule{
			Host: buildServiceHostname(meshService),
			TrafficPolicy: &api_v1alpha3.TrafficPolicy{
				Tls: &api_v1alpha3.TLSSettings{
					Mode: api_v1alpha3.TLSSettings_ISTIO_MUTUAL,
				},
			},
		},
	}
	err := destinationRuleClient.Upsert(ctx, destinationRule)
	if err != nil {
		return i.errorToStatus(err)
	}
	return nil
}

func (i *istioTrafficPolicyTranslator) ensureVirtualService(
	ctx context.Context,
	meshService *discovery_v1alpha1.MeshService,
	mergedTrafficPolicies []*networking_v1alpha1.TrafficPolicy,
	virtualServiceClient istio_networking.VirtualServiceClient,
) *networking_types.TrafficPolicyStatus_TranslatorError {
	computedVirtualService, err := i.translateIntoVirtualService(ctx, meshService, mergedTrafficPolicies)
	if err != nil {
		return i.errorToStatus(err)
	}
	// Upsert computed VirtualService
	err = virtualServiceClient.Upsert(ctx, computedVirtualService)
	if err != nil {
		return i.errorToStatus(err)
	}
	return nil
}

func (i *istioTrafficPolicyTranslator) translateIntoVirtualService(
	ctx context.Context,
	meshService *discovery_v1alpha1.MeshService,
	trafficPolicies []*networking_v1alpha1.TrafficPolicy,
) (*client_v1alpha3.VirtualService, error) {
	virtualService := &client_v1alpha3.VirtualService{
		ObjectMeta: clients.ResourceRefToObjectMeta(meshService.Spec.GetKubeService().GetRef()),
		Spec: api_v1alpha3.VirtualService{
			Hosts: []string{buildServiceHostname(meshService)},
			Http:  []*api_v1alpha3.HTTPRoute{},
		},
	}
	for _, trafficPolicy := range trafficPolicies {
		fault, err := i.translateFaultInjection(trafficPolicy)
		if err != nil {
			return nil, err
		}
		corsPolicy, err := i.translateCorsPolicy(trafficPolicy)
		if err != nil {
			return nil, err
		}
		requestMatchers, err := i.translateRequestMatchers(trafficPolicy)
		if err != nil {
			return nil, err
		}
		var mirrorPercentage *api_v1alpha3.Percent
		if trafficPolicy.Spec.GetMirror() != nil {
			mirrorPercentage = &api_v1alpha3.Percent{Value: trafficPolicy.Spec.GetMirror().GetPercentage()}
		}
		mirror, err := i.translateMirror(ctx, meshService, trafficPolicy)
		if err != nil {
			return nil, err
		}
		trafficShift, err := i.translateTrafficShift(ctx, meshService, trafficPolicy)
		if err != nil {
			return nil, err
		}
		virtualService.Spec.Http = append(virtualService.Spec.GetHttp(), &api_v1alpha3.HTTPRoute{
			Match:            requestMatchers,
			Route:            trafficShift,
			Timeout:          trafficPolicy.Spec.GetRequestTimeout(),
			Fault:            fault,
			CorsPolicy:       corsPolicy,
			Retries:          i.translateRetries(trafficPolicy),
			MirrorPercentage: mirrorPercentage,
			Mirror:           mirror,
			Headers:          i.translateHeaderManipulation(trafficPolicy),
		})
	}
	return virtualService, nil
}

func (i *istioTrafficPolicyTranslator) translateRequestMatchers(
	trafficPolicy *networking_v1alpha1.TrafficPolicy,
) ([]*api_v1alpha3.HTTPMatchRequest, error) {
	var translatedRequestMatcher []*api_v1alpha3.HTTPMatchRequest
	matchers := trafficPolicy.Spec.GetHttpRequestMatchers()
	if matchers != nil {
		translatedRequestMatcher = []*api_v1alpha3.HTTPMatchRequest{}
		for _, matcher := range matchers {
			headerMatchers, inverseHeaderMatchers := i.translateRequestMatcherHeaders(matcher.GetHeaders())
			uriMatcher, err := i.translateRequestMatcherPathSpecifier(matcher)
			if err != nil {
				return nil, err
			}
			matchRequest := &api_v1alpha3.HTTPMatchRequest{
				Uri:            uriMatcher,
				Method:         &api_v1alpha3.StringMatch{MatchType: &api_v1alpha3.StringMatch_Exact{Exact: matcher.GetMethod().String()}},
				QueryParams:    i.translateRequestMatcherQueryParams(matcher.GetQueryParameters()),
				Headers:        headerMatchers,
				WithoutHeaders: inverseHeaderMatchers,
				SourceLabels:   trafficPolicy.Spec.GetSourceSelector().GetLabels(),
			}

			if len(trafficPolicy.Spec.GetSourceSelector().GetNamespaces()) > 0 {
				for _, namespace := range trafficPolicy.Spec.GetSourceSelector().GetNamespaces() {
					matchRequestWithNamespace := *matchRequest
					matchRequestWithNamespace.SourceNamespace = namespace
					translatedRequestMatcher = append(translatedRequestMatcher, &matchRequestWithNamespace)
				}
			} else {
				translatedRequestMatcher = append(translatedRequestMatcher, matchRequest)
			}

		}
	}
	return translatedRequestMatcher, nil
}

func (i *istioTrafficPolicyTranslator) translateRequestMatcherPathSpecifier(matcher *networking_types.HttpMatcher) (*api_v1alpha3.StringMatch, error) {
	if matcher != nil && matcher.GetPathSpecifier() != nil {
		switch pathSpecifierType := matcher.GetPathSpecifier().(type) {
		case *networking_types.HttpMatcher_Exact:
			return &api_v1alpha3.StringMatch{MatchType: &api_v1alpha3.StringMatch_Exact{Exact: matcher.GetExact()}}, nil
		case *networking_types.HttpMatcher_Prefix:
			return &api_v1alpha3.StringMatch{MatchType: &api_v1alpha3.StringMatch_Prefix{Prefix: matcher.GetPrefix()}}, nil
		case *networking_types.HttpMatcher_Regex:
			return &api_v1alpha3.StringMatch{MatchType: &api_v1alpha3.StringMatch_Regex{Regex: matcher.GetRegex()}}, nil
		default:
			return nil, eris.Errorf("RequestMatchers[].PathSpecifier has unexpected type %T", pathSpecifierType)
		}
	}
	return nil, nil
}

func (i *istioTrafficPolicyTranslator) translateRequestMatcherQueryParams(matchers []*networking_types.QueryParameterMatcher) map[string]*api_v1alpha3.StringMatch {
	var translatedQueryParamMatcher map[string]*api_v1alpha3.StringMatch
	if matchers != nil {
		translatedQueryParamMatcher = map[string]*api_v1alpha3.StringMatch{}
		for _, matcher := range matchers {
			if matcher.GetRegex() {
				translatedQueryParamMatcher[matcher.GetName()] = &api_v1alpha3.StringMatch{
					MatchType: &api_v1alpha3.StringMatch_Regex{Regex: matcher.GetValue()},
				}
			} else {
				translatedQueryParamMatcher[matcher.GetName()] = &api_v1alpha3.StringMatch{
					MatchType: &api_v1alpha3.StringMatch_Exact{Exact: matcher.GetValue()},
				}
			}
		}
	}
	return translatedQueryParamMatcher
}

func (i *istioTrafficPolicyTranslator) translateRequestMatcherHeaders(matchers []*networking_types.HeaderMatcher) (
	map[string]*api_v1alpha3.StringMatch, map[string]*api_v1alpha3.StringMatch,
) {
	headerMatchers := map[string]*api_v1alpha3.StringMatch{}
	inverseHeaderMatchers := map[string]*api_v1alpha3.StringMatch{}
	var matcherMap map[string]*api_v1alpha3.StringMatch
	if matchers != nil {
		for _, matcher := range matchers {
			matcherMap = headerMatchers
			if matcher.GetInvertMatch() {
				matcherMap = inverseHeaderMatchers
			}
			if matcher.GetRegex() {
				matcherMap[matcher.GetName()] = &api_v1alpha3.StringMatch{
					MatchType: &api_v1alpha3.StringMatch_Regex{Regex: matcher.GetValue()},
				}
			} else {
				matcherMap[matcher.GetName()] = &api_v1alpha3.StringMatch{
					MatchType: &api_v1alpha3.StringMatch_Exact{Exact: matcher.GetValue()},
				}
			}
		}
	}
	// ensure field is set to nil if empty
	if len(headerMatchers) == 0 {
		headerMatchers = nil
	}
	if len(inverseHeaderMatchers) == 0 {
		inverseHeaderMatchers = nil
	}
	return headerMatchers, inverseHeaderMatchers
}

// ensure that subsets declared in this TrafficPolicy are reflected in the relevant kube Service's DestinationRules
// return name of Subset declared in DestinationRule
func (i *istioTrafficPolicyTranslator) translateSubset(
	ctx context.Context,
	destination *networking_types.MultiDestination_WeightedDestination,
) (string, error) {
	// fetch client for destination's cluster
	clusterName := destination.GetDestination().GetCluster().GetValue()
	dynamicClient, ok := i.dynamicClientGetter.GetClientForCluster(clusterName)
	if !ok {
		return "", ClientNotFoundErr(clusterName)
	}
	destinationRuleClient := i.destinationRuleClientFactory(dynamicClient)
	destinationRule, err := destinationRuleClient.Get(ctx, clients.ResourceRefToObjectKey(destination.GetDestination()))
	if err != nil {
		return "", err
	}
	for _, subset := range destinationRule.Spec.GetSubsets() {
		if labels.Equals(subset.GetLabels(), destination.GetSubset()) {
			return subset.GetName(), nil
		}
	}
	// subset doesn't yet exist, update the DestinationRule with it and return its generated name
	subsetName := generateUniqueSubsetName(destination.GetSubset())
	destinationRule.Spec.Subsets = append(destinationRule.Spec.Subsets, &api_v1alpha3.Subset{
		Name:   subsetName,
		Labels: destination.GetSubset(),
	})
	err = destinationRuleClient.Update(ctx, destinationRule)
	if err != nil {
		return "", err
	}
	return subsetName, nil
}

// For each Destination, create an Istio HTTPRouteDestination
func (i *istioTrafficPolicyTranslator) translateTrafficShift(
	ctx context.Context,
	meshService *discovery_v1alpha1.MeshService,
	trafficPolicy *networking_v1alpha1.TrafficPolicy,
) ([]*api_v1alpha3.HTTPRouteDestination, error) {
	var translatedTrafficShift []*api_v1alpha3.HTTPRouteDestination
	trafficShift := trafficPolicy.Spec.GetTrafficShift()
	if trafficShift != nil {
		translatedTrafficShift = []*api_v1alpha3.HTTPRouteDestination{}
		for _, destination := range trafficShift.GetDestinations() {
			hostnameForKubeService, err := i.getHostnameForKubeService(ctx, meshService, destination.GetDestination())
			if err != nil {
				return nil, err
			}
			httpRouteDestination := &api_v1alpha3.HTTPRouteDestination{
				Destination: &api_v1alpha3.Destination{
					Host: hostnameForKubeService,
				},
				Weight: int32(destination.GetWeight()),
			}
			if destination.Subset != nil {
				subsetName, err := i.translateSubset(ctx, destination)
				if err != nil {
					return nil, err
				}
				httpRouteDestination.Destination.Subset = subsetName
			}
			translatedTrafficShift = append(translatedTrafficShift, httpRouteDestination)
		}
	}
	return translatedTrafficShift, nil
}

func (i *istioTrafficPolicyTranslator) translateRetries(trafficPolicy *networking_v1alpha1.TrafficPolicy) *api_v1alpha3.HTTPRetry {
	var translatedRetries *api_v1alpha3.HTTPRetry
	retries := trafficPolicy.Spec.GetRetries()
	if retries != nil {
		translatedRetries = &api_v1alpha3.HTTPRetry{
			Attempts:      retries.GetAttempts(),
			PerTryTimeout: retries.GetPerTryTimeout(),
		}
	}
	return translatedRetries
}

func (i *istioTrafficPolicyTranslator) translateFaultInjection(trafficPolicy *networking_v1alpha1.TrafficPolicy) (*api_v1alpha3.HTTPFaultInjection, error) {
	var translatedFaultInjection *api_v1alpha3.HTTPFaultInjection
	faultInjection := trafficPolicy.Spec.GetFaultInjection()
	if faultInjection != nil {
		switch injectionType := faultInjection.GetFaultInjectionType().(type) {
		case *networking_types.FaultInjection_Abort_:
			abort := faultInjection.GetAbort()
			switch abortType := abort.GetErrorType().(type) {
			case *networking_types.FaultInjection_Abort_HttpStatus:
				translatedFaultInjection = &api_v1alpha3.HTTPFaultInjection{
					Abort: &api_v1alpha3.HTTPFaultInjection_Abort{
						ErrorType:  &api_v1alpha3.HTTPFaultInjection_Abort_HttpStatus{HttpStatus: abort.GetHttpStatus()},
						Percentage: &api_v1alpha3.Percent{Value: faultInjection.GetPercentage()},
					}}
			default:
				return nil, eris.Errorf("Abort.ErrorType has unexpected type %T", abortType)
			}
		case *networking_types.FaultInjection_Delay_:
			delay := faultInjection.GetDelay()
			switch delayType := delay.GetHttpDelayType().(type) {
			case *networking_types.FaultInjection_Delay_FixedDelay:
				translatedFaultInjection = &api_v1alpha3.HTTPFaultInjection{
					Delay: &api_v1alpha3.HTTPFaultInjection_Delay{
						HttpDelayType: &api_v1alpha3.HTTPFaultInjection_Delay_FixedDelay{FixedDelay: delay.GetFixedDelay()},
						Percentage:    &api_v1alpha3.Percent{Value: faultInjection.GetPercentage()},
					}}
			case *networking_types.FaultInjection_Delay_ExponentialDelay:
				translatedFaultInjection = &api_v1alpha3.HTTPFaultInjection{
					Delay: &api_v1alpha3.HTTPFaultInjection_Delay{
						HttpDelayType: &api_v1alpha3.HTTPFaultInjection_Delay_ExponentialDelay{ExponentialDelay: delay.GetExponentialDelay()},
						Percentage:    &api_v1alpha3.Percent{Value: faultInjection.GetPercentage()},
					}}
			default:
				return nil, eris.Errorf("Delay.HTTPDelayType has unexpected type %T", delayType)
			}
		default:
			return nil, eris.Errorf("FaultInjection.FaultInjectionType has unexpected type %T", injectionType)
		}
	}
	return translatedFaultInjection, nil
}

func (i *istioTrafficPolicyTranslator) translateMirror(
	ctx context.Context,
	meshService *discovery_v1alpha1.MeshService,
	trafficPolicy *networking_v1alpha1.TrafficPolicy,
) (*api_v1alpha3.Destination, error) {
	var mirror *api_v1alpha3.Destination
	if trafficPolicy.Spec.GetMirror() != nil {
		hostnameForKubeService, err := i.getHostnameForKubeService(ctx, meshService, trafficPolicy.Spec.GetMirror().GetDestination())
		if err != nil {
			return nil, err
		}
		mirror = &api_v1alpha3.Destination{
			Host: hostnameForKubeService,
		}
	}
	return mirror, nil
}

func (i *istioTrafficPolicyTranslator) translateHeaderManipulation(trafficPolicy *networking_v1alpha1.TrafficPolicy) *api_v1alpha3.Headers {
	var translatedHeaderManipulation *api_v1alpha3.Headers
	headerManipulation := trafficPolicy.Spec.GetHeaderManipulation()
	if headerManipulation != nil {
		translatedHeaderManipulation = &api_v1alpha3.Headers{
			Request: &api_v1alpha3.Headers_HeaderOperations{
				Add:    headerManipulation.GetAppendRequestHeaders(),
				Remove: headerManipulation.GetRemoveRequestHeaders(),
			},
			Response: &api_v1alpha3.Headers_HeaderOperations{
				Add:    headerManipulation.GetAppendResponseHeaders(),
				Remove: headerManipulation.GetRemoveResponseHeaders(),
			},
		}
	}
	return translatedHeaderManipulation
}

func (i *istioTrafficPolicyTranslator) translateCorsPolicy(trafficPolicy *networking_v1alpha1.TrafficPolicy) (*api_v1alpha3.CorsPolicy, error) {
	var translatedCorsPolicy *api_v1alpha3.CorsPolicy
	corsPolicy := trafficPolicy.Spec.GetCorsPolicy()
	if corsPolicy != nil {
		var allowOrigins []*api_v1alpha3.StringMatch
		for i, allowOrigin := range corsPolicy.GetAllowOrigins() {
			var stringMatch *api_v1alpha3.StringMatch
			switch matchType := allowOrigin.GetMatchType().(type) {
			case *networking_types.StringMatch_Exact:
				stringMatch = &api_v1alpha3.StringMatch{MatchType: &api_v1alpha3.StringMatch_Exact{Exact: allowOrigin.GetExact()}}
			case *networking_types.StringMatch_Prefix:
				stringMatch = &api_v1alpha3.StringMatch{MatchType: &api_v1alpha3.StringMatch_Prefix{Prefix: allowOrigin.GetPrefix()}}
			case *networking_types.StringMatch_Regex:
				stringMatch = &api_v1alpha3.StringMatch{MatchType: &api_v1alpha3.StringMatch_Regex{Regex: allowOrigin.GetRegex()}}
			default:
				return nil, eris.Errorf("AllowOrigins[%d].MatchType has unexpected type %T", i, matchType)
			}
			allowOrigins = append(allowOrigins, stringMatch)
		}
		translatedCorsPolicy = &api_v1alpha3.CorsPolicy{
			AllowOrigins:     allowOrigins,
			AllowMethods:     corsPolicy.GetAllowMethods(),
			AllowHeaders:     corsPolicy.GetAllowHeaders(),
			ExposeHeaders:    corsPolicy.GetExposeHeaders(),
			MaxAge:           corsPolicy.GetMaxAge(),
			AllowCredentials: corsPolicy.GetAllowCredentials(),
		}
	}
	return translatedCorsPolicy, nil
}

func (i *istioTrafficPolicyTranslator) errorToStatus(err error) *networking_types.TrafficPolicyStatus_TranslatorError {
	return &networking_types.TrafficPolicyStatus_TranslatorError{
		TranslatorId: TranslatorId,
		ErrorMessage: err.Error(),
	}
}

func (i *istioTrafficPolicyTranslator) getClusterNameForMeshService(
	ctx context.Context,
	meshService *discovery_v1alpha1.MeshService,
) (string, error) {
	mesh, err := i.meshClient.Get(ctx, clients.ResourceRefToObjectKey(meshService.Spec.GetMesh()))
	if err != nil {
		return "", err
	}
	return mesh.Spec.GetCluster().GetName(), nil
}

// If destination is in the same namespace as k8s Service, return k8s Service name.namespace
// Else, return k8s Service multicluster DNS name
func (i *istioTrafficPolicyTranslator) getHostnameForKubeService(
	ctx context.Context,
	meshService *discovery_v1alpha1.MeshService,
	destination *core_types.ResourceRef,
) (string, error) {
	destinationMeshService, err := i.meshServiceSelector.GetBackingMeshService(
		ctx, destination.GetName(), destination.GetNamespace(), destination.GetCluster().GetValue())
	if err != nil {
		return "", err
	}
	if common.AreResourcesOnLocalCluster(destination, meshService.Spec.GetKubeService().GetRef()) ||
		destination.GetCluster().GetValue() == meshService.Spec.GetKubeService().GetRef().GetCluster().GetValue() {
		// destination is on the same cluster as the MeshService's k8s Service
		return destinationMeshService.Spec.GetKubeService().GetRef().GetName() +
			"." + destinationMeshService.Spec.GetKubeService().GetRef().GetNamespace(), nil
	} else {
		// destination is on a remote cluster to the MeshService's k8s Service
		return destinationMeshService.Spec.GetFederation().GetMulticlusterDnsName(), nil
	}
}

func buildServiceHostname(meshService *discovery_v1alpha1.MeshService) string {
	return meshService.GetName()
}

// sort the label keys, then in order concatenate keys-values
func generateUniqueSubsetName(selectors map[string]string) string {
	subsetName := ""
	var keys []string
	for key := range selectors {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		subsetName += key + ":" + selectors[key] + ","
	}
	return subsetName
}
