package event_watcher_factories

import (
	core_controllers "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/controller"
	mc_manager "github.com/solo-io/service-mesh-hub/pkg/common/compute-target/k8s"
)

func NewPodEventWatcherFactory() PodEventWatcherFactory {
	return &defaultPodEventWatcherFactory{}
}

type defaultPodEventWatcherFactory struct{}

func (d *defaultPodEventWatcherFactory) Build(mgr mc_manager.AsyncManager, clusterName string) core_controllers.PodEventWatcher {
	return core_controllers.NewPodEventWatcher(clusterName, mgr.Manager())
}
