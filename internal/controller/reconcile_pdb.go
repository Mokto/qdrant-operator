package controller

import (
	"context"

	"github.com/go-logr/logr"
	v1policy "k8s.io/api/policy/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"

	qdrantv1alpha1 "qdrantoperator.io/operator/api/v1alpha1"
)

func (r *QdrantClusterReconciler) reconcilePodDisruptionBudget(ctx context.Context, log logr.Logger, namespace string, obj *qdrantv1alpha1.QdrantCluster) error {
	pdb := &v1policy.PodDisruptionBudget{
		ObjectMeta: v1meta.ObjectMeta{
			Name:      obj.Name,
			Namespace: namespace,
			OwnerReferences: []v1meta.OwnerReference{{
				APIVersion: obj.APIVersion,
				Kind:       obj.Kind,
				Name:       obj.Name,
				UID:        obj.UID,
			}},
		},
		TypeMeta: v1meta.TypeMeta{APIVersion: "policy/v1", Kind: "PodDisruptionBudget"},
		Spec: v1policy.PodDisruptionBudgetSpec{
			MaxUnavailable: &intstr.IntOrString{IntVal: 0},
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
		Namespace: namespace,
	}, existingPdb); err != nil {

		if err := r.Client.Create(ctx, pdb); err != nil {
			return err
		}
	} else {
		existingPdb.Spec.Selector = pdb.Spec.Selector

		if err := r.Client.Update(ctx, existingPdb); err != nil {
			return err
		}
	}
	return nil
}
