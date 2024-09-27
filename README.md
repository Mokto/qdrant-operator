# qdrant-operator

Qdrant-operator is an operator for Qdrant that allows:

- Graceful scaling up: will balance all shards automatically
- Graceful scaling down (will move all shards out of a node before scaling down)
- Supports multi stateful set
    - that way you can support one set on normal storage and one on localstorage for optimal performances/price for example.
    - also the operator makes sure there is 1 replica of each shard not on localstorage to avoid data loss






## Install the operator

- Deploy the operator with Helm

- helm install qdrant-operator -n qdrant-operator https://mokto.github.io/qdrant-operator/charts/qdrant-operator

## Install your cluster

- Recommended: using Qdrant version > v1.11.5 (prior versions haven't been tested although they should work)

- Start a cluster like this (this is just an example, feel free to adjust it to your needs):

```yaml
kind: QdrantCluster
apiVersion: qdrant.qdrantoperator.io/v1alpha1
metadata:
  name: qdrant
spec:
    image: "qdrant/qdrant:v1.11.5"
    apiKey: myapikey # optional
    config: | # set the config that you like
      storage:
        async_scorer: true
        snapshots_config:
          snapshots_storage: s3
      optimizers:
        default_segment_number: 15
        max_optimization_threads: 8
    statefulsets: 
    - replicas: 2
      name: main
      volumeClaim:
        storageClassName: gp3
        storage: 1Gi
      priorityClassName: "critical"
      resources:
        requests:
          memory: "1Gi"
          cpu: "100m"

    ## Optional: if you want to deploy ephemeral storage
    - replicas: 10
      name: nvme
      ephemeralStorage: true # mark this on the statefulset so that the operator can handle restarting the pods
      resources:
        requests:
          memory: "6Gi"
          cpu: "1"
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: nvme
                operator: Exists
      tolerations:
        - key: "nvme"
          operator: "Exists"
          effect: "NoSchedule"


```

- Feel free to change the statefulsets as you wish although you should note:
  - At the moment the operator doesn't delete statefulset if they are not part of the Custom resource anymore. It's instead recommended to use 0 as a replica value
  - You can edit replicas, resources, etc.. but don't change ephemeralStorage or volumeClaim values. It's untested and will probably break something. We will add some validation in the future to prevent those changes.

## Contributing

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

!!! If you change the Specs please run `make build-installer` and update in the helm charts

## License

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

