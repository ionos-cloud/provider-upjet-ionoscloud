# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]
- Updated Terraform provider to v6.7.36
- `image_password` (Volume, Server, CubeServer) and `hash` (InMemoryDBReplicaset `hashedPassword`) are now sourced from a Kubernetes `Secret` via a new `*SecretRef` field instead of a plaintext spec field, following the upstream provider marking these attributes `Sensitive`
- `server_side_encryption_customer_key` (ObjectStorage Object) and `server_side_encryption_customer_key` / `source_customer_key` (ObjectStorage ObjectCopy) are now sourced via `*SecretRef` fields for the same reason
- `secretKey` (ObjectStorage Key) and `secretkey` (ObjectStorage StorageAccesskey) are no longer exposed in `status.atProvider`, as these are now marked `Sensitive` upstream; the value is instead published in the resource's connection secret (`writeConnectionSecretToRef`), alongside a new `accesskey` entry, so both halves of the credential pair are available from a single secret
- Added `confidential` (Confidential Computing / SEV-SNP) and `enabledFeatures` fields to Server; added `enabledFeatures` to Datacenter
- Fixed stale/mismatched example manifests for Server, Volume, and CubeServer (cluster and namespaced) to use the new `*SecretRef` fields

## [0.5.8]
- Updated Terraform provider to v6.7.35
- Added MariaDB V2 cluster (ionoscloud_mariadb_cluster_v2) as a new Crossplane managed resource

## [0.5.7]
- Moved schema generation and the provider runtime from Terraform to OpenTofu v1.12.4
- Fixed a perpetual update loop on NIC `ips`, where a DHCP-assigned address copied into spec by late-initialization was continuously re-applied as a desired value
- Reworked the e2e image password: sourced from the `UPTEST_IMAGE_PASSWORD` GitHub secret and provisioned as the `example-password` secret, with compute examples referencing it via key `attribute.result`

## [0.5.6]
- Updated Terraform provider to v6.7.32
- Added InMemoryDB V2 cluster (ionoscloud_inmemorydb_cluster_v2) as a new Crossplane managed resource

## [0.5.5]
- Updated Terraform provider to v6.7.30

## [0.5.4]
- Update crossplane-runtime to v2.2.1
- Updated Terraform provider to v6.7.29
- Remove stale files from the `apis` directory

## [0.5.3]

### Changed
- Updated Terraform provider to v6.7.23
- Updated Go to v1.26.1
- Updated linter configuration
- Added `contract_number` support in Terraform client options and environment variable

## [v0.5.0] - 2025-09-11

> Note: v0.5.0 and v0.1.5 point to the same release.

### Added
- ClusterProviderConfig support (cluster-scoped provider configuration)
- Namespaced and cluster-scoped resource examples

### Fixed
- Fix referencing in cluster-scoped resources
- Fix linter issues with embedded types
- Fix CRD names for cluster-scoped resources
- Move examples to cluster scope
- Instantiate OperationTrackerStore

### Changed
- Updated build toolchain and dependencies
- Moved e2e tests to Chainsaw framework
- Updated docs to reference provider version 0.1.4

## [v0.1.5] - 2025-09-11

See v0.5.0.

## [v0.1.4] - 2025-08-01

### Added
- Async reconciliation for resources whose creation/deletion takes more than 1 minute

### Fixed
- e2e tests workflow improvements and small template fixes

### Changed
- Updated Terraform provider to v6.7.12

## [v0.1.3] - 2025-06-20

### Fixed
- Improved quickstart documentation
- Fix install documentation
- Fix version setting in release and publish steps (release-upbound.yml)

### Changed
- Updated Terraform provider to v6.7.7

## [v0.1.2] - 2025-06-17

### Added
- Upbound marketplace release workflow
- Publishing to Upbound registry (ionos-cloud org)

### Changed
- Separated release and CI workflows
- Updated dependencies (crossplane-runtime, upjet, Go modules)

## [v0.1.1] - 2025-03-03

### Changed
- Updated Terraform provider to v6.7.3

## [v0.1.0] - 2024-11-14

Initial release of the IONOS Cloud Crossplane provider based on Upjet.

### Added
- Compute resources: server, volume, NIC, firewall, IP block, LAN, snapshot, datacenter, image, IP failover, crossconnect, VCpu server, Cube server, boot device selection
- Application Load Balancer (ALB): load balancer, forwarding rule, target group
- Network Load Balancer (NLB): NLB, forwarding rule, flow log
- API Gateway: gateway, route
- Auto Scaling Group (ASG)
- CDN distribution
- Certificate Manager: certificate, auto-certificate, auto-certificate provider
- Container Registry: registry, repository, token
- Data Platform: cluster, node pool
- DNS: zone, record
- InMemoryDB: replica set
- Kafka: cluster, topic
- Kubernetes (K8s): cluster, node pool
- Logging pipeline
- MariaDB cluster
- MongoDB: cluster, collection, index, snapshot
- NAT Gateway: NAT gateway, rule, flow log
- NFS: cluster, share
- Object Storage: bucket, bucket ACL, bucket cors, bucket lifecycle, bucket policy, bucket public access block, bucket versioning, bucket website, object, object copy, access key
- PostgreSQL cluster
- VPN IPSec: gateway, tunnel
- VPN WireGuard: gateway, peer
- Group share resource
- Backup unit resource
- Quickstart guide and documentation
- E2E test framework