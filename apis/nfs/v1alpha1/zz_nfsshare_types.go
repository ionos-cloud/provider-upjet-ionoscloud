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

type ClientGroupsInitParameters struct {

	// Optional description for the clients groups.
	// Optional description for the clients groups.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A singular host allowed to connect to the share. The host can be specified as IP address and can be either IPv4 or IPv6.
	// A singular host allowed to connect to the share. The host can be specified as IP address and can be either IPv4 or IPv6.
	Hosts []*string `json:"hosts,omitempty" tf:"hosts,omitempty"`

	// The allowed host or network to which the export is being shared. The IP address can be either IPv4 or IPv6 and has to be given with CIDR notation.
	// The allowed host or network to which the export is being shared. The IP address can be either IPv4 or IPv6 and has to be given with CIDR notation.
	IPNetworks []*string `json:"ipNetworks,omitempty" tf:"ip_networks,omitempty"`

	// NFS specific configurations. Each configuration includes:
	NFS []ClientGroupsNFSInitParameters `json:"nfs,omitempty" tf:"nfs,omitempty"`
}

type ClientGroupsNFSInitParameters struct {

	// The squash mode for the export. The squash mode can be:
	// The squash mode for the export. The squash mode can be: none - No squash mode. no mapping, root-anonymous - Map root user to anonymous uid, all-anonymous - Map all users to anonymous uid.
	Squash *string `json:"squash,omitempty" tf:"squash,omitempty"`
}

type ClientGroupsNFSObservation struct {

	// The squash mode for the export. The squash mode can be:
	// The squash mode for the export. The squash mode can be: none - No squash mode. no mapping, root-anonymous - Map root user to anonymous uid, all-anonymous - Map all users to anonymous uid.
	Squash *string `json:"squash,omitempty" tf:"squash,omitempty"`
}

type ClientGroupsNFSParameters struct {

	// The squash mode for the export. The squash mode can be:
	// The squash mode for the export. The squash mode can be: none - No squash mode. no mapping, root-anonymous - Map root user to anonymous uid, all-anonymous - Map all users to anonymous uid.
	// +kubebuilder:validation:Optional
	Squash *string `json:"squash,omitempty" tf:"squash,omitempty"`
}

type ClientGroupsObservation struct {

	// Optional description for the clients groups.
	// Optional description for the clients groups.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A singular host allowed to connect to the share. The host can be specified as IP address and can be either IPv4 or IPv6.
	// A singular host allowed to connect to the share. The host can be specified as IP address and can be either IPv4 or IPv6.
	Hosts []*string `json:"hosts,omitempty" tf:"hosts,omitempty"`

	// The allowed host or network to which the export is being shared. The IP address can be either IPv4 or IPv6 and has to be given with CIDR notation.
	// The allowed host or network to which the export is being shared. The IP address can be either IPv4 or IPv6 and has to be given with CIDR notation.
	IPNetworks []*string `json:"ipNetworks,omitempty" tf:"ip_networks,omitempty"`

	// NFS specific configurations. Each configuration includes:
	NFS []ClientGroupsNFSObservation `json:"nfs,omitempty" tf:"nfs,omitempty"`
}

type ClientGroupsParameters struct {

	// Optional description for the clients groups.
	// Optional description for the clients groups.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A singular host allowed to connect to the share. The host can be specified as IP address and can be either IPv4 or IPv6.
	// A singular host allowed to connect to the share. The host can be specified as IP address and can be either IPv4 or IPv6.
	// +kubebuilder:validation:Optional
	Hosts []*string `json:"hosts" tf:"hosts,omitempty"`

	// The allowed host or network to which the export is being shared. The IP address can be either IPv4 or IPv6 and has to be given with CIDR notation.
	// The allowed host or network to which the export is being shared. The IP address can be either IPv4 or IPv6 and has to be given with CIDR notation.
	// +kubebuilder:validation:Optional
	IPNetworks []*string `json:"ipNetworks" tf:"ip_networks,omitempty"`

	// NFS specific configurations. Each configuration includes:
	// +kubebuilder:validation:Optional
	NFS []ClientGroupsNFSParameters `json:"nfs,omitempty" tf:"nfs,omitempty"`
}

type NFSShareInitParameters struct {

	// The groups of clients are the systems connecting to the Network File Storage cluster. Each group includes:
	// The groups of clients are the systems connecting to the Network File Storage cluster.
	ClientGroups []ClientGroupsInitParameters `json:"clientGroups,omitempty" tf:"client_groups,omitempty"`

	// The ID of the Network File Storage Cluster.
	// The ID of the Network File Storage Cluster.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/nfs/v1alpha1.NFSCluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a NFSCluster in nfs to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a NFSCluster in nfs to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// The group ID that will own the exported directory. If not set, anonymous (512) will be used.
	// The group ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	GID *float64 `json:"gid,omitempty" tf:"gid,omitempty"`

	// The location of the Network File Storage Cluster.
	// The location of the Network File Storage Cluster. Available locations: 'de/fra, 'de/txl'
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The directory being exported.
	// The directory being exported
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The quota in MiB for the export. The quota can restrict the amount of data that can be stored within the export. The quota can be disabled using 0. Default is 0.
	// The quota in MiB for the export. The quota can restrict the amount of data that can be stored within the export. The quota can be disabled using `0`.
	Quota *float64 `json:"quota,omitempty" tf:"quota,omitempty"`

	// The user ID that will own the exported directory. If not set, anonymous (512) will be used.
	// The user ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	UID *float64 `json:"uid,omitempty" tf:"uid,omitempty"`
}

type NFSShareObservation struct {

	// The groups of clients are the systems connecting to the Network File Storage cluster. Each group includes:
	// The groups of clients are the systems connecting to the Network File Storage cluster.
	ClientGroups []ClientGroupsObservation `json:"clientGroups,omitempty" tf:"client_groups,omitempty"`

	// The ID of the Network File Storage Cluster.
	// The ID of the Network File Storage Cluster.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// The group ID that will own the exported directory. If not set, anonymous (512) will be used.
	// The group ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	GID *float64 `json:"gid,omitempty" tf:"gid,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The location of the Network File Storage Cluster.
	// The location of the Network File Storage Cluster. Available locations: 'de/fra, 'de/txl'
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Path to the NFS export. The NFS path is the path to the directory being exported.
	NFSPath *string `json:"nfsPath,omitempty" tf:"nfs_path,omitempty"`

	// The directory being exported.
	// The directory being exported
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The quota in MiB for the export. The quota can restrict the amount of data that can be stored within the export. The quota can be disabled using 0. Default is 0.
	// The quota in MiB for the export. The quota can restrict the amount of data that can be stored within the export. The quota can be disabled using `0`.
	Quota *float64 `json:"quota,omitempty" tf:"quota,omitempty"`

	// The user ID that will own the exported directory. If not set, anonymous (512) will be used.
	// The user ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	UID *float64 `json:"uid,omitempty" tf:"uid,omitempty"`
}

type NFSShareParameters struct {

	// The groups of clients are the systems connecting to the Network File Storage cluster. Each group includes:
	// The groups of clients are the systems connecting to the Network File Storage cluster.
	// +kubebuilder:validation:Optional
	ClientGroups []ClientGroupsParameters `json:"clientGroups,omitempty" tf:"client_groups,omitempty"`

	// The ID of the Network File Storage Cluster.
	// The ID of the Network File Storage Cluster.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/nfs/v1alpha1.NFSCluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a NFSCluster in nfs to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a NFSCluster in nfs to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// The group ID that will own the exported directory. If not set, anonymous (512) will be used.
	// The group ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	// +kubebuilder:validation:Optional
	GID *float64 `json:"gid,omitempty" tf:"gid,omitempty"`

	// The location of the Network File Storage Cluster.
	// The location of the Network File Storage Cluster. Available locations: 'de/fra, 'de/txl'
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The directory being exported.
	// The directory being exported
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The quota in MiB for the export. The quota can restrict the amount of data that can be stored within the export. The quota can be disabled using 0. Default is 0.
	// The quota in MiB for the export. The quota can restrict the amount of data that can be stored within the export. The quota can be disabled using `0`.
	// +kubebuilder:validation:Optional
	Quota *float64 `json:"quota,omitempty" tf:"quota,omitempty"`

	// The user ID that will own the exported directory. If not set, anonymous (512) will be used.
	// The user ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	// +kubebuilder:validation:Optional
	UID *float64 `json:"uid,omitempty" tf:"uid,omitempty"`
}

// NFSShareSpec defines the desired state of NFSShare
type NFSShareSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NFSShareParameters `json:"forProvider"`
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
	InitProvider NFSShareInitParameters `json:"initProvider,omitempty"`
}

// NFSShareStatus defines the observed state of NFSShare.
type NFSShareStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NFSShareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NFSShare is the Schema for the NFSShares API. Creates and manages Network File Storage (NFS) Share objects on IonosCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type NFSShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.clientGroups) || (has(self.initProvider) && has(self.initProvider.clientGroups))",message="spec.forProvider.clientGroups is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   NFSShareSpec   `json:"spec"`
	Status NFSShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NFSShareList contains a list of NFSShares
type NFSShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NFSShare `json:"items"`
}

// Repository type metadata.
var (
	NFSShare_Kind             = "NFSShare"
	NFSShare_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NFSShare_Kind}.String()
	NFSShare_KindAPIVersion   = NFSShare_Kind + "." + CRDGroupVersion.String()
	NFSShare_GroupVersionKind = CRDGroupVersion.WithKind(NFSShare_Kind)
)

func init() {
	SchemeBuilder.Register(&NFSShare{}, &NFSShareList{})
}
