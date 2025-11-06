# Changelog

## [0.5.0](https://github.com/selesy/ethereum-mcp/compare/v0.4.1...v0.5.0) (2025-11-06)


### Features

* **gen:** generate "non-standard" methods from local OpenRPC ([7ab1920](https://github.com/selesy/ethereum-mcp/commit/7ab1920b047a83ca74588b2723829c3e6e6f407d))
* **gen:** include "non-standard" type definitions from local OpenRPC ([47d487a](https://github.com/selesy/ethereum-mcp/commit/47d487a049486afa4d6813a517c5f48fa194a105))
* **proxy:** Add ParamSpec and generated values ([3268a51](https://github.com/selesy/ethereum-mcp/commit/3268a51873be69dd19afcbf9829ea8761be614e5))

## [0.4.1](https://github.com/selesy/ethereum-mcp/compare/v0.4.0...v0.4.1) (2025-09-10)


### Bug Fixes

* generate valid JSONSchema filesmoves the package into the generator ([926ae00](https://github.com/selesy/ethereum-mcp/commit/926ae008da7e0db29b70d4f08ef89369f45c04b4))

## [0.4.0](https://github.com/selesy/ethereum-mcp/compare/v0.3.0...v0.4.0) (2025-09-02)


### Features

* **gen:** generate method parameter lists ([ef3c458](https://github.com/selesy/ethereum-mcp/commit/ef3c458a3f69ae2b78217c30d74631a739704c4c))

## [0.3.0](https://github.com/selesy/ethereum-mcp/compare/v0.2.0...v0.3.0) (2025-09-02)


### Features

* **tool:** generate mcp.Tool definitions ([cc02671](https://github.com/selesy/ethereum-mcp/commit/cc02671e8a3d8b998d5e93671bf672236457b6ec))

## [0.2.0](https://github.com/selesy/ethereum-mcp/compare/v0.1.1...v0.2.0) (2025-09-02)


### Features

* **gen:** expand generator to create schemas and tools ([eb764f9](https://github.com/selesy/ethereum-mcp/commit/eb764f97a8e872f2d7a66043d51d7455878bb58a))

## [0.1.1](https://github.com/selesy/ethereum-mcp/compare/v0.1.0...v0.1.1) (2025-09-02)


### Bug Fixes

* **gha:** try the easy fix for pre-release rate-limiting ([9cc45d4](https://github.com/selesy/ethereum-mcp/commit/9cc45d4ba9618215b59c9f37d13f40da4f72f8f0))

## [0.1.0](https://github.com/selesy/ethereum-mcp/compare/v0.0.0...v0.1.0) (2025-09-02)


### Features

* **gha:** copy workflows for pre-commit and release ([781e87b](https://github.com/selesy/ethereum-mcp/commit/781e87b15d9a1f507dc17eb980c81c1c405b4834))
* **openrpc:** add types representing the JSONSchema portions of OpenRPC ([b0af99e](https://github.com/selesy/ethereum-mcp/commit/b0af99e94c6796a13b5a851c4765c5cce13309a9))
* **openrpc:** recursively resolve references added to definitions ([cd3fd63](https://github.com/selesy/ethereum-mcp/commit/cd3fd63dc354b908cde4575aade20f0217dc5fe3))
* **schema:** add example creating an mcp.Tool from a value in schema ([5a007d2](https://github.com/selesy/ethereum-mcp/commit/5a007d237efb4d8bf40b68ef332ac8de0b5cdcd4))
* **schema:** add the generated, "raw" JSONSchemas for Ethereum methods ([3ceb83b](https://github.com/selesy/ethereum-mcp/commit/3ceb83b08e50224fd3590f7e8d5e40ecedd41c80))
* **schema:** adds a generator to scrape Ethereum execution APIs ([84d7430](https://github.com/selesy/ethereum-mcp/commit/84d7430854e8dbc13754df137c926159f75b9eaf))


### Bug Fixes

* **deps:** update ethereum-mcp version in gen ([ec5e821](https://github.com/selesy/ethereum-mcp/commit/ec5e821f1804e9011ebac0c6b5c38dc40e691e92))
* **gen/schema:** add trailing newline to end of JSONSchema files ([ebde15d](https://github.com/selesy/ethereum-mcp/commit/ebde15d96e3674575ae9e8e4e4d4c20130f7a109))
* **gen/schema:** correct typo in Go doc for embedded schema source ([0205915](https://github.com/selesy/ethereum-mcp/commit/02059151a4a2428b8042153db8af00b074fbd460))
* **gen/schema:** fix GitHub authentication token ([cb74d2f](https://github.com/selesy/ethereum-mcp/commit/cb74d2f22db9adf2c9db54c82e1ca18169452780))
* **gen/schema:** fix GitHub authentication token ([296ec84](https://github.com/selesy/ethereum-mcp/commit/296ec84e344fc31ee56526968687e26f95342d6d))
* **gen/schema:** fix GitHub authentication token ([7155910](https://github.com/selesy/ethereum-mcp/commit/71559104c04aac1cc7d7cf706c6c835b029d5be2))
* **gen/schema:** fix GitHub authentication token ([47d7fb5](https://github.com/selesy/ethereum-mcp/commit/47d7fb5164a7f345e3b4e9a8bf232a4c49479269))
* **gen/schema:** ignore auth token if not configured ([d4c2336](https://github.com/selesy/ethereum-mcp/commit/d4c233633a9f9efa364ee627b2702295c8e321d4))
* **gen/schema:** run schema package generation as a Go tool ([b4e2474](https://github.com/selesy/ethereum-mcp/commit/b4e2474e66a700046b81cd84485d830692cd5290))
* **gen/schema:** use pseudo-version for parent module ([6b2e115](https://github.com/selesy/ethereum-mcp/commit/6b2e115633e55fd7b3ee211503ee4f0b0cbec83d))
* **jsonschema:** switch to selesy fork of invopop.jsonschema ([277bd8d](https://github.com/selesy/ethereum-mcp/commit/277bd8d232be35a66eb28f3c688606e10105621e))
* **schema:** embed schemas as json.RawMessage instead of string ([43342f0](https://github.com/selesy/ethereum-mcp/commit/43342f0b219333e8003964105b5ce9d64c877e29))
