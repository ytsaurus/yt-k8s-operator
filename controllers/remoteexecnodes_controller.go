/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	ytv1 "github.com/ytsaurus/yt-k8s-operator/api/v1"
)

// RemoteExecNodesReconciler reconciles a RemoteExecNodes object
type RemoteExecNodesReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

//+kubebuilder:rbac:groups=cluster.ytsaurus.tech,resources=remoteexecnodes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cluster.ytsaurus.tech,resources=remoteexecnodes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cluster.ytsaurus.tech,resources=remoteexecnodes/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=statefulset,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps,resources=statefulset/status,verbs=get
//+kubebuilder:rbac:groups=core,resources=pod,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pod/status,verbs=get

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the RemoteExecNodes object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *RemoteExecNodesReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var nodes ytv1.RemoteExecNodes
	if err := r.Get(ctx, req.NamespacedName, &nodes); err != nil {
		logger.Error(err, "unable to fetch remote nodes")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	var remoteYtsaurus ytv1.RemoteYtsaurus
	ytsaurusName := types.NamespacedName{Name: nodes.Spec.RemoteClusterSpec.Name, Namespace: req.Namespace}
	if err := r.Get(ctx, ytsaurusName, &remoteYtsaurus); err != nil {
		logger.Error(err, "unable to fetch remote YTsaurus for the remote nodes")
		return ctrl.Result{RequeueAfter: time.Second * 10}, err
	}

	return r.Sync(ctx, &nodes, &remoteYtsaurus)
}

// SetupWithManager sets up the controller with the Manager.
func (r *RemoteExecNodesReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&ytv1.RemoteExecNodes{}).
		Complete(r)
}
