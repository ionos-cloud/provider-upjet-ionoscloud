# Migrating MongoDB Clusters from v1 to v2

IONOS Cloud has introduced a new DBaaS MongoDB API (v2) with breaking changes
compared to the original API (v1). All MongoDB clusters will eventually be
migrated to v2. In this provider, the two APIs are exposed as two separate
CRDs:

|                    | v1                         | v2                            |
|--------------------|----------------------------|-------------------------------|
| API group          | `mongodb.ionoscloud.io`    | `mongodbv2.ionoscloud.io`     |
| Kind               | `MongodbCluster`           | `MongodbCluster`              |
| Terraform resource | `ionoscloud_mongo_cluster` | `ionoscloud_mongo_cluster_v2` |

Migrating on the IONOS Cloud side is **in place**: an existing cluster keeps
its cluster ID, its data, and its connection string — it simply becomes
reachable through the v2 API. Nothing about your running database changes.

What *does* change is how you describe the cluster to Crossplane. Since v1
and v2 are different CRDs (different API group, different underlying Terraform
resource), Crossplane has no built-in conversion between them. You migrate by:

1. Telling Crossplane to stop deleting the real cluster if the old `v1`
   `MongodbCluster` custom resource (CR) is removed (`orphan` the resource).
2. Creating a new `v2` `MongodbCluster` CR whose fields describe the same
   cluster, with `crossplane.io/external-name` set to the existing cluster ID
   so Crossplane *adopts* the cluster instead of provisioning a new one.
3. Verifying the new CR reconciles against the existing cluster.
4. Deleting the old, now-orphaned `v1` CR.

This provider has shipped the `mongodbv2.ionoscloud.io` CRDs since provider
release `v0.6.2` (Terraform provider `v6.7.40`). Make sure you're running at
least that version before starting.

## Known limitations

v1 manages database users through a separate CRD, `MongodbUser` (in the
`mongodb.ionoscloud.io` group); the v1 `MongodbCluster` has no inline user at
all. **v2 has no equivalent `MongodbUser` CRD.** The v2 `MongodbCluster` spec
instead carries a single inline initial user/database under
`credentials.username` and `credentials.database`.

Because of this, the `credentials` block is **entirely new in v2** — there is
nothing in the v1 `MongodbCluster` to translate it from, so you always author
it by hand during the migration (see the field mapping and walkthrough below).

If you have `MongodbUser` resources, there is currently no way to represent
them under the v2 CRD. Options:

- Leave those `MongodbUser` v1 resources in place, pointed at the (now
  v2-managed) cluster via `clusterId`/`clusterIdRef` — they will keep
  reconciling against the same cluster, since the cluster ID doesn't change.
- Manage additional users directly with `mongosh` or another MongoDB client
  outside of Crossplane, until a v2 equivalent CRD exists.

## Field mapping

The shape of `spec.forProvider` changed substantially. Fields not listed below
are either unchanged or new-and-optional.

| v1 field                                      | v2 field                                     | Notes                                                                                                                       |
|-----------------------------------------------|----------------------------------------------|---------------------------------------------------------------------------------------------------------------------------|
| `displayName`                                 | `name`                                       | Renamed.                                                                                                                   |
| `mongodbVersion`                              | `version`                                    | Renamed, same values (e.g. `6.0`, `7.0`).                                                                                  |
| `edition`                                     | `edition`                                    | Value casing changed to upper case: `enterprise` → `ENTERPRISE`, `business` → `BUSINESS`. `playground` has no v2 equivalent. |
| `location` / `locationRef` / `locationSelector` | `location`                                 | v2 takes a **literal location string only** — there is no `locationRef`/`locationSelector` in v2.                          |
| `type` (`replicaset` / `sharded-cluster`)     | `sharding`                                   | No `type` field in v2. A sharded cluster is expressed with the `sharding` block (`enabled`, `shards`); a replica set simply omits `sharding`. |
| `shards`                                       | `sharding.shards`                            | Moved under the new `sharding` block.                                                                                       |
| `cores`                                        | `instances.cores`                            | Moved under the new `instances` block (ENTERPRISE edition).                                                                 |
| `ram` (MB)                                     | `instances.ram` (GB)                         | Moved under `instances`, **unit changed from MB to GB** (ENTERPRISE edition).                                               |
| `storageSize` (MB)                             | `instances.storageSize` (GB)                 | Moved under `instances`, **unit changed from MB to GB** (ENTERPRISE edition).                                               |
| `instances` (count)                            | `instances.count` / `instanceCount`          | For ENTERPRISE, the count moves under `instances.count`; for BUSINESS it becomes the top-level `instanceCount`.             |
| `templateId`                                   | `instanceTemplateId`                         | Renamed (BUSINESS edition).                                                                                                 |
| `storageType`                                  | *(removed)*                                  | No v2 equivalent; storage type is no longer selectable.                                                                     |
| `connections` (list, one entry)               | `connections` (single object)                | v1 accepted a list but only the first entry is used; v2 makes this a single object.                                         |
| `connections[].cidrList`                       | `connections.instanceAddresses`              | Renamed; still a list, one address per instance.                                                                            |
| `connections[].datacenterId(Ref/Selector)`     | `connections.datacenterId(Ref/Selector)`     | Unchanged (now inside the single object).                                                                                   |
| `connections[].lanId(Ref/Selector)`           | `connections.lanId(Ref/Selector)`            | Unchanged (now inside the single object).                                                                                   |
| `backup.location`                              | `snapshot.location`                          | The backup model changed: v2 uses a `snapshot` block.                                                                       |
| `backup.snapshotIntervalHours` / `backup.pointInTimeWindowHours` | *(removed)*                | No v2 equivalent.                                                                                                           |
| *(none)*                                        | `snapshot.retentionDays`                     | New and required. Accepted values: `3`, `4`, `5`, `6`, `7`, `15`, `30`, `60`, `90`, `120`, `180`, `360`.                   |
| `maintenanceWindow` (list, one entry)          | `maintenanceWindow` (single object)          | v1 accepted a list but only the first entry is used; v2 makes this a single object.                                         |
| *(none — separate `MongodbUser` CRD)*          | `credentials.username`                       | New. The initial user, inline on the cluster.                                                                              |
| *(none — secret on `MongodbUser`)*             | `credentials.passwordSecretRef`              | New. Write-only password, sourced from a Kubernetes `Secret`.                                                              |
| *(none)*                                        | `credentials.database`                       | New. Names the initial database inline.                                                                                     |
| *(none)*                                        | `description`                                | New, optional.                                                                                                             |
| *(none)*                                        | `logsEnabled` / `metricsEnabled`             | New, optional observability toggles.                                                                                       |
| *(none)*                                        | `restoreFromSnapshot.sourceSnapshotId` / `restoreFromSnapshot.recoveryTargetDatetime` | New, optional. Only for initializing from a snapshot; not needed when adopting an existing cluster.        |

v2 required fields are: `name`, `version`, `location`, `edition`,
`connections`, `credentials`, `maintenanceWindow`, and `snapshot`. In
addition, the sizing is edition-specific: **ENTERPRISE** requires the
`instances` block (`cores`, `count`, `ram`, `storageSize`), while **BUSINESS**
requires `instanceTemplateId` and `instanceCount`.

## Step-by-step walkthrough

### 1. Find the cluster ID of the existing v1 resource

```shell
kubectl get mongodbcluster.mongodb.ionoscloud.io example \
  -o jsonpath='{.metadata.annotations.crossplane\.io/external-name}'
```

```
7cf6b0b3-3edb-4e78-a039-4c5cef3d81ac
```

Record this ID — it's the actual IONOS Cloud cluster ID and will become the
`external-name` of the new v2 resource.

### 2. Orphan the old v1 resource

Before proceeding, verify that the running provider has management-policy
support enabled. It defaults to enabled, but `ENABLE_MANAGEMENT_POLICIES=false`
or a disabling command-line flag can turn it off. Do not continue while it is
disabled: the orphaning patch and Observe-only import below require this support.

Exclude `Delete` from `spec.managementPolicies` on the v1 `MongodbCluster` so
deleting the CR later does **not** delete the real cluster:

```shell
kubectl patch mongodbcluster.mongodb.ionoscloud.io example \
  --type=merge -p '{"spec":{"managementPolicies":["Observe"]}}'
```

This replaces the policy list with `Observe` only, preventing new create,
update, and delete operations from the v1 controller. Wait for v1 to reconcile
with this policy and for any in-flight operation to finish before applying or
enabling writes on the v2 resource.

### 3. Author the new v2 resource

There are two ways to build `spec.forProvider` for the new resource: translate
it by hand using the field mapping table (Option A), or let the provider read
the live cluster and generate it for you (Option B). Option B avoids manual
translation mistakes (e.g. the MB→GB unit change) and is the recommended
approach; Option A is documented for reference and for when you want full
control over the resulting manifest from the start.

#### Option A: Translate the spec by hand

Translate the old `spec.forProvider` using the field mapping table above. Set
`crossplane.io/external-name` to the cluster ID from step 1 so Crossplane
adopts the existing cluster instead of creating a new one.

Given a v1 (ENTERPRISE, replica set) resource like:

```yaml
apiVersion: mongodb.ionoscloud.io/v1alpha1
kind: MongodbCluster
metadata:
  name: example
spec:
  forProvider:
    displayName: MongoDB_cluster
    mongodbVersion: "7.0"
    edition: enterprise
    type: replicaset
    cores: 1
    ram: 8192        # MB
    storageSize: 10240 # MB
    storageType: HDD
    instances: 3
    connections:
        cidrList:
            - 192.168.1.108/24
            - 192.168.1.109/24
            - 192.168.1.110/24
        datacenterIdSelector:
          matchLabels:
            testing.upbound.io/example-name: datacenter_example
        lanIdSelector:
          matchLabels:
            testing.upbound.io/example-name: lan_example
    maintenanceWindow:
        dayOfTheWeek: Sunday
        time: "09:00:00"
    locationSelector:
      matchLabels:
        testing.upbound.io/example-name: datacenter_example
```

the equivalent v2 resource is:

```yaml
apiVersion: mongodbv2.ionoscloud.io/v1alpha1
kind: MongodbCluster
metadata:
  name: example
  annotations:
    crossplane.io/external-name: "7cf6b0b3-3edb-4e78-a039-4c5cef3d81ac" # cluster ID from step 1
spec:
  forProvider:
    name: MongoDB_cluster
    version: "7.0"
    edition: ENTERPRISE # was lower-case in v1
    location: gb/bhx # literal; v2 has no locationSelector/locationRef
    instances:
      cores: 1
      ram: 8 # MB -> GB
      storageSize: 10 # MB -> GB
      count: 3
    connections:
      instanceAddresses: # was cidrList
        - 192.168.1.108/24
        - 192.168.1.109/24
        - 192.168.1.110/24
      datacenterIdSelector:
        matchLabels:
          testing.upbound.io/example-name: datacenter_example
      lanIdSelector:
        matchLabels:
          testing.upbound.io/example-name: lan_example
    credentials: # new in v2 — no v1 equivalent to translate from
      username: dbadmin
      database: my_database
      passwordSecretRef:
        name: pwsecret
        namespace: upbound-system
        key: password
    snapshot:
      location: eu-central-3 # required in v2; see IONOS Cloud docs for a valid Object Storage location
      retentionDays: 7
    maintenanceWindow:
      dayOfTheWeek: Sunday
      time: "09:00:00"
```

> **Note:** `storageType`, `type`, the per-interval `backup` settings, and
> multiple `connections` entries have no v2 equivalent and are dropped.
> `credentials` is new in v2 and must be authored by hand (the password is
> write-only and comes from a `Secret` via `passwordSecretRef`). `snapshot`
> is required in v2 — check your cluster's actual snapshot/backup location in
> the IONOS Cloud console before setting it, since an incorrect value can
> force a new cluster. For a sharded cluster, add a `sharding` block
> (`enabled: true`, `shards: <n>`) instead of the v1 `type: sharded-cluster`.

Apply the new resource:

```shell
kubectl apply -f mongodbcluster-v2.yaml
```

#### Option B: Import with an Observe-only management policy (recommended)

Instead of hand-translating every field, apply a v2 resource with an
`Observe`-only management policy and a minimal `spec.forProvider`. The provider
will read the live cluster and populate `status.atProvider` for you — in the
*v2* field shape — which you then copy into `spec.forProvider` verbatim.

1. Apply a minimal v2 resource, with `crossplane.io/external-name` set to the
   cluster ID from step 1 and `managementPolicies: ["Observe"]`:

   ```yaml
   apiVersion: mongodbv2.ionoscloud.io/v1alpha1
   kind: MongodbCluster
   metadata:
     name: example
     annotations:
       crossplane.io/external-name: "7cf6b0b3-3edb-4e78-a039-4c5cef3d81ac" # cluster ID from step 1
   spec:
     managementPolicies: ["Observe"]
     forProvider: { location: gb/bhx } # location must match the cluster's actual location
   ```

   `Observe` skips the usual required-field validation on `forProvider` (e.g.
   `connections`, `credentials`, `snapshot`), but `location` is still needed
   even here — it's used client-side to pick the correct regional API endpoint
   before the provider can make any call at all, so it can't be left empty.
   Check the cluster's actual location in the IONOS Cloud console, or from the
   v1 resource's `spec.forProvider.location`/`locationSelector`.

   ```shell
   kubectl apply -f mongodbcluster-v2-observe.yaml
   ```

2. Wait for the resource to become `SYNCED`/`READY`, then read the observed
   state. `status.atProvider.id` and `status.atProvider.connectionUri` are
   observation-only — there is no matching field under `spec.forProvider`, so
   pasting them in as-is will be rejected by the CRD schema. Strip them with
   `jq` while extracting, so what you get back is already safe to paste into
   `spec.forProvider`:

   ```shell
   kubectl get mongodbcluster.mongodbv2.ionoscloud.io example -o json \
     | jq '.status.atProvider | del(.id, .connectionUri)'
   ```

   The result is already in the v2 field shape (`instances.cores`,
   `instances.ram` in GB, `snapshot.location`, etc.) — no manual unit
   conversion or field-mapping lookups required.

   **Note:** For this method, all IDs (e.g. `datacenterId`) will be returned as
   UUIDs, not as selectors or references. If you want to use a selector or
   reference, you will need to manually set that in the spec.

3. Copy the `jq` output into `spec.forProvider`. `status.atProvider.credentials`
   includes `username` and `database`, so (when the API returns them) the `jq`
   output already carries those under `credentials` and you don't need to
   re-type them. The only piece `status.atProvider` never contains is the
   write-only password, so you just need to add `credentials.passwordSecretRef`
   yourself — pointing at a `Secret` with the cluster's existing password.

4. Re-apply the manifest with the populated `spec.forProvider` from step 3,
   still with `managementPolicies: ["Observe"]`. This has to happen *before*
   removing `managementPolicies` since the required-field validation that
   `Observe` skips comes back the moment the policy changes:

   ```shell
   kubectl apply -f mongodbcluster-v2-observe.yaml
   ```

   Then switch the resource back to full management by removing
   `managementPolicies` (or setting it to `["*"]`):

   ```shell
   kubectl patch mongodbcluster.mongodbv2.ionoscloud.io example \
     --type=merge -p '{"spec":{"managementPolicies":null}}'
   ```

5. Confirm there's no drift: `SYNCED`/`READY` should stay `True` with no
   further update or replace triggered, since `spec.forProvider` now matches
   `status.atProvider` exactly.

### 4. Verify adoption

Confirm the new CR reconciles against the *same* cluster ID rather than
creating a new one:

```shell
kubectl get mongodbcluster.mongodbv2.ionoscloud.io example
```

```
NAME      SYNCED   READY   EXTERNAL-NAME                          AGE
example   True     True    7cf6b0b3-3edb-4e78-a039-4c5cef3d81ac   30s
```

The `EXTERNAL-NAME` must match the cluster ID recorded in step 1. If it
doesn't, Crossplane created a brand-new cluster instead of adopting the
existing one — delete the new CR (with the default `spec.managementPolicies`
of `["*"]`, which includes `Delete`) before it provisions further, and check
the `crossplane.io/external-name` annotation on the manifest you applied.

Also check for a clean diff between the applied spec and `status.atProvider` —
any drift here usually means a field was mistranslated in step 3 and will
trigger an update (or a replace, for immutable fields) against the real
cluster.

### 5. Remove the old v1 resource

Once the v2 resource is `SYNCED`/`READY` and matches the expected cluster:

**DO NOT** run `kubectl delete` on the old v1 resource unless you excluded
`Delete` from `spec.managementPolicies` in step 2.
If you delete the v1 resource without orphaning it first, Crossplane will
delete the real cluster.

```shell
kubectl delete mongodbcluster.mongodb.ionoscloud.io example
```

Because it was orphaned in step 2, this only removes the Kubernetes object —
the underlying cluster keeps running, now managed solely by the v2 resource.

If you also had `MongodbUser` resources for this cluster and want to keep
managing those users through Crossplane, leave them in place (see
[Known limitations](#known-limitations)) — do not delete them, as they still
reconcile against the same, unchanged cluster ID.

## Troubleshooting

- **New cluster created instead of adopted**: the
  `crossplane.io/external-name` annotation was missing or didn't match the
  real cluster ID exactly. Delete the erroneous CR and retry with the correct
  ID.
- **Update/replace triggered immediately after adoption**: a field in the new
  v2 spec doesn't match the cluster's actual current state (e.g. wrong
  `snapshot.location`, unit mix-up between MB and GB on `instances.ram`/
  `instances.storageSize`, or an `edition` left in lower case). Compare
  `status.atProvider` on the new resource against the old one's
  `status.atProvider` and correct the mismatched field.
- **`retentionDays` rejected**: `snapshot.retentionDays` only accepts a fixed
  set of values (`3`, `4`, `5`, `6`, `7`, `15`, `30`, `60`, `90`, `120`,
  `180`, `360`).
- **Old v1 resource deleted the real cluster**: this means `Delete` wasn't
  excluded from `spec.managementPolicies` (or the patch wasn't applied) before
  deletion. There is no way to recover a deleted cluster through Crossplane —
  restore from a snapshot via `restoreFromSnapshot` if one exists.
