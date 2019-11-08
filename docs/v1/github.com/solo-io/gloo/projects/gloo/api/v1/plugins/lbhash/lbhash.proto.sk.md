
---
title: "lbhash.proto"
weight: 5
---

<!-- Code generated by solo-kit. DO NOT EDIT. -->


### Package: `lbhash.plugins.gloo.solo.io` 
#### Types:


- [RouteActionHashConfig](#routeactionhashconfig)
- [Cookie](#cookie)
- [HashPolicy](#hashpolicy)
  



##### Source File: [github.com/solo-io/gloo/projects/gloo/api/v1/plugins/lbhash/lbhash.proto](https://github.com/solo-io/gloo/blob/master/projects/gloo/api/v1/plugins/lbhash/lbhash.proto)





---
### RouteActionHashConfig

 
Specifies the route’s hashing policy if the upstream cluster uses a hashing load balancer.
https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/route/route.proto#envoy-api-msg-route-routeaction-hashpolicy

```yaml
"hashPolicies": []lbhash.plugins.gloo.solo.io.HashPolicy

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `hashPolicies` | [[]lbhash.plugins.gloo.solo.io.HashPolicy](../lbhash.proto.sk/#hashpolicy) | The list of policies Envoy will use when generating a hash key for a hashing load balancer. |  |




---
### Cookie

 
Envoy supports two types of cookie affinity:
- Passive: Envoy reads the cookie from the headers
- Generated: Envoy uses the cookie spec to generate a cookie
In either case, the cookie is incorporated in the hash key.
additional notes https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/route/route.proto#envoy-api-msg-route-routeaction-hashpolicy-cookie

```yaml
"name": string
"ttl": .google.protobuf.Duration
"path": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `name` | `string` | required, the name of the cookie to be used to obtain the hash key. |  |
| `ttl` | [.google.protobuf.Duration](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/duration) | If specified, a cookie with the TTL will be generated if the cookie is not present. If the TTL is present and zero, the generated cookie will be a session cookie. |  |
| `path` | `string` | The name of the path for the cookie. If no path is specified here, no path will be set for the cookie. |  |




---
### HashPolicy

 
Specifies an element of Envoy's hashing policy for hashing load balancers

```yaml
"header": string
"cookie": .lbhash.plugins.gloo.solo.io.Cookie
"sourceIp": bool
"terminal": bool

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| `header` | `string` | Use a given header's value as a component of the hashing load balancer's hash key. Only one of `header`, or `sourceIp` can be set. |  |
| `cookie` | [.lbhash.plugins.gloo.solo.io.Cookie](../lbhash.proto.sk/#cookie) | Use a given cookie as a component of the hashing load balancer's hash key. Only one of `cookie`, or `sourceIp` can be set. |  |
| `sourceIp` | `bool` | Use the request's source IP address as a component of the hashing load balancer's hash key. Only one of `sourceIp`, or `cookie` can be set. |  |
| `terminal` | `bool` | If set, and a hash key is available after evaluating this policy, Envoy will skip the subsequent policies and use the key as it is. This is useful for defining "fallback" policies and limiting the time Envoy spends generating hash keys. |  |





<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->