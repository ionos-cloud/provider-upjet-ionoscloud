// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ConnectionPoolerInitParameters struct {

	// [bool]
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// [string] Represents different modes of connection pooling for the connection pooler.
	// Represents different modes of connection pooling for the connection pooler
	PoolMode *string `json:"poolMode,omitempty" tf:"pool_mode,omitempty"`
}

type ConnectionPoolerObservation struct {

	// [bool]
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// [string] Represents different modes of connection pooling for the connection pooler.
	// Represents different modes of connection pooling for the connection pooler
	PoolMode *string `json:"poolMode,omitempty" tf:"pool_mode,omitempty"`
}

type ConnectionPoolerParameters struct {

	// [bool]
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// [string] Represents different modes of connection pooling for the connection pooler.
	// Represents different modes of connection pooling for the connection pooler
	// +kubebuilder:validation:Optional
	PoolMode *string `json:"poolMode" tf:"pool_mode,omitempty"`
}

type ConnectionsInitParameters struct {

	// [true] The IP and subnet for the database. Note the following unavailable IP ranges: 10.233.64.0/18, 10.233.0.0/18, 10.233.114.0/24. Please enter in the correct format like IP/Subnet, exp: 192.168.10.0/24. See Private IPs and Configuring the network.
	// The IP and subnet for the database.
	// Note the following unavailable IP ranges:
	// 10.233.64.0/18
	// 10.233.0.0/18
	// 10.233.114.0/24
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// [true] The datacenter to connect your cluster to.
	// The datacenter to connect your cluster to.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// [true] The LAN to connect your cluster to.
	// The LAN to connect your cluster to.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Lan
	LanID *string `json:"lanId,omitempty" tf:"lan_id,omitempty"`

	// Reference to a Lan in compute to populate lanId.
	// +kubebuilder:validation:Optional
	LanIDRef *v1.Reference `json:"lanIdRef,omitempty" tf:"-"`

	// Selector for a Lan in compute to populate lanId.
	// +kubebuilder:validation:Optional
	LanIDSelector *v1.Selector `json:"lanIdSelector,omitempty" tf:"-"`
}

type ConnectionsObservation struct {

	// [true] The IP and subnet for the database. Note the following unavailable IP ranges: 10.233.64.0/18, 10.233.0.0/18, 10.233.114.0/24. Please enter in the correct format like IP/Subnet, exp: 192.168.10.0/24. See Private IPs and Configuring the network.
	// The IP and subnet for the database.
	// Note the following unavailable IP ranges:
	// 10.233.64.0/18
	// 10.233.0.0/18
	// 10.233.114.0/24
	Cidr *string `json:"cidr,omitempty" tf:"cidr,omitempty"`

	// [true] The datacenter to connect your cluster to.
	// The datacenter to connect your cluster to.
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// [true] The LAN to connect your cluster to.
	// The LAN to connect your cluster to.
	LanID *string `json:"lanId,omitempty" tf:"lan_id,omitempty"`
}

type ConnectionsParameters struct {

	// [true] The IP and subnet for the database. Note the following unavailable IP ranges: 10.233.64.0/18, 10.233.0.0/18, 10.233.114.0/24. Please enter in the correct format like IP/Subnet, exp: 192.168.10.0/24. See Private IPs and Configuring the network.
	// The IP and subnet for the database.
	// Note the following unavailable IP ranges:
	// 10.233.64.0/18
	// 10.233.0.0/18
	// 10.233.114.0/24
	// +kubebuilder:validation:Optional
	Cidr *string `json:"cidr" tf:"cidr,omitempty"`

	// [true] The datacenter to connect your cluster to.
	// The datacenter to connect your cluster to.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	// +kubebuilder:validation:Optional
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// [true] The LAN to connect your cluster to.
	// The LAN to connect your cluster to.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Lan
	// +kubebuilder:validation:Optional
	LanID *string `json:"lanId,omitempty" tf:"lan_id,omitempty"`

	// Reference to a Lan in compute to populate lanId.
	// +kubebuilder:validation:Optional
	LanIDRef *v1.Reference `json:"lanIdRef,omitempty" tf:"-"`

	// Selector for a Lan in compute to populate lanId.
	// +kubebuilder:validation:Optional
	LanIDSelector *v1.Selector `json:"lanIdSelector,omitempty" tf:"-"`
}

type CredentialsInitParameters struct {

	// [string]
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// [string] The username for the initial postgres user. Some system usernames are restricted (e.g. "postgres", "admin", "standby")
	// the username for the initial postgres user. some system usernames are restricted (e.g. "postgres", "admin", "standby")
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type CredentialsObservation struct {

	// [string] The username for the initial postgres user. Some system usernames are restricted (e.g. "postgres", "admin", "standby")
	// the username for the initial postgres user. some system usernames are restricted (e.g. "postgres", "admin", "standby")
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type CredentialsParameters struct {

	// [string]
	// +kubebuilder:validation:Optional
	PasswordSecretRef v1.SecretKeySelector `json:"passwordSecretRef" tf:"-"`

	// [string] The username for the initial postgres user. Some system usernames are restricted (e.g. "postgres", "admin", "standby")
	// the username for the initial postgres user. some system usernames are restricted (e.g. "postgres", "admin", "standby")
	// +kubebuilder:validation:Optional
	Username *string `json:"username" tf:"username,omitempty"`
}

type FromBackupInitParameters struct {

	// [string] The unique ID of the backup you want to restore.
	// The unique ID of the backup you want to restore.
	BackupID *string `json:"backupId,omitempty" tf:"backup_id,omitempty"`

	// [string] If this value is supplied as ISO 8601 timestamp, the backup will be replayed up until the given timestamp. If empty, the backup will be applied completely.
	// If this value is supplied as ISO 8601 timestamp, the backup will be replayed up until the given timestamp. If empty, the backup will be applied completely.
	RecoveryTargetTime *string `json:"recoveryTargetTime,omitempty" tf:"recovery_target_time,omitempty"`
}

type FromBackupObservation struct {

	// [string] The unique ID of the backup you want to restore.
	// The unique ID of the backup you want to restore.
	BackupID *string `json:"backupId,omitempty" tf:"backup_id,omitempty"`

	// [string] If this value is supplied as ISO 8601 timestamp, the backup will be replayed up until the given timestamp. If empty, the backup will be applied completely.
	// If this value is supplied as ISO 8601 timestamp, the backup will be replayed up until the given timestamp. If empty, the backup will be applied completely.
	RecoveryTargetTime *string `json:"recoveryTargetTime,omitempty" tf:"recovery_target_time,omitempty"`
}

type FromBackupParameters struct {

	// [string] The unique ID of the backup you want to restore.
	// The unique ID of the backup you want to restore.
	// +kubebuilder:validation:Optional
	BackupID *string `json:"backupId" tf:"backup_id,omitempty"`

	// [string] If this value is supplied as ISO 8601 timestamp, the backup will be replayed up until the given timestamp. If empty, the backup will be applied completely.
	// If this value is supplied as ISO 8601 timestamp, the backup will be replayed up until the given timestamp. If empty, the backup will be applied completely.
	// +kubebuilder:validation:Optional
	RecoveryTargetTime *string `json:"recoveryTargetTime,omitempty" tf:"recovery_target_time,omitempty"`
}

type MaintenanceWindowInitParameters struct {

	// [string]
	DayOfTheWeek *string `json:"dayOfTheWeek,omitempty" tf:"day_of_the_week,omitempty"`

	// [string]
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type MaintenanceWindowObservation struct {

	// [string]
	DayOfTheWeek *string `json:"dayOfTheWeek,omitempty" tf:"day_of_the_week,omitempty"`

	// [string]
	Time *string `json:"time,omitempty" tf:"time,omitempty"`
}

type MaintenanceWindowParameters struct {

	// [string]
	// +kubebuilder:validation:Optional
	DayOfTheWeek *string `json:"dayOfTheWeek" tf:"day_of_the_week,omitempty"`

	// [string]
	// +kubebuilder:validation:Optional
	Time *string `json:"time" tf:"time,omitempty"`
}

type PostgresqlClusterInitParameters struct {

	// (Computed)[string] The S3 location where the backups will be stored. Possible values are: de, eu-south-2, eu-central-2. This attribute is immutable (disallowed in update requests).
	// The S3 location where the backups will be stored.
	BackupLocation *string `json:"backupLocation,omitempty" tf:"backup_location,omitempty"`

	// [object]
	// Configuration options for the connection pooler
	ConnectionPooler *ConnectionPoolerInitParameters `json:"connectionPooler,omitempty" tf:"connection_pooler,omitempty"`

	// [string] Details about the network connection for your cluster.
	// Details about the network connection for your cluster.
	Connections *ConnectionsInitParameters `json:"connections,omitempty" tf:"connections,omitempty"`

	// [int] The number of CPU cores per replica.
	// The number of CPU cores per replica.
	Cores *float64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// [string] Credentials for the database user to be created. This attribute is immutable(disallowed in update requests).
	// Credentials for the database user to be created.
	Credentials []CredentialsInitParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// [string] The friendly name of your cluster.
	// The friendly name of your cluster.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// [string] The unique ID of the backup you want to restore. This attribute is immutable(disallowed in update requests).
	// Creates the cluster based on the existing backup.
	FromBackup *FromBackupInitParameters `json:"fromBackup,omitempty" tf:"from_backup,omitempty"`

	// [int] The total number of instances in the cluster (one master and n-1 standbys)
	// The total number of instances in the cluster (one master and n-1 standbys)
	Instances *float64 `json:"instances,omitempty" tf:"instances,omitempty"`

	// [string] The physical location where the cluster will be created. This will be where all of your instances live. Property cannot be modified after datacenter creation. Possible values are: de/fra, de/txl, gb/lhr, es/vit, us/ewr, us/las. This attribute is immutable(disallowed in update requests).
	// The physical location where the cluster will be created. This will be where all of your instances live. Property cannot be modified after datacenter creation (disallowed in update requests)
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	// +crossplane:generate:reference:extractor=github.com/ionos-cloud/provider-upjet-ionoscloud/config/common.DatacenterLocation()
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Reference to a Datacenter in compute to populate location.
	// +kubebuilder:validation:Optional
	LocationRef *v1.Reference `json:"locationRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate location.
	// +kubebuilder:validation:Optional
	LocationSelector *v1.Selector `json:"locationSelector,omitempty" tf:"-"`

	// (Computed)[string] A weekly 4 hour-long window, during which maintenance might occur
	// a weekly 4 hour-long window, during which maintenance might occur
	MaintenanceWindow *MaintenanceWindowInitParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] The PostgreSQL version of your cluster.
	// The PostgreSQL version of your cluster.
	PostgresVersion *string `json:"postgresVersion,omitempty" tf:"postgres_version,omitempty"`

	// [int] The amount of memory per instance in megabytes. Has to be a multiple of 1024.
	// The amount of memory per instance in megabytes. Has to be a multiple of 1024.
	RAM *float64 `json:"ram,omitempty" tf:"ram,omitempty"`

	// [int] The amount of storage per instance in MB. Has to be a multiple of 2048.
	// The amount of storage per instance in megabytes. Has to be a multiple of 2048.
	StorageSize *float64 `json:"storageSize,omitempty" tf:"storage_size,omitempty"`

	// [string] SSD, SSD Standard, SSD Premium, or HDD. Value "SSD" is deprecated, use the equivalent "SSD Premium" instead. This attribute is immutable(disallowed in update requests).
	// The storage type used in your cluster.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// [string] Represents different modes of replication. Can have one of the following values: ASYNCHRONOUS, SYNCHRONOUS, STRICTLY_SYNCHRONOUS. This attribute is immutable(disallowed in update requests).
	// Represents different modes of replication.
	SynchronizationMode *string `json:"synchronizationMode,omitempty" tf:"synchronization_mode,omitempty"`
}

type PostgresqlClusterObservation struct {

	// (Computed)[string] The S3 location where the backups will be stored. Possible values are: de, eu-south-2, eu-central-2. This attribute is immutable (disallowed in update requests).
	// The S3 location where the backups will be stored.
	BackupLocation *string `json:"backupLocation,omitempty" tf:"backup_location,omitempty"`

	// [object]
	// Configuration options for the connection pooler
	ConnectionPooler *ConnectionPoolerObservation `json:"connectionPooler,omitempty" tf:"connection_pooler,omitempty"`

	// [string] Details about the network connection for your cluster.
	// Details about the network connection for your cluster.
	Connections *ConnectionsObservation `json:"connections,omitempty" tf:"connections,omitempty"`

	// [int] The number of CPU cores per replica.
	// The number of CPU cores per replica.
	Cores *float64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// [string] Credentials for the database user to be created. This attribute is immutable(disallowed in update requests).
	// Credentials for the database user to be created.
	Credentials []CredentialsObservation `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// (Computed)[string] The DNS name pointing to your cluster.
	// The DNS name pointing to your cluster
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// [string] The friendly name of your cluster.
	// The friendly name of your cluster.
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// [string] The unique ID of the backup you want to restore. This attribute is immutable(disallowed in update requests).
	// Creates the cluster based on the existing backup.
	FromBackup *FromBackupObservation `json:"fromBackup,omitempty" tf:"from_backup,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [int] The total number of instances in the cluster (one master and n-1 standbys)
	// The total number of instances in the cluster (one master and n-1 standbys)
	Instances *float64 `json:"instances,omitempty" tf:"instances,omitempty"`

	// [string] The physical location where the cluster will be created. This will be where all of your instances live. Property cannot be modified after datacenter creation. Possible values are: de/fra, de/txl, gb/lhr, es/vit, us/ewr, us/las. This attribute is immutable(disallowed in update requests).
	// The physical location where the cluster will be created. This will be where all of your instances live. Property cannot be modified after datacenter creation (disallowed in update requests)
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// (Computed)[string] A weekly 4 hour-long window, during which maintenance might occur
	// a weekly 4 hour-long window, during which maintenance might occur
	MaintenanceWindow *MaintenanceWindowObservation `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] The PostgreSQL version of your cluster.
	// The PostgreSQL version of your cluster.
	PostgresVersion *string `json:"postgresVersion,omitempty" tf:"postgres_version,omitempty"`

	// [int] The amount of memory per instance in megabytes. Has to be a multiple of 1024.
	// The amount of memory per instance in megabytes. Has to be a multiple of 1024.
	RAM *float64 `json:"ram,omitempty" tf:"ram,omitempty"`

	// [int] The amount of storage per instance in MB. Has to be a multiple of 2048.
	// The amount of storage per instance in megabytes. Has to be a multiple of 2048.
	StorageSize *float64 `json:"storageSize,omitempty" tf:"storage_size,omitempty"`

	// [string] SSD, SSD Standard, SSD Premium, or HDD. Value "SSD" is deprecated, use the equivalent "SSD Premium" instead. This attribute is immutable(disallowed in update requests).
	// The storage type used in your cluster.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// [string] Represents different modes of replication. Can have one of the following values: ASYNCHRONOUS, SYNCHRONOUS, STRICTLY_SYNCHRONOUS. This attribute is immutable(disallowed in update requests).
	// Represents different modes of replication.
	SynchronizationMode *string `json:"synchronizationMode,omitempty" tf:"synchronization_mode,omitempty"`
}

type PostgresqlClusterParameters struct {

	// (Computed)[string] The S3 location where the backups will be stored. Possible values are: de, eu-south-2, eu-central-2. This attribute is immutable (disallowed in update requests).
	// The S3 location where the backups will be stored.
	// +kubebuilder:validation:Optional
	BackupLocation *string `json:"backupLocation,omitempty" tf:"backup_location,omitempty"`

	// [object]
	// Configuration options for the connection pooler
	// +kubebuilder:validation:Optional
	ConnectionPooler *ConnectionPoolerParameters `json:"connectionPooler,omitempty" tf:"connection_pooler,omitempty"`

	// [string] Details about the network connection for your cluster.
	// Details about the network connection for your cluster.
	// +kubebuilder:validation:Optional
	Connections *ConnectionsParameters `json:"connections,omitempty" tf:"connections,omitempty"`

	// [int] The number of CPU cores per replica.
	// The number of CPU cores per replica.
	// +kubebuilder:validation:Optional
	Cores *float64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// [string] Credentials for the database user to be created. This attribute is immutable(disallowed in update requests).
	// Credentials for the database user to be created.
	// +kubebuilder:validation:Optional
	Credentials []CredentialsParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// [string] The friendly name of your cluster.
	// The friendly name of your cluster.
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// [string] The unique ID of the backup you want to restore. This attribute is immutable(disallowed in update requests).
	// Creates the cluster based on the existing backup.
	// +kubebuilder:validation:Optional
	FromBackup *FromBackupParameters `json:"fromBackup,omitempty" tf:"from_backup,omitempty"`

	// [int] The total number of instances in the cluster (one master and n-1 standbys)
	// The total number of instances in the cluster (one master and n-1 standbys)
	// +kubebuilder:validation:Optional
	Instances *float64 `json:"instances,omitempty" tf:"instances,omitempty"`

	// [string] The physical location where the cluster will be created. This will be where all of your instances live. Property cannot be modified after datacenter creation. Possible values are: de/fra, de/txl, gb/lhr, es/vit, us/ewr, us/las. This attribute is immutable(disallowed in update requests).
	// The physical location where the cluster will be created. This will be where all of your instances live. Property cannot be modified after datacenter creation (disallowed in update requests)
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	// +crossplane:generate:reference:extractor=github.com/ionos-cloud/provider-upjet-ionoscloud/config/common.DatacenterLocation()
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Reference to a Datacenter in compute to populate location.
	// +kubebuilder:validation:Optional
	LocationRef *v1.Reference `json:"locationRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate location.
	// +kubebuilder:validation:Optional
	LocationSelector *v1.Selector `json:"locationSelector,omitempty" tf:"-"`

	// (Computed)[string] A weekly 4 hour-long window, during which maintenance might occur
	// a weekly 4 hour-long window, during which maintenance might occur
	// +kubebuilder:validation:Optional
	MaintenanceWindow *MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// [string] The PostgreSQL version of your cluster.
	// The PostgreSQL version of your cluster.
	// +kubebuilder:validation:Optional
	PostgresVersion *string `json:"postgresVersion,omitempty" tf:"postgres_version,omitempty"`

	// [int] The amount of memory per instance in megabytes. Has to be a multiple of 1024.
	// The amount of memory per instance in megabytes. Has to be a multiple of 1024.
	// +kubebuilder:validation:Optional
	RAM *float64 `json:"ram,omitempty" tf:"ram,omitempty"`

	// [int] The amount of storage per instance in MB. Has to be a multiple of 2048.
	// The amount of storage per instance in megabytes. Has to be a multiple of 2048.
	// +kubebuilder:validation:Optional
	StorageSize *float64 `json:"storageSize,omitempty" tf:"storage_size,omitempty"`

	// [string] SSD, SSD Standard, SSD Premium, or HDD. Value "SSD" is deprecated, use the equivalent "SSD Premium" instead. This attribute is immutable(disallowed in update requests).
	// The storage type used in your cluster.
	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`

	// [string] Represents different modes of replication. Can have one of the following values: ASYNCHRONOUS, SYNCHRONOUS, STRICTLY_SYNCHRONOUS. This attribute is immutable(disallowed in update requests).
	// Represents different modes of replication.
	// +kubebuilder:validation:Optional
	SynchronizationMode *string `json:"synchronizationMode,omitempty" tf:"synchronization_mode,omitempty"`
}

// PostgresqlClusterSpec defines the desired state of PostgresqlCluster
type PostgresqlClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PostgresqlClusterParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider PostgresqlClusterInitParameters `json:"initProvider,omitempty"`
}

// PostgresqlClusterStatus defines the observed state of PostgresqlCluster.
type PostgresqlClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PostgresqlClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PostgresqlCluster is the Schema for the PostgresqlClusters API. Creates and manages DbaaS Postgres Cluster objects.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type PostgresqlCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cores) || (has(self.initProvider) && has(self.initProvider.cores))",message="spec.forProvider.cores is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.credentials) || (has(self.initProvider) && has(self.initProvider.credentials))",message="spec.forProvider.credentials is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.instances) || (has(self.initProvider) && has(self.initProvider.instances))",message="spec.forProvider.instances is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.postgresVersion) || (has(self.initProvider) && has(self.initProvider.postgresVersion))",message="spec.forProvider.postgresVersion is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ram) || (has(self.initProvider) && has(self.initProvider.ram))",message="spec.forProvider.ram is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageSize) || (has(self.initProvider) && has(self.initProvider.storageSize))",message="spec.forProvider.storageSize is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageType) || (has(self.initProvider) && has(self.initProvider.storageType))",message="spec.forProvider.storageType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.synchronizationMode) || (has(self.initProvider) && has(self.initProvider.synchronizationMode))",message="spec.forProvider.synchronizationMode is a required parameter"
	Spec   PostgresqlClusterSpec   `json:"spec"`
	Status PostgresqlClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgresqlClusterList contains a list of PostgresqlClusters
type PostgresqlClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgresqlCluster `json:"items"`
}

// Repository type metadata.
var (
	PostgresqlCluster_Kind             = "PostgresqlCluster"
	PostgresqlCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PostgresqlCluster_Kind}.String()
	PostgresqlCluster_KindAPIVersion   = PostgresqlCluster_Kind + "." + CRDGroupVersion.String()
	PostgresqlCluster_GroupVersionKind = CRDGroupVersion.WithKind(PostgresqlCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&PostgresqlCluster{}, &PostgresqlClusterList{})
}
