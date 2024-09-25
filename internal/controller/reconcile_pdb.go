package controller

import (
	"context"

	"github.com/go-logr/logr"
	v1policy "k8s.io/api/policy/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"

	qdrantv1alpha1 "qdrantoperator.io/operator/api/v1alpha1"
	"qdrantoperator.io/operator/internal/qdrant"
)

func (r *QdrantClusterReconciler) reconcilePodDisruptionBudget(ctx context.Context, log logr.Logger, obj *qdrantv1alpha1.QdrantCluster) error {
	maxUnavailable := int32(1)
	for _, collection := range obj.Status.Collections {
		if collection.Status != qdrant.CollectionStatus_Green.String() {
			maxUnavailable = 0
			break
		}
	}
	pdb := &v1policy.PodDisruptionBudget{
		ObjectMeta: v1meta.ObjectMeta{
			Name:      obj.Name,
			Namespace: obj.Namespace,
			OwnerReferences: []v1meta.OwnerReference{{
				APIVersion: obj.APIVersion,
				Kind:       obj.Kind,
				Name:       obj.Name,
				UID:        obj.UID,
			}},
		},
		TypeMeta: v1meta.TypeMeta{APIVersion: "policy/v1", Kind: "PodDisruptionBudget"},
		Spec: v1policy.PodDisruptionBudgetSpec{
			MaxUnavailable: &intstr.IntOrString{IntVal: maxUnavailable},
			Selector: &v1meta.LabelSelector{
				MatchLabels: map[string]string{
					"cluster": obj.Name,
				},
			},
		},
	}

	existingPdb := &v1policy.PodDisruptionBudget{}

	log.Info("Deploying PDB")
	if err := r.Get(ctx, types.NamespacedName{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}, existingPdb); err != nil {

		if err := r.Client.Create(ctx, pdb); err != nil {
			return err
		}
	} else {
		existingPdb.Spec.Selector = pdb.Spec.Selector
		existingPdb.Spec.MaxUnavailable = pdb.Spec.MaxUnavailable

		if err := r.Client.Update(ctx, existingPdb); err != nil {
			return err
		}
	}
	return nil
}
