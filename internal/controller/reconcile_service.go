package controller

import (
	"context"

	"github.com/go-logr/logr"
	v1core "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	qdrantv1alpha1 "qdrantoperator.io/operator/api/v1alpha1"
)

func (r *QdrantClusterReconciler) reconcileService(ctx context.Context, log logr.Logger, obj *qdrantv1alpha1.QdrantCluster) error {

	err := r.serviceHandle(ctx, log, obj, obj.GetServiceName(), map[string]string{
		"cluster": obj.Name,
	})
	if err != nil {
		return err
	}

	err = r.serviceHandle(ctx, log, obj, obj.GetServiceName()+"-ephemeral", map[string]string{
		"cluster":                  obj.Name,
		"qdrant-ephemeral-storage": "true",
	})
	if err != nil {
		return err
	}

	err = r.serviceHandle(ctx, log, obj, obj.GetServiceName()+"-persistent", map[string]string{
		"cluster":                  obj.Name,
		"qdrant-ephemeral-storage": "false",
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *QdrantClusterReconciler) serviceHandle(ctx context.Context, _ logr.Logger, obj *qdrantv1alpha1.QdrantCluster, serviceName string, selector map[string]string) error {

	objectMeta := v1meta.ObjectMeta{
		Name:      serviceName,
		Namespace: obj.Namespace,
		OwnerReferences: []v1meta.OwnerReference{{
			APIVersion: obj.APIVersion,
			Kind:       obj.Kind,
			Name:       obj.Name,
			UID:        obj.UID,
		}},
	}
	typeMeta := v1meta.TypeMeta{APIVersion: "v1", Kind: "Service"}
	ports := []v1core.ServicePort{{
		Name:     "http",
		Port:     6333,
		Protocol: "TCP",
	}, {
		Name:     "grpc",
		Port:     6334,
		Protocol: "TCP",
	}, {
		Name:     "p2p",
		Port:     6335,
		Protocol: "TCP",
	}}

	/**
	* Main
	 */
	service := &v1core.Service{
		ObjectMeta: objectMeta,
		TypeMeta:   typeMeta,
		Spec: v1core.ServiceSpec{
			Ports:    ports,
			Type:     "ClusterIP",
			Selector: selector,
		},
	}

	existingService := &v1core.Service{}
	if err := r.Get(ctx, types.NamespacedName{
		Name:      serviceName,
		Namespace: obj.Namespace,
	}, existingService); err != nil {

		if err := r.Client.Create(ctx, service); err != nil {
			return err
		}
	} else {
		existingService.Spec = service.Spec

		if err := r.Client.Update(ctx, existingService); err != nil {
			return err
		}
	}

	return nil
}

func (r *QdrantClusterReconciler) reconcileHeadlessService(ctx context.Context, log logr.Logger, obj *qdrantv1alpha1.QdrantCluster) error {
	service := &v1core.Service{
		ObjectMeta: v1meta.ObjectMeta{
			Name:      obj.GetHeadlessServiceName(),
			Namespace: obj.Namespace,
			OwnerReferences: []v1meta.OwnerReference{{
				APIVersion: obj.APIVersion,
				Kind:       obj.Kind,
				Name:       obj.Name,
				UID:        obj.UID,
			}},
		},
		TypeMeta: v1meta.TypeMeta{APIVersion: "v1", Kind: "Service"},
		Spec: v1core.ServiceSpec{
			Ports: []v1core.ServicePort{{
				Name:     "http",
				Port:     6333,
				Protocol: "TCP",
			}, {
				Name:     "grpc",
				Port:     6334,
				Protocol: "TCP",
			}, {
				Name:     "p2p",
				Port:     6335,
				Protocol: "TCP",
			}},

			PublishNotReadyAddresses: true,
			ClusterIP:                "None",
			Selector: map[string]string{
				"cluster": obj.Name,
			},
		},
	}

	existingService := &v1core.Service{}

	if err := r.Get(ctx, types.NamespacedName{
		Name:      obj.GetHeadlessServiceName(),
		Namespace: obj.Namespace,
	}, existingService); err != nil {
		log.Info("Deploying headless Service...")
		if err := r.Client.Create(ctx, service); err != nil {
			return err
		}
	} else {
		existingService.Spec.Ports = service.Spec.Ports

		if err := r.Client.Update(ctx, existingService); err != nil {
			return err
		}
	}
	return nil
}
