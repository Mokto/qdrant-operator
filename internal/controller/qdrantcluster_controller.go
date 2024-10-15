/*
Copyright 2024.

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

package controller

import (
	"context"
	"time"

	"google.golang.org/grpc/metadata"
	v1 "k8s.io/api/apps/v1"
	v1core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/uuid"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	qdrantv1alpha1 "qdrantoperator.io/operator/api/v1alpha1"
)

// QdrantClusterReconciler reconciles a QdrantCluster object
type QdrantClusterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=qdrant.qdrantoperator.io,resources=qdrantclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=qdrant.qdrantoperator.io,resources=qdrantclusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=qdrant.qdrantoperator.io,resources=qdrantclusters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the QdrantCluster object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
// +kubebuilder:rbac:groups=examples.itamar.marom,resources=qdrantclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=examples.itamar.marom,resources=qdrantclusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=examples.itamar.marom,resources=qdrantclusters/finalizers,verbs=update
// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=get;list;watch;create;update;patch;delete
func (r *QdrantClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	requeueResult := ctrl.Result{
		RequeueAfter: time.Second * 30,
	}

	log := log.FromContext(ctx).WithValues("reconcileID", uuid.NewUUID())

	obj := &qdrantv1alpha1.QdrantCluster{}
	if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
		log.Error(err, "unable to fetch QdrantCluster")
		return requeueResult, client.IgnoreNotFound(err)
	}

	if obj.Spec.ApiKey != "" {
		ctx = metadata.AppendToOutgoingContext(ctx, "api-key", obj.Spec.ApiKey)
	}

	checksum, err := r.reconcileConfigmap(ctx, log, obj)
	if err != nil {
		log.Error(err, "unable to fetch update ConfigMap")
		return requeueResult, err
	}

	err = r.reconcileService(ctx, log, obj)
	if err != nil {
		log.Error(err, "unable to fetch update Service")
		return requeueResult, err
	}

	err = r.reconcileHeadlessService(ctx, log, obj)
	if err != nil {
		log.Error(err, "unable to fetch update headless Service")
		return requeueResult, err
	}

	err = r.reconcilePodDisruptionBudget(ctx, log, obj)
	if err != nil {
		log.Error(err, "unable to fetch update PodDisruptionBudget")
		return requeueResult, err
	}

	err = r.reconcileStatefulsets(ctx, log, obj, checksum)
	if err != nil {
		log.Error(err, "unable to fetch update StatefulSets")
		return requeueResult, err
	}

	log.Info("Reconcilied QdrantCluster " + obj.Name)

	if obj.Status.Peers.GetLeader() == nil {
		return requeueResult, nil
	}
	err = r.clearEmptyNodesFromScaleDown(ctx, log, obj)
	if err != nil {
		log.Error(err, "unable to clear empty nodes")
		return requeueResult, err
	}

	hasReplicatedShards, err := r.replicateMissingShards(ctx, log, obj)
	if err != nil {
		log.Error(err, "unable to duplicate shards")
		return requeueResult, err
	}
	if hasReplicatedShards {
		return requeueResult, nil
	}

	hasMovedShards, err := r.ensureShardsSafe(ctx, log, obj)
	if err != nil {
		log.Error(err, "unable to move shards to main nodes")
		return requeueResult, err
	}
	if hasMovedShards {
		return requeueResult, nil
	}

	err = r.balanceShards(ctx, log, obj)
	if err != nil {
		log.Error(err, "unable to trigger moving shards")
		return requeueResult, err
	}

	return requeueResult, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *QdrantClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&qdrantv1alpha1.QdrantCluster{}).
		Owns(&v1.StatefulSet{}).
		Owns(&v1core.ConfigMap{}).
		Owns(&v1core.Service{}).
		Owns(&v1core.ConfigMap{}).
		Watches(
			&v1core.ConfigMap{},
			handler.EnqueueRequestsFromMapFunc(r.findObjects),
		).
		Watches(
			&v1core.Service{},
			handler.EnqueueRequestsFromMapFunc(r.findObjects),
		).
		Watches(
			&v1.StatefulSet{},
			handler.EnqueueRequestsFromMapFunc(r.findObjects),
		).
		Complete(r)
}

func (r *QdrantClusterReconciler) findObjects(ctx context.Context, configMap client.Object) []reconcile.Request {

	if len(configMap.GetOwnerReferences()) == 0 {
		return []reconcile.Request{}
	}
	if configMap.GetOwnerReferences()[0].Name == "QdrantCluster" {
		return []reconcile.Request{
			{
				NamespacedName: types.NamespacedName{
					Name:      configMap.GetName(),
					Namespace: configMap.GetNamespace(),
				},
			},
		}
	}

	return []reconcile.Request{}
}
