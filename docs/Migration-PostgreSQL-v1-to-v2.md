# Migrating PostgreSQL Clusters from v1 to v2

IONOS Cloud has introduced a new DBaaS PostgreSQL API (v2) with breaking
changes compared to the original API (v1). All PostgreSQL clusters will
eventually be migrated to v2. In this provider, the two APIs are exposed as
two separate CRDs:

|                    | v1                         | v2                           |
|--------------------|----------------------------|------------------------------|
| API group          | `postgresql.ionoscloud.io` | `postgresqlv2.ionoscloud.io` |
| Kind               | `PostgresqlCluster`        | `PostgresqlCluster`          |
| Terraform resource | `ionoscloud_pg_cluster`    | `ionoscloud_pg_cluster_v2`   |

Migrating on the IONOS Cloud side is **in place**: an existing cluster keeps
its cluster ID, its data, and its DNS name — it simply becomes reachable
through the v2 API. Nothing about your running database changes.

What *does* change is how you describe the cluster to Crossplane. Since v1
and v2 are different CRDs (different API group, different underlying
Terraform resource), Crossplane has no built-in conversion between them. You
migrate by:

1. Telling Crossplane to stop deleting the real cluster if the old `v1`
   `PostgresqlCluster` custom resource (CR) is removed (`orphan` the
   resource).
2. Creating a new `v2` `PostgresqlCluster` CR whose fields describe the same
   cluster, with `crossplane.io/external-name` set to the existing cluster
   ID so Crossplane *adopts* the cluster instead of provisioning a new one.
3. Verifying the new CR reconciles against the existing cluster.
4. Deleting the old, now-orphaned `v1` CR.

This provider has shipped the `postgresqlv2.ionoscloud.io` CRDs since
provider release `v0.5.4` (Terraform provider `v6.7.29`). Make sure you're
running at least that version before starting.

## Known limitations

v1 exposes multiple users and databases per cluster through separate CRDs:
`PostgresqlUser` and `PostgresqlDatabase` (in the `postgresql.ionoscloud.io`
group). **v2 has no equivalent CRDs.** The v2 `PostgresqlCluster` spec only
supports a single inline user/database pair, under `credentials.username`
and `credentials.database`.

If you have additional `PostgresqlUser` or `PostgresqlDatabase` resources
beyond the cluster's initial user/database, there is currently no way to
represent them under the v2 CRD. Options:

- Leave those `PostgresqlUser`/`PostgresqlDatabase` v1 resources in place,
  pointed at the (now v2-managed) cluster via `clusterId`/`clusterIdRef` —
  they will keep reconciling against the same cluster, since the cluster ID
  doesn't change.
- Manage additional users/databases directly with `psql` or another
  PostgreSQL client outside of Crossplane, until a v2 equivalent CRD exists.

## Field mapping

The shape of `spec.forProvider` changed substantially. Fields not listed
below (e.g. `location`/`locationSelector`, `maintenanceWindow`) are
unchanged.

| v1 field                                                 | v2 field                                   | Notes                                                                                                                  |
|----------------------------------------------------------|--------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| `displayName`                                            | `name`                                     | Renamed.                                                                                                               |
| `postgresVersion`                                        | `version`                                  | Renamed.                                                                                                               |
| `synchronizationMode`                                    | `replicationMode`                          | Renamed, same values (`ASYNCHRONOUS`, `STRICTLY_SYNCHRONOUS`).                                                         |
| `cores`                                                  | `instances.cores`                          | Moved under the new `instances` block.                                                                                 |
| `ram` (MB)                                               | `instances.ram` (GB)                       | Moved under `instances`, **unit changed from MB to GB**.                                                               |
| `storageSize` (MB)                                       | `instances.storageSize` (GB)               | Moved under `instances`, **unit changed from MB to GB**.                                                               |
| `instances` (count)                                      | `instances.count`                          | Moved under `instances` alongside `cores`/`ram`/`storageSize`.                                                         |
| `storageType`                                            | *(removed)*                                | No v2 equivalent; storage type is no longer selectable.                                                                |
| `credentials` (list, one entry)                          | `credentials` (single object)              | v1 accepted a list but only the first entry is used; v2 makes this a single object.                                    |
| `credentials[].username`                                 | `credentials.username`                     | Same meaning.                                                                                                          |
| `credentials[].passwordSecretRef`                        | `credentials.passwordSecretRef`            | Same meaning.                                                                                                          |
| *(none — required separate `PostgresqlDatabase`)*        | `credentials.database`                     | v2 lets you name the initial database inline.                                                                          |
| *(none)*                                                 | `credentials.passwordVersion`              | New. Bump this string to trigger a password update; the password itself is write-only in v2.                           |
| `connections.cidr`                                       | `connections.primaryInstanceAddress`       | The IP/subnet addressing model changed: v1 took a CIDR for the whole cluster, v2 takes the primary instance's address. |
| `connections.datacenterId(Ref/Selector)`                 | `connections.datacenterId(Ref/Selector)`   | Unchanged.                                                                                                             |
| `connections.lanId(Ref/Selector)`                        | `connections.lanId(Ref/Selector)`          | Unchanged.                                                                                                             |
| `backupLocation`                                         | `backup.location`                          | Moved under a new `backup` block.                                                                                      |
| *(none)*                                                 | `backup.retentionDays`                     | New.                                                                                                                   |
| `connectionPooler.enabled` / `connectionPooler.poolMode` | `connectionPooler`                         | Changed from an object to a single string: `DISABLED`, `TRANSACTION`, or `SESSION`.                                    |
| `fromBackup.backupId`                                    | `restoreFromBackup.sourceBackupId`         | Renamed block and field.                                                                                               |
| `fromBackup.recoveryTargetTime`                          | `restoreFromBackup.recoveryTargetDatetime` | Renamed.                                                                                                               |
| `allowReplace`                                           | *(removed)*                                | No v2 equivalent.                                                                                                      |

v2 also has different required fields: `backup`, `connections`,
`credentials`, `instances`, `maintenanceWindow`, `name`, `replicationMode`,
and `version` are all required (compare with v1's `cores`, `credentials`,
`displayName`, `instances`, `postgresVersion`, `ram`, `storageSize`,
`storageType`, `synchronizationMode`).

## Step-by-step walkthrough

### 1. Find the cluster ID of the existing v1 resource

```shell
kubectl get postgresqlcluster.postgresql.ionoscloud.io example \
  -o jsonpath='{.metadata.annotations.crossplane\.io/external-name}'
```

```
7cf6b0b3-3edb-4e78-a039-4c5cef3d81ac
```

Record this ID — it's the actual IONOS Cloud cluster ID and will become the
`external-name` of the new v2 resource.

### 2. Orphan the old v1 resource

Set `spec.deletionPolicy: Orphan` on the v1 `PostgresqlCluster` so deleting
the CR later does **not** delete the real cluster:

```shell
kubectl patch postgresqlcluster.postgresql.ionoscloud.io example \
  --type=merge -p '{"spec":{"deletionPolicy":"Orphan"}}'
```

### 3. Author the new v2 resource

There are two ways to build `spec.forProvider` for the new resource: translate
it by hand using the field mapping table (Option A), or let the provider read
the live cluster and generate it for you (Option B). Option B avoids manual
translation mistakes (e.g. the MB→GB unit change) and is the recommended
approach; Option A is documented for reference and for when you want full
control over the resulting manifest from the start.

#### Option A: Translate the spec by hand

Translate the old `spec.forProvider` using the field mapping table above.
Set `crossplane.io/external-name` to the cluster ID from step 1 so
Crossplane adopts the existing cluster instead of creating a new one.

Given a v1 resource like:

```yaml
apiVersion: postgresql.ionoscloud.io/v1alpha1
kind: PostgresqlCluster
metadata:
  name: example
spec:
  forProvider:
    displayName: PostgreSQL_cluster
    postgresVersion: "12"
    synchronizationMode: ASYNCHRONOUS
    cores: 1
    ram: 2048
    storageSize: 2048
    storageType: HDD
    instances: 1
    credentials:
      - username: username
        passwordSecretRef:
          name: pwsecret
          namespace: upbound-system
          key: password
    connections:
      cidr: 192.168.100.1/24
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
apiVersion: postgresqlv2.ionoscloud.io/v1alpha1
kind: PostgresqlCluster
metadata:
  name: example
  annotations:
    crossplane.io/external-name: "7cf6b0b3-3edb-4e78-a039-4c5cef3d81ac" # cluster ID from step 1
spec:
  forProvider:
    name: PostgreSQL_cluster
    version: "12"
    replicationMode: ASYNCHRONOUS
    instances:
      cores: 1
      ram: 2 # MB -> GB
      storageSize: 2 # MB -> GB
      count: 1
    credentials:
      username: username
      passwordSecretRef:
        name: pwsecret
        namespace: upbound-system
        key: password
      passwordVersion: "1"
    connections:
      primaryInstanceAddress: 192.168.100.1/24
      datacenterIdSelector:
        matchLabels:
          testing.upbound.io/example-name: datacenter_example
      lanIdSelector:
        matchLabels:
          testing.upbound.io/example-name: lan_example
    backup:
      location: eu-central-3 # required in v2; see IONOS Cloud docs for your cluster's backup location
      retentionDays: 7
    maintenanceWindow:
      dayOfTheWeek: Sunday
      time: "09:00:00"
    locationSelector:
      matchLabels:
        testing.upbound.io/example-name: datacenter_example
```

> **Note:** `storageType`, `allowReplace`, and multiple `credentials` entries
> have no v2 equivalent and are dropped. `backup.location` is required in
> v2 even though `backupLocation` was optional/computed in v1 — check your
> cluster's actual backup location in the IONOS Cloud console or via the
> `ionoscloud_pg_backup_location_v2` data source before setting it, since an
> incorrect value can force a new cluster.

Apply the new resource:

```shell
kubectl apply -f postgresqlcluster-v2.yaml
```

#### Option B: Import with an Observe-only management policy (recommended)

Instead of hand-translating every field, apply a v2 resource with an
`Observe`-only management policy and an empty (or minimal)
`spec.forProvider`. The provider will read the live cluster and populate
`status.atProvider` for you — in the *v2* field shape — which you then copy
into `spec.forProvider` verbatim.

1. Apply a minimal v2 resource, with `crossplane.io/external-name` set to the
   cluster ID from step 1 and `managementPolicies: ["Observe"]`:

   ```yaml
   apiVersion: postgresqlv2.ionoscloud.io/v1alpha1
   kind: PostgresqlCluster
   metadata:
     name: example
     annotations:
       crossplane.io/external-name: "7cf6b0b3-3edb-4e78-a039-4c5cef3d81ac" # cluster ID from step 1
   spec:
     managementPolicies: ["Observe"]
     forProvider: {}
   ```

   `Observe` skips the usual required-field validation on `forProvider`, so
   this applies cleanly even though `forProvider` is empty.

   ```shell
   kubectl apply -f postgresqlcluster-v2-observe.yaml
   ```

2. Wait for the resource to become `SYNCED`/`READY`, then read the observed
   state. `status.atProvider.id` and `status.atProvider.dnsName` are
   observation-only — there is no matching `id`/`dnsName` field under
   `spec.forProvider`, so pasting them in as-is will be rejected by the CRD
   schema. Strip them with `jq` while extracting, so what you get back is
   already safe to paste into `spec.forProvider`:

   ```shell
   kubectl get postgresqlcluster.postgresqlv2.ionoscloud.io example -o json \
     | jq '.status.atProvider | del(.id, .dnsName)'
   ```

   The result is already in the v2 field shape (`instances.cores`,
   `instances.ram` in GB, `replicationMode`, `backup.location`, etc.) — no
   manual unit conversion or field-mapping lookups required.

    **Note**: for this method, all ids (like datacenterID) will be returned as a UUID, not using a selector or reference. 
If you want to use a selector or reference, you will need to manually set that in the spec.

3. Copy the `jq` output into `spec.forProvider`. `status.atProvider` never
   contains write-only fields, so you still need to set
   `credentials.username`, `credentials.database`,
   `credentials.passwordSecretRef`, and `credentials.passwordVersion`
   yourself (from what you already know about the cluster) — these aren't
   observable and won't appear under `status.atProvider` regardless of the
   `jq` filter.

4. Switch the resource back to full management by removing
   `managementPolicies` (or setting it to `["*"]`) and apply again:

   ```shell
   kubectl patch postgresqlcluster.postgresqlv2.ionoscloud.io example \
     --type=merge -p '{"spec":{"managementPolicies":null}}'
   ```

   Then re-apply the manifest with the populated `spec.forProvider` from
   step 3.

5. Confirm there's no drift: `SYNCED`/`READY` should stay `True` with no
   further update or replace triggered, since `spec.forProvider` now matches
   `status.atProvider` exactly.

### 4. Verify adoption

Confirm the new CR reconciles against the *same* cluster ID rather than
creating a new one:

```shell
kubectl get postgresqlcluster.postgresqlv2.ionoscloud.io example
```

```
NAME      SYNCED   READY   EXTERNAL-NAME                          AGE
example   True     True    7cf6b0b3-3edb-4e78-a039-4c5cef3d81ac   30s
```

The `EXTERNAL-NAME` must match the cluster ID recorded in step 1. If it
doesn't, Crossplane created a brand-new cluster instead of adopting the
existing one — delete the new CR (with `deletionPolicy: Delete`, the
default) before it provisions further, and check the
`crossplane.io/external-name` annotation on the manifest you applied.

Also check for a clean diff between the applied spec and
`status.atProvider` — any drift here usually means a field was mistranslated
in step 3 and will trigger an update (or a replace, for immutable fields)
against the real cluster.

### 5. Remove the old v1 resource

Once the v2 resource is `SYNCED`/`READY` and matches the expected cluster:

**DO NOT** run `kubectl delete` on the old v1 resource unless you set `spec.deletionPolicy: Orphan` in step 2. 
If you delete the v1 resource without orphaning it first, Crossplane will delete the real cluster.

```shell
kubectl delete postgresqlcluster.postgresql.ionoscloud.io example
```

Because it was orphaned in step 2, this only removes the Kubernetes object —
the underlying cluster keeps running, now managed solely by the v2 resource.

## Troubleshooting

- **New cluster created instead of adopted**: the
  `crossplane.io/external-name` annotation was missing or didn't match the
  real cluster ID exactly. Delete the erroneous CR and retry with the
  correct ID.
- **Update/replace triggered immediately after adoption**: a field in the
  new v2 spec doesn't match the cluster's actual current state (e.g. wrong
  `backup.location`, unit mix-up between MB and GB on `instances.ram`/
  `instances.storageSize`). Compare `status.atProvider` on the new resource
  against the old one's `status.atProvider` and correct the mismatched
  field.
- **Old v1 resource deleted the real cluster**: this means
  `deletionPolicy: Orphan` wasn't set (or wasn't applied) before deletion.
  There is no way to recover a deleted cluster through Crossplane — restore
  from a backup via `restoreFromBackup` if one exists.
