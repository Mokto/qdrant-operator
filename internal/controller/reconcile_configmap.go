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
}

func (r *QdrantClusterReconciler) reconcileConfigmap(ctx context.Context, log logr.Logger, obj *qdrantv1alpha1.QdrantCluster) (string, error) {

	falseValue := false

	var configObj map[string]interface{}

	if obj.Spec.Config != "" {
		if err := yaml.Unmarshal([]byte(obj.Spec.Config), &configObj); err != nil {
			log.Error(err, "Unmarshaling config object")
		}
	}

	for key, value := range defaultConfigObj {
		if _, ok := configObj[key]; !ok {
			configObj[key] = value
		}
	}

	if obj.Spec.ApiKey != "" {
		if configObj["service"] == nil {
			configObj["service"] = map[string]interface{}{}
		}
		configObj["service"].(map[string]interface{})["api_key"] = obj.Spec.ApiKey
	}

	bytes, err := yaml.Marshal(configObj)

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
	if obj.Status.HasBeenInited != nil && *obj.Status.HasBeenInited {
		dns := ""
		if leader != nil && leader.IsReady && !leader.EphemeralStorage {
			dns = leader.DNS
		}
		for _, peer := range obj.Status.Peers {
			if peer.IsReady && !peer.EphemeralStorage {
				dns = peer.DNS
				break
			}
		}
		if dns == "" {
			for _, peer := range obj.Status.Peers {
				if peer.IsReady {
					dns = peer.DNS
					break
				}
			}
		}
		if dns == "" {
			log.Info("No ready peers found. Not updating config")
			return "", nil
		}
		config = `
#!/bin/sh
LEADER=http://` + dns + `:6335
SELF=http://$HOSTNAME.` + obj.GetHeadlessServiceName() + `.` + obj.GetNamespace() + `:6335
echo "Leader is $LEADER. Self is $SELF"
if [[ "$LEADER" == "$SELF" ]]; then
	echo "Leader is self. Waiting for a new leader to be elected and restarting the pod."
	sleep 5
	exit 0
else
	exec ./entrypoint.sh --bootstrap $LEADER --uri $SELF
fi

`
		// exec ./entrypoint.sh --bootstrap 'http://` + leader.DNS + `:6335' --uri 'http://'"$HOSTNAME"'.` + obj.GetHeadlessServiceName() + `.` + obj.GetNamespace() + `:6335'
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
