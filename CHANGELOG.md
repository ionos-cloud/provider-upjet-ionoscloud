# Changelog

All notable changes to this project will be documented in this file.

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