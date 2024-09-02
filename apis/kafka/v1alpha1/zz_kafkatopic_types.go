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

type KafkaTopicInitParameters struct {

	// [string] ID of the Kafka Cluster that the topic belongs to.
	// The ID of the Kafka Cluster to which the topic belongs.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/kafka/v1alpha1.Kafka
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a Kafka in kafka to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a Kafka in kafka to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// [string] The location of the Kafka Cluster Topic. Possible values: de/fra, de/txl,
	// es/vit,gb/lhr, us/ewr, us/las, us/mci, fr/par
	// The location of your Kafka Cluster Topic. Supported locations: de/fra, de/txl, es/vit, gb/lhr, us/ewr, us/las, us/mci, fr/par
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// [string] Name of the Kafka Cluster.
	// The name of your Kafka Cluster Topic. Must be 63 characters or less and must begin and end with an alphanumeric character (`[a-z0-9A-Z]`) with dashes (`-`), underscores (`_`), dots (`.`), and alphanumerics between.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [int] The number of partitions of the topic. Partitions allow for parallel
	// processing of messages. The partition count must be greater than or equal to the replication factor. Minimum value: 1.
	// Default value: 3.
	// The number of partitions of the topic. Partitions allow for parallel processing of messages. The partition count must be greater than or equal to the replication factor.
	NumberOfPartitions *float64 `json:"numberOfPartitions,omitempty" tf:"number_of_partitions,omitempty"`

	// [int] The number of replicas of the topic. The replication factor determines how many
	// copies of the topic are stored on different brokers. The replication factor must be less than or equal to the number
	// of brokers in the Kafka Cluster. Minimum value: 1. Default value: 3.
	// The number of replicas of the topic. The replication factor determines how many copies of the topic are stored on different brokers. The replication factor must be less than or equal to the number of brokers in the Kafka Cluster.
	ReplicationFactor *float64 `json:"replicationFactor,omitempty" tf:"replication_factor,omitempty"`

	// [int] This configuration controls the maximum time we will retain a log before we will
	// discard old log segments to free up space. This represents an SLA on how soon consumers must read their data. If set
	// to -1, no time limit is applied. Default value: 604800000.
	// This configuration controls the maximum time we will retain a log before we will discard old log segments to free up space. This represents an SLA on how soon consumers must read their data. If set to -1, no time limit is applied.
	RetentionTime *float64 `json:"retentionTime,omitempty" tf:"retention_time,omitempty"`

	// [int] This configuration controls the segment file size for the log. Retention and
	// cleaning is always done a file at a time so a larger segment size means fewer files but less granular control over
	// retention. Default value: 1073741824.
	// This configuration controls the segment file size for the log. Retention and cleaning is always done a file at a time so a larger segment size means fewer files but less granular control over retention.
	SegmentBytes *float64 `json:"segmentBytes,omitempty" tf:"segment_bytes,omitempty"`
}

type KafkaTopicObservation struct {

	// [string] ID of the Kafka Cluster that the topic belongs to.
	// The ID of the Kafka Cluster to which the topic belongs.
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// (Computed)[string] The UUID of the Kafka Cluster Topic.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [string] The location of the Kafka Cluster Topic. Possible values: de/fra, de/txl,
	// es/vit,gb/lhr, us/ewr, us/las, us/mci, fr/par
	// The location of your Kafka Cluster Topic. Supported locations: de/fra, de/txl, es/vit, gb/lhr, us/ewr, us/las, us/mci, fr/par
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// [string] Name of the Kafka Cluster.
	// The name of your Kafka Cluster Topic. Must be 63 characters or less and must begin and end with an alphanumeric character (`[a-z0-9A-Z]`) with dashes (`-`), underscores (`_`), dots (`.`), and alphanumerics between.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [int] The number of partitions of the topic. Partitions allow for parallel
	// processing of messages. The partition count must be greater than or equal to the replication factor. Minimum value: 1.
	// Default value: 3.
	// The number of partitions of the topic. Partitions allow for parallel processing of messages. The partition count must be greater than or equal to the replication factor.
	NumberOfPartitions *float64 `json:"numberOfPartitions,omitempty" tf:"number_of_partitions,omitempty"`

	// [int] The number of replicas of the topic. The replication factor determines how many
	// copies of the topic are stored on different brokers. The replication factor must be less than or equal to the number
	// of brokers in the Kafka Cluster. Minimum value: 1. Default value: 3.
	// The number of replicas of the topic. The replication factor determines how many copies of the topic are stored on different brokers. The replication factor must be less than or equal to the number of brokers in the Kafka Cluster.
	ReplicationFactor *float64 `json:"replicationFactor,omitempty" tf:"replication_factor,omitempty"`

	// [int] This configuration controls the maximum time we will retain a log before we will
	// discard old log segments to free up space. This represents an SLA on how soon consumers must read their data. If set
	// to -1, no time limit is applied. Default value: 604800000.
	// This configuration controls the maximum time we will retain a log before we will discard old log segments to free up space. This represents an SLA on how soon consumers must read their data. If set to -1, no time limit is applied.
	RetentionTime *float64 `json:"retentionTime,omitempty" tf:"retention_time,omitempty"`

	// [int] This configuration controls the segment file size for the log. Retention and
	// cleaning is always done a file at a time so a larger segment size means fewer files but less granular control over
	// retention. Default value: 1073741824.
	// This configuration controls the segment file size for the log. Retention and cleaning is always done a file at a time so a larger segment size means fewer files but less granular control over retention.
	SegmentBytes *float64 `json:"segmentBytes,omitempty" tf:"segment_bytes,omitempty"`
}

type KafkaTopicParameters struct {

	// [string] ID of the Kafka Cluster that the topic belongs to.
	// The ID of the Kafka Cluster to which the topic belongs.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/kafka/v1alpha1.Kafka
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a Kafka in kafka to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a Kafka in kafka to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// [string] The location of the Kafka Cluster Topic. Possible values: de/fra, de/txl,
	// es/vit,gb/lhr, us/ewr, us/las, us/mci, fr/par
	// The location of your Kafka Cluster Topic. Supported locations: de/fra, de/txl, es/vit, gb/lhr, us/ewr, us/las, us/mci, fr/par
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// [string] Name of the Kafka Cluster.
	// The name of your Kafka Cluster Topic. Must be 63 characters or less and must begin and end with an alphanumeric character (`[a-z0-9A-Z]`) with dashes (`-`), underscores (`_`), dots (`.`), and alphanumerics between.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [int] The number of partitions of the topic. Partitions allow for parallel
	// processing of messages. The partition count must be greater than or equal to the replication factor. Minimum value: 1.
	// Default value: 3.
	// The number of partitions of the topic. Partitions allow for parallel processing of messages. The partition count must be greater than or equal to the replication factor.
	// +kubebuilder:validation:Optional
	NumberOfPartitions *float64 `json:"numberOfPartitions,omitempty" tf:"number_of_partitions,omitempty"`

	// [int] The number of replicas of the topic. The replication factor determines how many
	// copies of the topic are stored on different brokers. The replication factor must be less than or equal to the number
	// of brokers in the Kafka Cluster. Minimum value: 1. Default value: 3.
	// The number of replicas of the topic. The replication factor determines how many copies of the topic are stored on different brokers. The replication factor must be less than or equal to the number of brokers in the Kafka Cluster.
	// +kubebuilder:validation:Optional
	ReplicationFactor *float64 `json:"replicationFactor,omitempty" tf:"replication_factor,omitempty"`

	// [int] This configuration controls the maximum time we will retain a log before we will
	// discard old log segments to free up space. This represents an SLA on how soon consumers must read their data. If set
	// to -1, no time limit is applied. Default value: 604800000.
	// This configuration controls the maximum time we will retain a log before we will discard old log segments to free up space. This represents an SLA on how soon consumers must read their data. If set to -1, no time limit is applied.
	// +kubebuilder:validation:Optional
	RetentionTime *float64 `json:"retentionTime,omitempty" tf:"retention_time,omitempty"`

	// [int] This configuration controls the segment file size for the log. Retention and
	// cleaning is always done a file at a time so a larger segment size means fewer files but less granular control over
	// retention. Default value: 1073741824.
	// This configuration controls the segment file size for the log. Retention and cleaning is always done a file at a time so a larger segment size means fewer files but less granular control over retention.
	// +kubebuilder:validation:Optional
	SegmentBytes *float64 `json:"segmentBytes,omitempty" tf:"segment_bytes,omitempty"`
}

// KafkaTopicSpec defines the desired state of KafkaTopic
type KafkaTopicSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     KafkaTopicParameters `json:"forProvider"`
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
	InitProvider KafkaTopicInitParameters `json:"initProvider,omitempty"`
}

// KafkaTopicStatus defines the observed state of KafkaTopic.
type KafkaTopicStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        KafkaTopicObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// KafkaTopic is the Schema for the KafkaTopics API. Creates and manages IonosCloud Kafka Cluster Topic objects.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type KafkaTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   KafkaTopicSpec   `json:"spec"`
	Status KafkaTopicStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KafkaTopicList contains a list of KafkaTopics
type KafkaTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KafkaTopic `json:"items"`
}

// Repository type metadata.
var (
	KafkaTopic_Kind             = "KafkaTopic"
	KafkaTopic_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: KafkaTopic_Kind}.String()
	KafkaTopic_KindAPIVersion   = KafkaTopic_Kind + "." + CRDGroupVersion.String()
	KafkaTopic_GroupVersionKind = CRDGroupVersion.WithKind(KafkaTopic_Kind)
)

func init() {
	SchemeBuilder.Register(&KafkaTopic{}, &KafkaTopicList{})
}
