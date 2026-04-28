# Changelog

## 0.1.0-alpha.17 (2026-04-28)

Full Changelog: [v0.1.0-alpha.16...v0.1.0-alpha.17](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.1.0-alpha.16...v0.1.0-alpha.17)

### Features

* support setting headers via env ([81e12ea](https://github.com/bluehive-health/bluehive-sdk-go/commit/81e12ea5e68dfa00b00ee0b9196bb987b82742fe))

## 0.1.0-alpha.16 (2026-04-26)

Full Changelog: [v0.1.0-alpha.15...v0.1.0-alpha.16](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.1.0-alpha.15...v0.1.0-alpha.16)

### Features

* **go:** add default http client with timeout ([9896e9a](https://github.com/bluehive-health/bluehive-sdk-go/commit/9896e9aa1067a478e10144960a4de6d283ab2e44))


### Chores

* **internal:** more robust bootstrap script ([f0a6f59](https://github.com/bluehive-health/bluehive-sdk-go/commit/f0a6f59993f8b86f3f36b9cb80422c5ab4166e1a))

## 0.1.0-alpha.15 (2026-04-09)

Full Changelog: [v0.1.0-alpha.14...v0.1.0-alpha.15](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.1.0-alpha.14...v0.1.0-alpha.15)

### Features

* **api:** api update ([568aa91](https://github.com/bluehive-health/bluehive-sdk-go/commit/568aa917bfd7ec2a116ba8ffbe7b4133a6a09086))
* **api:** api update ([e0c3ab6](https://github.com/bluehive-health/bluehive-sdk-go/commit/e0c3ab638169c854c02de778a64a3335ca0dedb0))
* **api:** api update ([1f971cd](https://github.com/bluehive-health/bluehive-sdk-go/commit/1f971cd55e4e82f75ac056877156425ede7a2f36))
* **api:** api update ([d5c7ee1](https://github.com/bluehive-health/bluehive-sdk-go/commit/d5c7ee16c1a6e8715b02084b7f7236ace50bd086))
* **api:** api update ([6743aa0](https://github.com/bluehive-health/bluehive-sdk-go/commit/6743aa00dbab38e1e131fa660818c9526e1467d1))
* **api:** api update ([a870ef3](https://github.com/bluehive-health/bluehive-sdk-go/commit/a870ef3227cb9dc16c6307664a0f14c7e58e3bb4))
* **api:** api update ([8a0034c](https://github.com/bluehive-health/bluehive-sdk-go/commit/8a0034c38c604fc22c89a3514b536e165b45a1cd))
* **api:** api update ([e394587](https://github.com/bluehive-health/bluehive-sdk-go/commit/e394587fccfbc33898730b9afe96e94e94ec0307))
* **api:** manual updates ([b08e8f4](https://github.com/bluehive-health/bluehive-sdk-go/commit/b08e8f4686fe293c5146dc542e4e461b5d4125dd))
* **api:** manual updates ([3f73510](https://github.com/bluehive-health/bluehive-sdk-go/commit/3f735109c8f3cbd5b46c9e0c8faec3899a930697))
* **client:** add a convenient param.SetJSON helper ([b32ee94](https://github.com/bluehive-health/bluehive-sdk-go/commit/b32ee949ab9389c96938489985d012805e790e2c))
* **encoder:** support bracket encoding form-data object members ([5beec67](https://github.com/bluehive-health/bluehive-sdk-go/commit/5beec678b963f170e1ceab99d16868304cb1cfc5))
* **internal:** support comma format in multipart form encoding ([256862a](https://github.com/bluehive-health/bluehive-sdk-go/commit/256862ad5e1a01d51ff0b9a60b75c7ec5103d8e2))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([4c791ea](https://github.com/bluehive-health/bluehive-sdk-go/commit/4c791eacee66577ab615ff99796b27222c67328f))
* **client:** correctly specify Accept header with */* instead of empty ([da4aace](https://github.com/bluehive-health/bluehive-sdk-go/commit/da4aace3b4dc17c738263f780c74c386a258fa30))
* **client:** make sure to import param package when used ([7985aad](https://github.com/bluehive-health/bluehive-sdk-go/commit/7985aadd0a6c75d537d6e2f347ebe0a834ba018f))
* **client:** use correct format specifier for header serialization ([126df2d](https://github.com/bluehive-health/bluehive-sdk-go/commit/126df2d5f74652998056ca4e1861044b9cb6e929))
* **docs:** add missing pointer prefix to api.md return types ([5347162](https://github.com/bluehive-health/bluehive-sdk-go/commit/53471626928eac618ea9bcf44ae78b169e865097))
* **docs:** fix mcp installation instructions for remote servers ([50f4075](https://github.com/bluehive-health/bluehive-sdk-go/commit/50f4075f7a6647f59c374030c2c6da5362dc1891))
* **encoder:** correctly serialize NullStruct ([aa17980](https://github.com/bluehive-health/bluehive-sdk-go/commit/aa179808de85eed847e5c4dc3bc4f09676e6b5ae))
* fix for union type names ([740ba25](https://github.com/bluehive-health/bluehive-sdk-go/commit/740ba2597dd2b5e01bc923f5cc7714060749d93d))
* **mcp:** correct code tool API endpoint ([5054b4a](https://github.com/bluehive-health/bluehive-sdk-go/commit/5054b4a2f6f7b1af2cf30ed71981639ca9652cd3))
* prevent duplicate ? in query params ([5e5c3ea](https://github.com/bluehive-health/bluehive-sdk-go/commit/5e5c3ea1be8c863a804c32823fbb62008993c956))
* rename param to avoid collision ([6b3c1a7](https://github.com/bluehive-health/bluehive-sdk-go/commit/6b3c1a7e9a9dc261b1cc91b178699c282b71bfe6))
* skip usage tests that don't work with Prism ([16321fc](https://github.com/bluehive-health/bluehive-sdk-go/commit/16321fcc99b9a01c0a04e1e0515ab18887b999e4))


### Chores

* add float64 to valid types for RegisterFieldValidator ([fc27a8b](https://github.com/bluehive-health/bluehive-sdk-go/commit/fc27a8b3ae6eb2c2ead205fbb6a096bd5f903cad))
* bump gjson version ([cc15261](https://github.com/bluehive-health/bluehive-sdk-go/commit/cc152612205d9b4e853d5c3a0d4155744c2f7bcf))
* **ci:** skip lint on metadata-only changes ([1d0498b](https://github.com/bluehive-health/bluehive-sdk-go/commit/1d0498bec0c1d0e6289d14339b9e682ecdb0b6a2))
* **ci:** skip uploading artifacts on stainless-internal branches ([1f2f07d](https://github.com/bluehive-health/bluehive-sdk-go/commit/1f2f07dc1eff60ec816105d9d8f1e6f3d85846d7))
* **ci:** support opting out of skipping builds on metadata-only commits ([f0dee03](https://github.com/bluehive-health/bluehive-sdk-go/commit/f0dee033818ad5157697ed185ccba7438dcdf752))
* **client:** fix multipart serialisation of Default() fields ([8e2fb1b](https://github.com/bluehive-health/bluehive-sdk-go/commit/8e2fb1bbe38b76108c3adb688a2568cbacc2ed0a))
* elide duplicate aliases ([254622c](https://github.com/bluehive-health/bluehive-sdk-go/commit/254622c56501c05195f4c0b70fd996ef734bfa18))
* **internal:** codegen related update ([3eea328](https://github.com/bluehive-health/bluehive-sdk-go/commit/3eea328ca84e7fa2a3778b8832b5a88d1827a2ee))
* **internal:** codegen related update ([eeda857](https://github.com/bluehive-health/bluehive-sdk-go/commit/eeda85794cfdefda1a412b56071a8262b963be07))
* **internal:** codegen related update ([41893a0](https://github.com/bluehive-health/bluehive-sdk-go/commit/41893a08f4518df9747d9b5caa50f6ea234c3d96))
* **internal:** codegen related update ([9a0b14e](https://github.com/bluehive-health/bluehive-sdk-go/commit/9a0b14e5d04a5a92dc15a2621870f1fe66cc3713))
* **internal:** grammar fix (it's -&gt; its) ([6b45c7e](https://github.com/bluehive-health/bluehive-sdk-go/commit/6b45c7e2b2c94dc65cd22d3adcff599efb603fe7))
* **internal:** minor cleanup ([a473228](https://github.com/bluehive-health/bluehive-sdk-go/commit/a47322817ffa83b121adb5eda45071486f55112d))
* **internal:** move custom custom `json` tags to `api` ([6eb7fad](https://github.com/bluehive-health/bluehive-sdk-go/commit/6eb7fad412b7b88e4f2e4c343f5bf82736ab8d49))
* **internal:** remove mock server code ([8c835c4](https://github.com/bluehive-health/bluehive-sdk-go/commit/8c835c4a03c9aee4f2b81970aa0555e60b9f1110))
* **internal:** support default value struct tag ([2aadfee](https://github.com/bluehive-health/bluehive-sdk-go/commit/2aadfee4d9aabee67aa775ad1ce9db94905576cf))
* **internal:** tweak CI branches ([9e02176](https://github.com/bluehive-health/bluehive-sdk-go/commit/9e02176989242616a8dba95aabb0f26bfeded945))
* **internal:** update `actions/checkout` version ([e9826dc](https://github.com/bluehive-health/bluehive-sdk-go/commit/e9826dccc2f370400d51cea3b55bc9938a972e70))
* **internal:** update gitignore ([cdbc6ba](https://github.com/bluehive-health/bluehive-sdk-go/commit/cdbc6ba79a1d80cb7e9f0d63f4b2d77855feda76))
* **internal:** use explicit returns ([08e3994](https://github.com/bluehive-health/bluehive-sdk-go/commit/08e3994d65d3ebc77468eb7725d5a7138610d0a4))
* **internal:** use explicit returns in more places ([647b9a4](https://github.com/bluehive-health/bluehive-sdk-go/commit/647b9a4dcf8ffdd0d0e27e44103acf6b33440350))
* remove unnecessary error check for url parsing ([6be7dfb](https://github.com/bluehive-health/bluehive-sdk-go/commit/6be7dfba26b1494ce9c3b6d4d75a7c46f7cd3c25))
* update docs for api:"required" ([6fd8c71](https://github.com/bluehive-health/bluehive-sdk-go/commit/6fd8c7147af7c97925f07afa900a1cc0c1e294ab))
* update mock server docs ([61d6e3e](https://github.com/bluehive-health/bluehive-sdk-go/commit/61d6e3e79cc40f7ba4f7548759a72e47ebca43e9))


### Documentation

* prominently feature MCP server setup in root SDK readmes ([1739026](https://github.com/bluehive-health/bluehive-sdk-go/commit/17390262897acfa7cc8b0978574d688d350b1115))

## 0.1.0-alpha.14 (2025-10-05)

Full Changelog: [v0.1.0-alpha.13...v0.1.0-alpha.14](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.1.0-alpha.13...v0.1.0-alpha.14)

### Features

* **api:** api update ([7e1f54e](https://github.com/bluehive-health/bluehive-sdk-go/commit/7e1f54e9a1dd1445aaa41de8e381bf8d202f94d3))
* **api:** api update ([6fc7d2a](https://github.com/bluehive-health/bluehive-sdk-go/commit/6fc7d2a9733000760e61dc549fac75a4f38c5d07))
* **api:** manual updates ([1284584](https://github.com/bluehive-health/bluehive-sdk-go/commit/1284584774a84a624e762a1bf1b014f059ed0533))

## 0.1.0-alpha.13 (2025-10-03)

Full Changelog: [v0.1.0-alpha.12...v0.1.0-alpha.13](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.1.0-alpha.12...v0.1.0-alpha.13)

### Features

* **api:** api update ([cc72e14](https://github.com/bluehive-health/bluehive-sdk-go/commit/cc72e145a4e5b596487572dec075091f6fab23fa))

## 0.1.0-alpha.12 (2025-10-03)

Full Changelog: [v0.1.0-alpha.11...v0.1.0-alpha.12](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.1.0-alpha.11...v0.1.0-alpha.12)

### Features

* **api:** api update ([57d1541](https://github.com/bluehive-health/bluehive-sdk-go/commit/57d1541f1e322000b3e889e8b185f5f39a206153))

## 0.1.0-alpha.11 (2025-09-26)

Full Changelog: [v0.1.0-alpha.10...v0.1.0-alpha.11](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.1.0-alpha.10...v0.1.0-alpha.11)

### Bug Fixes

* bugfix for setting JSON keys with special characters ([47a0ba7](https://github.com/bluehive-health/bluehive-sdk-go/commit/47a0ba779b5bc8a3bb0d68f7eadf79c529ae35ff))

## 0.1.0-alpha.10 (2025-09-20)

Full Changelog: [v0.1.0-alpha.9...v0.1.0-alpha.10](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.1.0-alpha.9...v0.1.0-alpha.10)

### Bug Fixes

* use slices.Concat instead of sometimes modifying r.Options ([f6f9ecc](https://github.com/bluehive-health/bluehive-sdk-go/commit/f6f9eccdd0434535e4e32bda7ab88d9643dba398))


### Chores

* bump minimum go version to 1.22 ([876effd](https://github.com/bluehive-health/bluehive-sdk-go/commit/876effde6e997136b2daa9c4176187ff484416f9))
* do not install brew dependencies in ./scripts/bootstrap by default ([53e46f4](https://github.com/bluehive-health/bluehive-sdk-go/commit/53e46f4cb74a927244617bb97044f5e7ca8a40a8))
* update more docs for 1.22 ([465efa0](https://github.com/bluehive-health/bluehive-sdk-go/commit/465efa0f8bbd716f1f85485942a10130d2d1ff71))

## 0.1.0-alpha.9 (2025-09-06)

Full Changelog: [v0.1.0-alpha.8...v0.1.0-alpha.9](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.1.0-alpha.8...v0.1.0-alpha.9)

### Bug Fixes

* **internal:** unmarshal correctly when there are multiple discriminators ([23bbdd1](https://github.com/bluehive-health/bluehive-sdk-go/commit/23bbdd1b10f5f32c0c9b64a89c5cee42503bdefb))


### Chores

* **internal:** codegen related update ([4d0f136](https://github.com/bluehive-health/bluehive-sdk-go/commit/4d0f136e9b70d48401c838870ba515069ffdc443))

## 0.1.0-alpha.8 (2025-08-29)

Full Changelog: [v0.1.0-alpha.7...v0.1.0-alpha.8](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.1.0-alpha.7...v0.1.0-alpha.8)

### Bug Fixes

* close body before retrying ([b7857b0](https://github.com/bluehive-health/bluehive-sdk-go/commit/b7857b0d1853e4a195ae9c84474256f99096afa8))

## 0.1.0-alpha.7 (2025-08-13)

Full Changelog: [v0.1.0-alpha.6...v0.1.0-alpha.7](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.1.0-alpha.6...v0.1.0-alpha.7)

### Features

* **api:** api update ([0e78e21](https://github.com/bluehive-health/bluehive-sdk-go/commit/0e78e212a1ae26d03b69fb46765cdc5a938bcde9))

## 0.1.0-alpha.6 (2025-08-13)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Features

* **api:** api update ([400298d](https://github.com/bluehive-health/bluehive-sdk-go/commit/400298db458e45cccab197c8f462c16d6d43d319))

## 0.1.0-alpha.5 (2025-08-13)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Features

* **api:** manual updates ([b67d3df](https://github.com/bluehive-health/bluehive-sdk-go/commit/b67d3dffd825a8c3c357df0fe1ba9013f3e5a07a))

## 0.1.0-alpha.4 (2025-08-13)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** api update ([bbe45e7](https://github.com/bluehive-health/bluehive-sdk-go/commit/bbe45e78144bd170ee36085ba64f1370867540f6))

## 0.1.0-alpha.3 (2025-08-13)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** manual updates ([16e1418](https://github.com/bluehive-health/bluehive-sdk-go/commit/16e1418ac375e8044559306736f654a896ac4b17))

## 0.1.0-alpha.2 (2025-08-13)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Chores

* **internal:** codegen related update ([04b0fd4](https://github.com/bluehive-health/bluehive-sdk-go/commit/04b0fd4813b1e0b91ded144e4183b5ad1483dd81))

## 0.1.0-alpha.1 (2025-08-09)

Full Changelog: [v0.0.1-alpha.4...v0.1.0-alpha.1](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.0.1-alpha.4...v0.1.0-alpha.1)

### Features

* **client:** support optional json html escaping ([4c282fb](https://github.com/bluehive-health/bluehive-sdk-go/commit/4c282fb4f48296e7a990db66a44a18d66307cf9f))


### Chores

* **internal:** update comment in script ([c73b43f](https://github.com/bluehive-health/bluehive-sdk-go/commit/c73b43f1b04e13249d11e3402b6e6aa21227ed36))
* update @stainless-api/prism-cli to v5.15.0 ([3fe4a47](https://github.com/bluehive-health/bluehive-sdk-go/commit/3fe4a47803f24dacbc8716a6299b3ce8e03eaaf6))

## 0.0.1-alpha.4 (2025-07-22)

Full Changelog: [v0.0.1-alpha.3...v0.0.1-alpha.4](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.0.1-alpha.3...v0.0.1-alpha.4)

### Bug Fixes

* **client:** process custom base url ahead of time ([af9233d](https://github.com/bluehive-health/bluehive-sdk-go/commit/af9233dde50dee94fb85ddb9fd9fd0ee9fbd691c))


### Chores

* lint tests in subpackages ([0911f5f](https://github.com/bluehive-health/bluehive-sdk-go/commit/0911f5f8030aae9b1a2023a600b9e40332e54965))

## 0.0.1-alpha.3 (2025-07-09)

Full Changelog: [v0.0.1-alpha.2...v0.0.1-alpha.3](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.0.1-alpha.2...v0.0.1-alpha.3)

### Chores

* **internal:** fix lint script for tests ([edcfef1](https://github.com/bluehive-health/bluehive-sdk-go/commit/edcfef168adebae4820fe4e727b57be5d42d25df))

## 0.0.1-alpha.2 (2025-07-08)

Full Changelog: [v0.0.1-alpha.1...v0.0.1-alpha.2](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.0.1-alpha.1...v0.0.1-alpha.2)

### Chores

* lint tests ([310e678](https://github.com/bluehive-health/bluehive-sdk-go/commit/310e678ddf13aab96ef21439e17a64a0bb485df1))

## 0.0.1-alpha.1 (2025-07-06)

Full Changelog: [v0.0.1-alpha.0...v0.0.1-alpha.1](https://github.com/bluehive-health/bluehive-sdk-go/compare/v0.0.1-alpha.0...v0.0.1-alpha.1)

### Chores

* configure new SDK language ([f92d55e](https://github.com/bluehive-health/bluehive-sdk-go/commit/f92d55e069b3411230ea4cc3eaaa4c3b3388ddad))
* update SDK settings ([f4cd4dd](https://github.com/bluehive-health/bluehive-sdk-go/commit/f4cd4ddebf96c7d913648a0efaad7d1a673db4f0))
* update SDK settings ([045f3f1](https://github.com/bluehive-health/bluehive-sdk-go/commit/045f3f187d99e4836a2401a3c977fa5ba0a43a07))
