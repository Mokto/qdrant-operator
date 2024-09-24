package controller

import (
	"context"

	"github.com/go-logr/logr"
	"gopkg.in/yaml.v3"
	v1core "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	qdrantv1alpha1 "qdrantoperator.io/operator/api/v1alpha1"
)

// var defaultConfig = "cluster:\n\tconsensus: \n\t\ttick_period_ms: 100"
var defaultConfigObj = map[string]interface{}{
	"cluster": map[string]interface{}{
		"consensus": map[string]interface{}{
			"tick_period_ms": 100,
		},
		"enabled": true,
		"p2p": map[string]interface{}{
			"port": 6335,
		},
	},
	"optimizers": map[string]interface{}{
		"default_segment_number":   15,
		"max_optimization_threads": 8,
	},
	"service": map[string]interface{}{
		// "api_key":             "0191bae0-49a3-707f-a5b7-ce08599bf5e5",
		"host":                "::",
		"max_request_size_mb": 32,
	},
	"storage": map[string]interface{}{
		"async_scorer": true,
	},
}

func (r *QdrantClusterReconciler) reconcileConfigmap(ctx context.Context, log logr.Logger, obj *qdrantv1alpha1.QdrantCluster) (string, error) {

	falseValue := false

	bytes, err := yaml.Marshal(defaultConfigObj)

	if err != nil {
		log.Error(err, "Marshaling config object")
		return "", err
	}

	data := map[string]string{
		"production.yaml": string(bytes),
	}
	hash := hashMap(data)

	config := `
	#!/bin/sh
	exec ./entrypoint.sh --uri 'http://'"$HOSTNAME"'.` + obj.GetHeadlessServiceName() + `.` + obj.GetNamespace() + `:6335'
	`

	leader := obj.Status.Peers.GetLeader()
	if leader != nil {
		config = `
#!/bin/sh
exec ./entrypoint.sh --bootstrap 'http://` + leader.DNS + `:6335' --uri 'http://'"$HOSTNAME"'.` + obj.GetHeadlessServiceName() + `.` + obj.GetNamespace() + `:6335'
`
	}

	data["initialize.sh"] = config

	configmap := &v1core.ConfigMap{
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
		Immutable: &falseValue,
		TypeMeta:  v1meta.TypeMeta{APIVersion: "v1", Kind: "ConfigMap"},
		Data:      data,
	}

	existingConfigmap := &v1core.ConfigMap{}

	log.Info("Deploying Configmap")
	if err := r.Get(ctx, types.NamespacedName{
		Name:      obj.Name,
		Namespace: obj.Namespace,
	}, existingConfigmap); err != nil {

		if err := r.Client.Create(ctx, configmap); err != nil {
			return "", err
		}
	} else {
		existingConfigmap.Data = configmap.Data

		if err := r.Client.Update(ctx, existingConfigmap); err != nil {
			return "", err
		}
	}
	return hash, nil
}
