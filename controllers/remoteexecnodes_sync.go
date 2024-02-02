package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log"

	ytv1 "github.com/ytsaurus/yt-k8s-operator/api/v1"
	"github.com/ytsaurus/yt-k8s-operator/pkg/apiproxy"
	"github.com/ytsaurus/yt-k8s-operator/pkg/components"
	"github.com/ytsaurus/yt-k8s-operator/pkg/ytconfig"
)

func (r *RemoteExecNodesReconciler) Sync(
	ctx context.Context,
	resource *ytv1.RemoteExecNodes,
	remoteYtsaurus *ytv1.RemoteYtsaurus,
) (ctrl.Result, error) {
	logger := log.FromContext(ctx).WithValues("component", "remotedatanodes")
	apiProxy := apiproxy.NewAPIProxy(resource, r.Client, r.Recorder, r.Scheme)

	cfgen := ytconfig.NewRemoteNodeGenerator(
		types.NamespacedName{Name: resource.Name, Namespace: resource.Namespace},
		getClusterDomain(r.Client),
		resource.Spec.ConfigurationSpec,
		remoteYtsaurus.Spec.MasterConnectionSpec,
	)

	component := components.NewRemoteExecNodes(
		cfgen,
		resource,
		apiProxy,
		resource.Spec.ExecNodesSpec,
		resource.Spec.ConfigurationSpec,
	)
	err := component.Fetch(ctx)
	if err != nil {
		logger.Error(err, "failed to fetch remote nodes")
		return ctrl.Result{Requeue: true}, err
	}

	err = component.Sync(ctx)
	if err != nil {
		logger.Error(err, "failed to sync remote nodes")
		return ctrl.Result{Requeue: true}, err
	}
	return ctrl.Result{}, nil
}
