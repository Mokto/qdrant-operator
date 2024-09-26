# qdrant-operator

Qdrant-operator is an operator for Qdrant that allows:

- Graceful scaling up: will balance all shards automatically
- Graceful scaling down (will move all shards out of a node before scaling down)
- Supports multi stateful set
    - that way you can support one set on normal storage and one on localstorage for optimal performances/price for example.
    - also makes sure there is 1 replica of each shard not on localstorage to avoid data loss






## Getting Started



### To Deploy on the cluster


## Contributing

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

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

