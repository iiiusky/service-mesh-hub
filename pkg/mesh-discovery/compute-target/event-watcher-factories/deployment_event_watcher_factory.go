package event_watcher_factories

import (
	apps_controllers "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/controller"
	mc_manager "github.com/solo-io/service-mesh-hub/pkg/common/compute-target/k8s"
)

func NewDeploymentEventWatcherFactory() DeploymentEventWatcherFactory {
	return &defaultDeploymentEventWatcherFactory{}
}

type defaultDeploymentEventWatcherFactory struct{}

func (d *defaultDeploymentEventWatcherFactory) Build(mgr mc_manager.AsyncManager, clusterName string) apps_controllers.DeploymentEventWatcher {
	// just directly return the generated autopilot implementation
	return apps_controllers.NewDeploymentEventWatcher(clusterName, mgr.Manager())
}
