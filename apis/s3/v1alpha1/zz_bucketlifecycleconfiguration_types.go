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

type AbortIncompleteMultipartUploadInitParameters struct {

	// [int] Specifies the number of days after which IONOS S3 Object Storage aborts an incomplete multipart upload.
	// Specifies the number of days after which IONOS S3 Object Storage aborts an incomplete multipart upload.
	DaysAfterInitiation *float64 `json:"daysAfterInitiation,omitempty" tf:"days_after_initiation,omitempty"`
}

type AbortIncompleteMultipartUploadObservation struct {

	// [int] Specifies the number of days after which IONOS S3 Object Storage aborts an incomplete multipart upload.
	// Specifies the number of days after which IONOS S3 Object Storage aborts an incomplete multipart upload.
	DaysAfterInitiation *float64 `json:"daysAfterInitiation,omitempty" tf:"days_after_initiation,omitempty"`
}

type AbortIncompleteMultipartUploadParameters struct {

	// [int] Specifies the number of days after which IONOS S3 Object Storage aborts an incomplete multipart upload.
	// Specifies the number of days after which IONOS S3 Object Storage aborts an incomplete multipart upload.
	// +kubebuilder:validation:Optional
	DaysAfterInitiation *float64 `json:"daysAfterInitiation" tf:"days_after_initiation,omitempty"`
}

type BucketLifecycleConfigurationInitParameters struct {

	// [string] The name of the bucket where the object will be stored.
	// The name of the S3 bucket.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/s3/v1alpha1.Bucket
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// A list of lifecycle rules for objects in the bucket.
	Rule []RuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketLifecycleConfigurationObservation struct {

	// [string] The name of the bucket where the object will be stored.
	// The name of the S3 bucket.
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// [int] Container for the Contract Number of the owner
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of lifecycle rules for objects in the bucket.
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type BucketLifecycleConfigurationParameters struct {

	// [string] The name of the bucket where the object will be stored.
	// The name of the S3 bucket.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/s3/v1alpha1.Bucket
	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// A list of lifecycle rules for objects in the bucket.
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type ExpirationInitParameters struct {

	// [string] Specifies the date after which you want the specific rule action to take effect.
	// Specifies the date when the object expires. Required if 'days' is not specified.
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// [int] Specifies the number of days after object creation when the object expires. Required if 'date' is not specified.
	// Specifies the number of days after object creation when the object expires. Required if 'date' is not specified.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// [bool] Indicates whether IONOS S3 Object Storage will remove a delete marker with no noncurrent versions. If set to true, the delete marker will be expired; if set to false the policy takes no operation. This cannot be specified with Days or Date in a Lifecycle Expiration Policy.
	// Indicates whether IONOS S3 Object Storage will remove a delete marker with no noncurrent versions. If set to true, the delete marker will be expired; if set to false the policy takes no operation. This cannot be specified with Days or Date in a Lifecycle Expiration Policy.
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type ExpirationObservation struct {

	// [string] Specifies the date after which you want the specific rule action to take effect.
	// Specifies the date when the object expires. Required if 'days' is not specified.
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// [int] Specifies the number of days after object creation when the object expires. Required if 'date' is not specified.
	// Specifies the number of days after object creation when the object expires. Required if 'date' is not specified.
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// [bool] Indicates whether IONOS S3 Object Storage will remove a delete marker with no noncurrent versions. If set to true, the delete marker will be expired; if set to false the policy takes no operation. This cannot be specified with Days or Date in a Lifecycle Expiration Policy.
	// Indicates whether IONOS S3 Object Storage will remove a delete marker with no noncurrent versions. If set to true, the delete marker will be expired; if set to false the policy takes no operation. This cannot be specified with Days or Date in a Lifecycle Expiration Policy.
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type ExpirationParameters struct {

	// [string] Specifies the date after which you want the specific rule action to take effect.
	// Specifies the date when the object expires. Required if 'days' is not specified.
	// +kubebuilder:validation:Optional
	Date *string `json:"date,omitempty" tf:"date,omitempty"`

	// [int] Specifies the number of days after object creation when the object expires. Required if 'date' is not specified.
	// Specifies the number of days after object creation when the object expires. Required if 'date' is not specified.
	// +kubebuilder:validation:Optional
	Days *float64 `json:"days,omitempty" tf:"days,omitempty"`

	// [bool] Indicates whether IONOS S3 Object Storage will remove a delete marker with no noncurrent versions. If set to true, the delete marker will be expired; if set to false the policy takes no operation. This cannot be specified with Days or Date in a Lifecycle Expiration Policy.
	// Indicates whether IONOS S3 Object Storage will remove a delete marker with no noncurrent versions. If set to true, the delete marker will be expired; if set to false the policy takes no operation. This cannot be specified with Days or Date in a Lifecycle Expiration Policy.
	// +kubebuilder:validation:Optional
	ExpiredObjectDeleteMarker *bool `json:"expiredObjectDeleteMarker,omitempty" tf:"expired_object_delete_marker,omitempty"`
}

type NoncurrentVersionExpirationInitParameters struct {

	// [int] Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action.
	// Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action.
	NoncurrentDays *float64 `json:"noncurrentDays,omitempty" tf:"noncurrent_days,omitempty"`
}

type NoncurrentVersionExpirationObservation struct {

	// [int] Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action.
	// Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action.
	NoncurrentDays *float64 `json:"noncurrentDays,omitempty" tf:"noncurrent_days,omitempty"`
}

type NoncurrentVersionExpirationParameters struct {

	// [int] Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action.
	// Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action.
	// +kubebuilder:validation:Optional
	NoncurrentDays *float64 `json:"noncurrentDays" tf:"noncurrent_days,omitempty"`
}

type RuleInitParameters struct {

	// [block] Specifies the days since the initiation of an incomplete multipart upload that IONOS S3 Object Storage will wait before permanently removing all parts of the upload.
	// Specifies the days since the initiation of an incomplete multipart upload that IONOS S3 Object Storage will wait before permanently removing all parts of the upload.
	AbortIncompleteMultipartUpload *AbortIncompleteMultipartUploadInitParameters `json:"abortIncompleteMultipartUpload,omitempty" tf:"abort_incomplete_multipart_upload,omitempty"`

	// [block]  A lifecycle rule for when an object expires.
	// A lifecycle rule for when an object expires.
	Expiration *ExpirationInitParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// [int] Container for the Contract Number of the owner
	// Unique identifier for the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [block] A lifecycle rule for when non-current object versions expire.
	// A lifecycle rule for when non-current object versions expire.
	NoncurrentVersionExpiration *NoncurrentVersionExpirationInitParameters `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// [string] Prefix identifying one or more objects to which the rule applies.
	// Object key prefix identifying one or more objects to which the rule applies.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// [string] The lifecycle rule status. Valid values are Enabled or Disabled.
	// Whether the rule is currently being applied. Valid values: Enabled or Disabled.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RuleObservation struct {

	// [block] Specifies the days since the initiation of an incomplete multipart upload that IONOS S3 Object Storage will wait before permanently removing all parts of the upload.
	// Specifies the days since the initiation of an incomplete multipart upload that IONOS S3 Object Storage will wait before permanently removing all parts of the upload.
	AbortIncompleteMultipartUpload *AbortIncompleteMultipartUploadObservation `json:"abortIncompleteMultipartUpload,omitempty" tf:"abort_incomplete_multipart_upload,omitempty"`

	// [block]  A lifecycle rule for when an object expires.
	// A lifecycle rule for when an object expires.
	Expiration *ExpirationObservation `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// [int] Container for the Contract Number of the owner
	// Unique identifier for the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [block] A lifecycle rule for when non-current object versions expire.
	// A lifecycle rule for when non-current object versions expire.
	NoncurrentVersionExpiration *NoncurrentVersionExpirationObservation `json:"noncurrentVersionExpiration,omitempty" tf:"noncurrent_version_expiration,omitempty"`

	// [string] Prefix identifying one or more objects to which the rule applies.
	// Object key prefix identifying one or more objects to which the rule applies.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// [string] The lifecycle rule status. Valid values are Enabled or Disabled.
	// Whether the rule is currently being applied. Valid values: Enabled or Disabled.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type RuleParameters struct {

	// [block] Specifies the days since the initiation of an incomplete multipart upload that IONOS S3 Object Storage will wait before permanently removing all parts of the upload.
	// Specifies the days since the initiation of an incomplete multipart upload that IONOS S3 Object Storage will wait before permanently removing all parts of the upload.
	// +kubebuilder:validation:Optional
	AbortIncompleteMultipartUpload *AbortIncompleteMultipartUploadParameters `json:"abortIncompleteMultipartUpload" tf:"abort_incomplete_multipart_upload,omitempty"`

	// [block]  A lifecycle rule for when an object expires.
	// A lifecycle rule for when an object expires.
	// +kubebuilder:validation:Optional
	Expiration *ExpirationParameters `json:"expiration,omitempty" tf:"expiration,omitempty"`

	// [int] Container for the Contract Number of the owner
	// Unique identifier for the rule.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [block] A lifecycle rule for when non-current object versions expire.
	// A lifecycle rule for when non-current object versions expire.
	// +kubebuilder:validation:Optional
	NoncurrentVersionExpiration *NoncurrentVersionExpirationParameters `json:"noncurrentVersionExpiration" tf:"noncurrent_version_expiration,omitempty"`

	// [string] Prefix identifying one or more objects to which the rule applies.
	// Object key prefix identifying one or more objects to which the rule applies.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix" tf:"prefix,omitempty"`

	// [string] The lifecycle rule status. Valid values are Enabled or Disabled.
	// Whether the rule is currently being applied. Valid values: Enabled or Disabled.
	// +kubebuilder:validation:Optional
	Status *string `json:"status" tf:"status,omitempty"`
}

// BucketLifecycleConfigurationSpec defines the desired state of BucketLifecycleConfiguration
type BucketLifecycleConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketLifecycleConfigurationParameters `json:"forProvider"`
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
	InitProvider BucketLifecycleConfigurationInitParameters `json:"initProvider,omitempty"`
}

// BucketLifecycleConfigurationStatus defines the observed state of BucketLifecycleConfiguration.
type BucketLifecycleConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketLifecycleConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BucketLifecycleConfiguration is the Schema for the BucketLifecycleConfigurations API. Manages Buckets lifecycle configuration on IonosCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type BucketLifecycleConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketLifecycleConfigurationSpec   `json:"spec"`
	Status            BucketLifecycleConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketLifecycleConfigurationList contains a list of BucketLifecycleConfigurations
type BucketLifecycleConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketLifecycleConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketLifecycleConfiguration_Kind             = "BucketLifecycleConfiguration"
	BucketLifecycleConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketLifecycleConfiguration_Kind}.String()
	BucketLifecycleConfiguration_KindAPIVersion   = BucketLifecycleConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketLifecycleConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketLifecycleConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketLifecycleConfiguration{}, &BucketLifecycleConfigurationList{})
}
