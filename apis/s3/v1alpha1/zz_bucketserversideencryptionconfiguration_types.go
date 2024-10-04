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

type ApplyServerSideEncryptionByDefaultInitParameters struct {

	// [string] Server-side encryption algorithm to use. Valid values are 'AES256'
	// Server-side encryption algorithm to use. Valid values are 'AES256'
	SseAlgorithm *string `json:"sseAlgorithm,omitempty" tf:"sse_algorithm,omitempty"`
}

type ApplyServerSideEncryptionByDefaultObservation struct {

	// [string] Server-side encryption algorithm to use. Valid values are 'AES256'
	// Server-side encryption algorithm to use. Valid values are 'AES256'
	SseAlgorithm *string `json:"sseAlgorithm,omitempty" tf:"sse_algorithm,omitempty"`
}

type ApplyServerSideEncryptionByDefaultParameters struct {

	// [string] Server-side encryption algorithm to use. Valid values are 'AES256'
	// Server-side encryption algorithm to use. Valid values are 'AES256'
	// +kubebuilder:validation:Optional
	SseAlgorithm *string `json:"sseAlgorithm" tf:"sse_algorithm,omitempty"`
}

type BucketServerSideEncryptionConfigurationInitParameters struct {

	// [string] The name of the bucket where the object will be stored.
	// The name of the bucket
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/s3/v1alpha1.Bucket
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// [block] A block of rule as defined below.
	// Specifies the default server-side encryption configuration.
	Rule []BucketServerSideEncryptionConfigurationRuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketServerSideEncryptionConfigurationObservation struct {

	// [string] The name of the bucket where the object will be stored.
	// The name of the bucket
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [block] A block of rule as defined below.
	// Specifies the default server-side encryption configuration.
	Rule []BucketServerSideEncryptionConfigurationRuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketServerSideEncryptionConfigurationParameters struct {

	// [string] The name of the bucket where the object will be stored.
	// The name of the bucket
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/s3/v1alpha1.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// [block] A block of rule as defined below.
	// Specifies the default server-side encryption configuration.
	// +kubebuilder:validation:Optional
	Rule []BucketServerSideEncryptionConfigurationRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketServerSideEncryptionConfigurationRuleInitParameters struct {

	// [block] Defines the default encryption settings.
	// Defines the default encryption settings.
	ApplyServerSideEncryptionByDefault *ApplyServerSideEncryptionByDefaultInitParameters `json:"applyServerSideEncryptionByDefault,omitempty" tf:"apply_server_side_encryption_by_default,omitempty"`
}

type BucketServerSideEncryptionConfigurationRuleObservation struct {

	// [block] Defines the default encryption settings.
	// Defines the default encryption settings.
	ApplyServerSideEncryptionByDefault *ApplyServerSideEncryptionByDefaultObservation `json:"applyServerSideEncryptionByDefault,omitempty" tf:"apply_server_side_encryption_by_default,omitempty"`
}

type BucketServerSideEncryptionConfigurationRuleParameters struct {

	// [block] Defines the default encryption settings.
	// Defines the default encryption settings.
	// +kubebuilder:validation:Optional
	ApplyServerSideEncryptionByDefault *ApplyServerSideEncryptionByDefaultParameters `json:"applyServerSideEncryptionByDefault" tf:"apply_server_side_encryption_by_default,omitempty"`
}

// BucketServerSideEncryptionConfigurationSpec defines the desired state of BucketServerSideEncryptionConfiguration
type BucketServerSideEncryptionConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketServerSideEncryptionConfigurationParameters `json:"forProvider"`
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
	InitProvider BucketServerSideEncryptionConfigurationInitParameters `json:"initProvider,omitempty"`
}

// BucketServerSideEncryptionConfigurationStatus defines the observed state of BucketServerSideEncryptionConfiguration.
type BucketServerSideEncryptionConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketServerSideEncryptionConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BucketServerSideEncryptionConfiguration is the Schema for the BucketServerSideEncryptionConfigurations API. Manages Buckets server side encryption configuration on IonosCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type BucketServerSideEncryptionConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketServerSideEncryptionConfigurationSpec   `json:"spec"`
	Status            BucketServerSideEncryptionConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketServerSideEncryptionConfigurationList contains a list of BucketServerSideEncryptionConfigurations
type BucketServerSideEncryptionConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketServerSideEncryptionConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketServerSideEncryptionConfiguration_Kind             = "BucketServerSideEncryptionConfiguration"
	BucketServerSideEncryptionConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketServerSideEncryptionConfiguration_Kind}.String()
	BucketServerSideEncryptionConfiguration_KindAPIVersion   = BucketServerSideEncryptionConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketServerSideEncryptionConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketServerSideEncryptionConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketServerSideEncryptionConfiguration{}, &BucketServerSideEncryptionConfigurationList{})
}
