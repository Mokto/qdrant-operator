# Changelog

## [0.17.1](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.17.0...qdrant-operator-v0.17.1) (2024-12-05)


### Bug Fixes

* deleting non ephemereal ([8380ae9](https://github.com/Mokto/qdrant-operator/commit/8380ae921611e3bf381d5ef55caab17d15f7b9a7))

## [0.17.0](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.16.4...qdrant-operator-v0.17.0) (2024-12-02)


### Features

* replicate shards faster ([f3626bc](https://github.com/Mokto/qdrant-operator/commit/f3626bc85fe01e8effd11389c970944ce7bacdc1))

## [0.16.4](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.16.3...qdrant-operator-v0.16.4) (2024-11-13)


### Bug Fixes

* wip ([b19dc37](https://github.com/Mokto/qdrant-operator/commit/b19dc37c4c83c5994c531244b97dbc595ef21f1f))

## [0.16.3](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.16.2...qdrant-operator-v0.16.3) (2024-11-13)


### Bug Fixes

* ready container ([debae2b](https://github.com/Mokto/qdrant-operator/commit/debae2bddb2f80fc104835b99c9f110ff09e4fa1))

## [0.16.2](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.16.1...qdrant-operator-v0.16.2) (2024-11-13)


### Bug Fixes

* better waiting system ([e0b7844](https://github.com/Mokto/qdrant-operator/commit/e0b784432bce8e85da0fa41a9d777a6b287e7b54))

## [0.16.1](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.16.0...qdrant-operator-v0.16.1) (2024-11-13)


### Bug Fixes

* ready container for collections with small amount of shards ([4a6f2f3](https://github.com/Mokto/qdrant-operator/commit/4a6f2f348855450eb1bd70f1074ca7a741bb22c8))

## [0.16.0](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.15.0...qdrant-operator-v0.16.0) (2024-11-12)


### Features

* improve stability ([616f855](https://github.com/Mokto/qdrant-operator/commit/616f85581a8f4a0a920b3470164c018417e077ef))

## [0.15.0](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.14.0...qdrant-operator-v0.15.0) (2024-11-11)


### Features

* only copy active shards ([01566f2](https://github.com/Mokto/qdrant-operator/commit/01566f2e7c153eff29a98c4962f5b5502b2d9d2b))

## [0.14.0](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.13.1...qdrant-operator-v0.14.0) (2024-11-08)


### Features

* deleting failed pods automatically ([7dfc198](https://github.com/Mokto/qdrant-operator/commit/7dfc19885f1d0f410acb344c1bcc17dbcf963d35))

## [0.13.1](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.13.0...qdrant-operator-v0.13.1) (2024-11-08)


### Bug Fixes

* sleep in the infinite loop ([4b332aa](https://github.com/Mokto/qdrant-operator/commit/4b332aae94602358250bc3ea68202421719d7b0b))

## [0.13.0](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.12.0...qdrant-operator-v0.13.0) (2024-11-08)


### Features

* log on which peer the leader hasn't been found ([8f9b618](https://github.com/Mokto/qdrant-operator/commit/8f9b6184567e46eb3c6fe9863cc7a86cb5f7b8ba))

## [0.12.0](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.11.0...qdrant-operator-v0.12.0) (2024-11-08)


### Features

* run clusters independently for more stability ([7b45527](https://github.com/Mokto/qdrant-operator/commit/7b45527e17d7bb34c87a5f365e6872ad7f6212c1))

## [0.11.0](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.10.0...qdrant-operator-v0.11.0) (2024-10-28)


### Features

* allow overidding resources ([2ce72b0](https://github.com/Mokto/qdrant-operator/commit/2ce72b04d246fb711bdb24a7d3212fd2cf7e8128))

## [0.10.0](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.9.0...qdrant-operator-v0.10.0) (2024-10-28)


### Features

* snapshot to copy shards ([e5fd6ce](https://github.com/Mokto/qdrant-operator/commit/e5fd6cebe529d6fc8da1f60f73dff2e499432268))

## [0.9.0](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.8.0...qdrant-operator-v0.9.0) (2024-10-22)


### Features

* new container ([49adb14](https://github.com/Mokto/qdrant-operator/commit/49adb14f15a6030de303be5789d59905a4dd335e))

## [0.8.0](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.7.0...qdrant-operator-v0.8.0) (2024-10-22)


### Features

* delete dead shards after copying is done ([fc3c978](https://github.com/Mokto/qdrant-operator/commit/fc3c978cd63e2bfde1095de97a870c35d43987e2))

## [0.7.0](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.6.4...qdrant-operator-v0.7.0) (2024-10-22)


### Features

* count a dead shard as missing ([931cdc1](https://github.com/Mokto/qdrant-operator/commit/931cdc193897ee5cbba37d9ee8b9d90fed1e8bde))

## [0.6.4](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.6.3...qdrant-operator-v0.6.4) (2024-10-15)


### Bug Fixes

* longer waiting period for the sidecar ([e2ad987](https://github.com/Mokto/qdrant-operator/commit/e2ad987ce63c2b6142f40e62ef30dea807578b77))

## [0.6.3](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.6.2...qdrant-operator-v0.6.3) (2024-10-15)


### Bug Fixes

* longer waiting period for the sidecar ([cb9b3d9](https://github.com/Mokto/qdrant-operator/commit/cb9b3d96758300f1193818eab164fb61718d1562))

## [0.6.2](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.6.1...qdrant-operator-v0.6.2) (2024-10-15)


### Bug Fixes

* all ephemeral storage ([866d7ed](https://github.com/Mokto/qdrant-operator/commit/866d7ed669ce65ecdb5bcae25f9d7ffbe34b1b13))

## [0.6.1](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.6.0...qdrant-operator-v0.6.1) (2024-10-15)


### Bug Fixes

* sidecar healthcheck was too strict ([9896e2d](https://github.com/Mokto/qdrant-operator/commit/9896e2d9a63890c301b0a744aa8422d4a09b9625))

## [0.6.0](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.5.4...qdrant-operator-v0.6.0) (2024-10-15)


### Features

* replicate missing shards to best candidate ([5379ec6](https://github.com/Mokto/qdrant-operator/commit/5379ec60397f836ca8077aa4a99bbeecb932948c))

## [0.5.4](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.5.3...qdrant-operator-v0.5.4) (2024-10-14)


### Bug Fixes

* only move shards if you find a candidate for it ([7c57b2d](https://github.com/Mokto/qdrant-operator/commit/7c57b2d7882fdaa4d78197120e7f369dc970deb6))

## [0.5.3](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.5.2...qdrant-operator-v0.5.3) (2024-10-13)


### Bug Fixes

* try to connect to the raft leader in priority ([8acad11](https://github.com/Mokto/qdrant-operator/commit/8acad11bca3b5841d905cc55e6b9c8d86dc6562c))

## [0.5.2](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.5.1...qdrant-operator-v0.5.2) (2024-10-13)


### Bug Fixes

* improved logging again ([3b00b28](https://github.com/Mokto/qdrant-operator/commit/3b00b28cc520555be943ea7a859ab03501872e85))

## [0.5.1](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.5.0...qdrant-operator-v0.5.1) (2024-10-12)


### Bug Fixes

* wrong docker image ([af998af](https://github.com/Mokto/qdrant-operator/commit/af998af1f2aac144fc794090732073afb5569c22))

## [0.5.0](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.4.0...qdrant-operator-v0.5.0) (2024-10-12)


### Features

* better logging ([8fb676a](https://github.com/Mokto/qdrant-operator/commit/8fb676ab7abac0f130ba82e65abd6d0e8dab8086))

## [0.4.0](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.3.11...qdrant-operator-v0.4.0) (2024-10-12)


### Features

* startup probes on ephemeral storages ([9f63842](https://github.com/Mokto/qdrant-operator/commit/9f63842a893e81a208080f8854a935b12ac66619))

## [0.3.11](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.3.10...qdrant-operator-v0.3.11) (2024-10-12)


### Bug Fixes

* crd ([a77b447](https://github.com/Mokto/qdrant-operator/commit/a77b447b8f7350de12140dc695e725a5bdab5fab))

## [0.3.10](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.3.9...qdrant-operator-v0.3.10) (2024-10-12)


### Bug Fixes

* hasbeeninited ([1fce96a](https://github.com/Mokto/qdrant-operator/commit/1fce96a9ba12111ffd8a26b62198637e48f38a53))
* hasbeeninited ([9ac957f](https://github.com/Mokto/qdrant-operator/commit/9ac957f60b8f127d00e6195eb93cd39ae1873d3b))

## [0.3.9](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.3.8...qdrant-operator-v0.3.9) (2024-10-11)


### Bug Fixes

* make leader election even safer ([44e6577](https://github.com/Mokto/qdrant-operator/commit/44e6577250bf66525d160e689bdc0dc4518b8fa8))

## [0.3.8](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.3.7...qdrant-operator-v0.3.8) (2024-10-10)


### Bug Fixes

* stability issues ([de39c0f](https://github.com/Mokto/qdrant-operator/commit/de39c0f5b91ad77d55313b9b6d39a79b6cf11ab8))

## [0.3.7](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.3.6...qdrant-operator-v0.3.7) (2024-10-07)


### Bug Fixes

* unknown status ([935b5cc](https://github.com/Mokto/qdrant-operator/commit/935b5ccc9d4e989161c077b36cd4bf88036a3507))

## [0.3.6](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.3.5...qdrant-operator-v0.3.6) (2024-10-03)


### Bug Fixes

* api key ([389987a](https://github.com/Mokto/qdrant-operator/commit/389987aa3ca3ceade20d389065ebf303af2841d9))

## [0.3.5](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.3.4...qdrant-operator-v0.3.5) (2024-10-03)


### Bug Fixes

* abort shards transfers first if needed ([64ff68f](https://github.com/Mokto/qdrant-operator/commit/64ff68fcc4eeef54fb81b0641d274a5a5e200c25))

## [0.3.4](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.3.3...qdrant-operator-v0.3.4) (2024-09-30)


### Bug Fixes

* max unavailable 0 on status unknown ([988a347](https://github.com/Mokto/qdrant-operator/commit/988a347917564110906cd80d93569a10bb0378dc))

## [0.3.3](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.3.2...qdrant-operator-v0.3.3) (2024-09-27)


### Bug Fixes

* memory requests ([1d74c65](https://github.com/Mokto/qdrant-operator/commit/1d74c65aaced9b0f342555792d1273116f8501f1))

## [0.3.2](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.3.1...qdrant-operator-v0.3.2) (2024-09-27)


### Bug Fixes

* inter namespace connection ([b88b260](https://github.com/Mokto/qdrant-operator/commit/b88b260c50492c6f592382d54b5c81729c8a546e))
* qdrant operator permission & image ([52c095e](https://github.com/Mokto/qdrant-operator/commit/52c095e40392952116e95ceb20ae32838d39117e))

## [0.3.1](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.3.0...qdrant-operator-v0.3.1) (2024-09-27)


### Bug Fixes

* env building ([a8d357c](https://github.com/Mokto/qdrant-operator/commit/a8d357ca44869f0fee24c9835b8e920996352852))

## [0.3.0](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.2.0...qdrant-operator-v0.3.0) (2024-09-27)


### Features

* deploy 2 additional services for ephemeral handling ([21848fa](https://github.com/Mokto/qdrant-operator/commit/21848faff679e9112530670f2329979b099360dd))

## [0.2.0](https://github.com/Mokto/qdrant-operator/compare/qdrant-operator-v0.1.0...qdrant-operator-v0.2.0) (2024-09-27)


### Features

* deploy helm chart & docker image ([#1](https://github.com/Mokto/qdrant-operator/issues/1)) ([a1d9b4f](https://github.com/Mokto/qdrant-operator/commit/a1d9b4fce9c9ff247cfc1d3398f45f4657fc3dbc))
