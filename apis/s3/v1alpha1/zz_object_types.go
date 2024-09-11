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

type ObjectInitParameters struct {

	// [string] The name of the bucket where the object will be stored. Must be between 3 and 63 characters.
	// The name of the bucket
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/s3/v1alpha1.Bucket
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// Reference to a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketRef *v1.Reference `json:"bucketRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucket.
	// +kubebuilder:validation:Optional
	BucketSelector *v1.Selector `json:"bucketSelector,omitempty" tf:"-"`

	// [string] Specifies caching behavior along the request/reply chain.
	// Can be used to specify caching behavior along the request/reply chain
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// [string] Inline content of the object.
	// The utf-8 content of the object
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// [string] Specifies presentational information for the object.
	// Specifies presentational information for the object
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// [string] Specifies what content encodings have been applied to the object.
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// [string] The natural language or languages of the intended audience for the object.
	// The natural language or languages of the intended audience for the object
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// [string] A standard MIME type describing the format of the contents.
	// A standard MIME type describing the format of the contents
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// [string] The date and time at which the object is no longer cacheable.
	// The date and time at which the object is no longer cacheable
	Expires *string `json:"expires,omitempty" tf:"expires,omitempty"`

	// [bool] If true, the object will be destroyed if versioning is enabled then all versions will be destroyed. Default is false.
	// Specifies whether to delete the object even if it has a governance-type Object Lock in place. You must explicitly pass a value of true for this parameter to delete the object.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// [string] The key of the object. Must be at least 1 character long.
	// The key of the object
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// [map] A map of metadata to store with the object in IONOS S3 Object Storage. Metadata keys must be lowercase alphanumeric characters.
	// A map of metadata to store with the object in IONOS S3 Object Storage
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// [string]The concatenation of the authentication device's serial number, a space, and the value displayed on your authentication device.
	// The concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device. Required to permanently delete a versioned object if versioning is configured with MFA Delete enabled.
	Mfa *string `json:"mfa,omitempty" tf:"mfa,omitempty"`

	// [string] Indicates whether a legal hold is in effect for the object. Valid values are ON and OFF.
	// Specifies whether a legal hold will be applied to this object.
	ObjectLockLegalHold *string `json:"objectLockLegalHold,omitempty" tf:"object_lock_legal_hold,omitempty"`

	// [string] The object lock mode that you want to apply to the object. Valid values are GOVERNANCE and COMPLIANCE.
	// Confirms that the requester knows that they will be charged for the request. Bucket owners need not specify this parameter in their requests.
	ObjectLockMode *string `json:"objectLockMode,omitempty" tf:"object_lock_mode,omitempty"`

	// [string] The date and time when the object lock retention expires.Must be in RFC3999 format
	// The date and time when you want this object's Object Lock to expire. Must be formatted as a timestamp parameter.
	ObjectLockRetainUntilDate *string `json:"objectLockRetainUntilDate,omitempty" tf:"object_lock_retain_until_date,omitempty"`

	// [string] Confirms that the requester knows that they will be charged for the request.
	// Confirms that the requester knows that they will be charged for the request. Bucket owners need not specify this parameter in their requests.
	RequestPayer *string `json:"requestPayer,omitempty" tf:"request_payer,omitempty"`

	// [string] The server-side encryption algorithm used when storing this object in IONOS S3 Object Storage. Valid value is AES256.
	// The server-side encryption algorithm used when storing this object in IONOS S3 Object Storage (AES256).
	ServerSideEncryption *string `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

	// [string] Specifies the IONOS S3 Object Storage Encryption Context for object encryption.
	// Specifies the IONOS S3 Object Storage Encryption Context to use for object encryption. The value of this header is a base64-encoded UTF-8 string holding JSON with the encryption context key-value pairs.
	ServerSideEncryptionContextSecretRef *v1.SecretKeySelector `json:"serverSideEncryptionContextSecretRef,omitempty" tf:"-"`

	// [string] Specifies the algorithm to use for encrypting the object. Valid value is AES256.
	// Specifies the algorithm to use to when encrypting the object (e.g., AES256).
	ServerSideEncryptionCustomerAlgorithm *string `json:"serverSideEncryptionCustomerAlgorithm,omitempty" tf:"server_side_encryption_customer_algorithm,omitempty"`

	// [string] Specifies the 256-bit, base64-encoded encryption key to use to encrypt and decrypt your data.
	// Specifies the 256-bit, base64-encoded encryption key to use to encrypt and decrypt your data
	ServerSideEncryptionCustomerKey *string `json:"serverSideEncryptionCustomerKey,omitempty" tf:"server_side_encryption_customer_key,omitempty"`

	// [string] Specifies the 128-bit MD5 digest of the encryption key.
	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. IONOS S3 Object Storage uses this header for a message integrity check  to ensure that the encryption key was transmitted without error
	ServerSideEncryptionCustomerKeyMd5 *string `json:"serverSideEncryptionCustomerKeyMd5,omitempty" tf:"server_side_encryption_customer_key_md5,omitempty"`

	// [string] The path to the file to upload.
	// The path to the file to upload
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// [string] The storage class of the object. Valid value is STANDARD. Default is STANDARD.
	// The storage class of the object. Valid value is 'STANDARD'.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// [map] The tag-set for the object.
	// The tag-set for the object
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// [string] Redirects requests for this object to another object in the same bucket or to an external URL.
	// If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. IONOS S3 Object Storage stores the value of this header in the object metadata
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

type ObjectObservation struct {

	// [string] The name of the bucket where the object will be stored. Must be between 3 and 63 characters.
	// The name of the bucket
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// [string] Specifies caching behavior along the request/reply chain.
	// Can be used to specify caching behavior along the request/reply chain
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// [string] Inline content of the object.
	// The utf-8 content of the object
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// [string] Specifies presentational information for the object.
	// Specifies presentational information for the object
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// [string] Specifies what content encodings have been applied to the object.
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// [string] The natural language or languages of the intended audience for the object.
	// The natural language or languages of the intended audience for the object
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// [string] A standard MIME type describing the format of the contents.
	// A standard MIME type describing the format of the contents
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// (Computed)[string] An entity tag (ETag) is an opaque identifier assigned by a web server to a specific version of a resource found at a URL.
	// An entity tag (ETag) is an opaque identifier assigned by a web server to a specific version of a resource found at a URL.
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	// [string] The date and time at which the object is no longer cacheable.
	// The date and time at which the object is no longer cacheable
	Expires *string `json:"expires,omitempty" tf:"expires,omitempty"`

	// [bool] If true, the object will be destroyed if versioning is enabled then all versions will be destroyed. Default is false.
	// Specifies whether to delete the object even if it has a governance-type Object Lock in place. You must explicitly pass a value of true for this parameter to delete the object.
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [string] The key of the object. Must be at least 1 character long.
	// The key of the object
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// [map] A map of metadata to store with the object in IONOS S3 Object Storage. Metadata keys must be lowercase alphanumeric characters.
	// A map of metadata to store with the object in IONOS S3 Object Storage
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// [string]The concatenation of the authentication device's serial number, a space, and the value displayed on your authentication device.
	// The concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device. Required to permanently delete a versioned object if versioning is configured with MFA Delete enabled.
	Mfa *string `json:"mfa,omitempty" tf:"mfa,omitempty"`

	// [string] Indicates whether a legal hold is in effect for the object. Valid values are ON and OFF.
	// Specifies whether a legal hold will be applied to this object.
	ObjectLockLegalHold *string `json:"objectLockLegalHold,omitempty" tf:"object_lock_legal_hold,omitempty"`

	// [string] The object lock mode that you want to apply to the object. Valid values are GOVERNANCE and COMPLIANCE.
	// Confirms that the requester knows that they will be charged for the request. Bucket owners need not specify this parameter in their requests.
	ObjectLockMode *string `json:"objectLockMode,omitempty" tf:"object_lock_mode,omitempty"`

	// [string] The date and time when the object lock retention expires.Must be in RFC3999 format
	// The date and time when you want this object's Object Lock to expire. Must be formatted as a timestamp parameter.
	ObjectLockRetainUntilDate *string `json:"objectLockRetainUntilDate,omitempty" tf:"object_lock_retain_until_date,omitempty"`

	// [string] Confirms that the requester knows that they will be charged for the request.
	// Confirms that the requester knows that they will be charged for the request. Bucket owners need not specify this parameter in their requests.
	RequestPayer *string `json:"requestPayer,omitempty" tf:"request_payer,omitempty"`

	// [string] The server-side encryption algorithm used when storing this object in IONOS S3 Object Storage. Valid value is AES256.
	// The server-side encryption algorithm used when storing this object in IONOS S3 Object Storage (AES256).
	ServerSideEncryption *string `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

	// [string] Specifies the algorithm to use for encrypting the object. Valid value is AES256.
	// Specifies the algorithm to use to when encrypting the object (e.g., AES256).
	ServerSideEncryptionCustomerAlgorithm *string `json:"serverSideEncryptionCustomerAlgorithm,omitempty" tf:"server_side_encryption_customer_algorithm,omitempty"`

	// [string] Specifies the 256-bit, base64-encoded encryption key to use to encrypt and decrypt your data.
	// Specifies the 256-bit, base64-encoded encryption key to use to encrypt and decrypt your data
	ServerSideEncryptionCustomerKey *string `json:"serverSideEncryptionCustomerKey,omitempty" tf:"server_side_encryption_customer_key,omitempty"`

	// [string] Specifies the 128-bit MD5 digest of the encryption key.
	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. IONOS S3 Object Storage uses this header for a message integrity check  to ensure that the encryption key was transmitted without error
	ServerSideEncryptionCustomerKeyMd5 *string `json:"serverSideEncryptionCustomerKeyMd5,omitempty" tf:"server_side_encryption_customer_key_md5,omitempty"`

	// [string] The path to the file to upload.
	// The path to the file to upload
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// [string] The storage class of the object. Valid value is STANDARD. Default is STANDARD.
	// The storage class of the object. Valid value is 'STANDARD'.
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// [map] The tag-set for the object.
	// The tag-set for the object
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// (Computed)[string] The version of the object.
	// The version of the object
	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`

	// [string] Redirects requests for this object to another object in the same bucket or to an external URL.
	// If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. IONOS S3 Object Storage stores the value of this header in the object metadata
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

type ObjectParameters struct {

	// [string] The name of the bucket where the object will be stored. Must be between 3 and 63 characters.
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

	// [string] Specifies caching behavior along the request/reply chain.
	// Can be used to specify caching behavior along the request/reply chain
	// +kubebuilder:validation:Optional
	CacheControl *string `json:"cacheControl,omitempty" tf:"cache_control,omitempty"`

	// [string] Inline content of the object.
	// The utf-8 content of the object
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// [string] Specifies presentational information for the object.
	// Specifies presentational information for the object
	// +kubebuilder:validation:Optional
	ContentDisposition *string `json:"contentDisposition,omitempty" tf:"content_disposition,omitempty"`

	// [string] Specifies what content encodings have been applied to the object.
	// Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field
	// +kubebuilder:validation:Optional
	ContentEncoding *string `json:"contentEncoding,omitempty" tf:"content_encoding,omitempty"`

	// [string] The natural language or languages of the intended audience for the object.
	// The natural language or languages of the intended audience for the object
	// +kubebuilder:validation:Optional
	ContentLanguage *string `json:"contentLanguage,omitempty" tf:"content_language,omitempty"`

	// [string] A standard MIME type describing the format of the contents.
	// A standard MIME type describing the format of the contents
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// [string] The date and time at which the object is no longer cacheable.
	// The date and time at which the object is no longer cacheable
	// +kubebuilder:validation:Optional
	Expires *string `json:"expires,omitempty" tf:"expires,omitempty"`

	// [bool] If true, the object will be destroyed if versioning is enabled then all versions will be destroyed. Default is false.
	// Specifies whether to delete the object even if it has a governance-type Object Lock in place. You must explicitly pass a value of true for this parameter to delete the object.
	// +kubebuilder:validation:Optional
	ForceDestroy *bool `json:"forceDestroy,omitempty" tf:"force_destroy,omitempty"`

	// [string] The key of the object. Must be at least 1 character long.
	// The key of the object
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// [map] A map of metadata to store with the object in IONOS S3 Object Storage. Metadata keys must be lowercase alphanumeric characters.
	// A map of metadata to store with the object in IONOS S3 Object Storage
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// [string]The concatenation of the authentication device's serial number, a space, and the value displayed on your authentication device.
	// The concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device. Required to permanently delete a versioned object if versioning is configured with MFA Delete enabled.
	// +kubebuilder:validation:Optional
	Mfa *string `json:"mfa,omitempty" tf:"mfa,omitempty"`

	// [string] Indicates whether a legal hold is in effect for the object. Valid values are ON and OFF.
	// Specifies whether a legal hold will be applied to this object.
	// +kubebuilder:validation:Optional
	ObjectLockLegalHold *string `json:"objectLockLegalHold,omitempty" tf:"object_lock_legal_hold,omitempty"`

	// [string] The object lock mode that you want to apply to the object. Valid values are GOVERNANCE and COMPLIANCE.
	// Confirms that the requester knows that they will be charged for the request. Bucket owners need not specify this parameter in their requests.
	// +kubebuilder:validation:Optional
	ObjectLockMode *string `json:"objectLockMode,omitempty" tf:"object_lock_mode,omitempty"`

	// [string] The date and time when the object lock retention expires.Must be in RFC3999 format
	// The date and time when you want this object's Object Lock to expire. Must be formatted as a timestamp parameter.
	// +kubebuilder:validation:Optional
	ObjectLockRetainUntilDate *string `json:"objectLockRetainUntilDate,omitempty" tf:"object_lock_retain_until_date,omitempty"`

	// [string] Confirms that the requester knows that they will be charged for the request.
	// Confirms that the requester knows that they will be charged for the request. Bucket owners need not specify this parameter in their requests.
	// +kubebuilder:validation:Optional
	RequestPayer *string `json:"requestPayer,omitempty" tf:"request_payer,omitempty"`

	// [string] The server-side encryption algorithm used when storing this object in IONOS S3 Object Storage. Valid value is AES256.
	// The server-side encryption algorithm used when storing this object in IONOS S3 Object Storage (AES256).
	// +kubebuilder:validation:Optional
	ServerSideEncryption *string `json:"serverSideEncryption,omitempty" tf:"server_side_encryption,omitempty"`

	// [string] Specifies the IONOS S3 Object Storage Encryption Context for object encryption.
	// Specifies the IONOS S3 Object Storage Encryption Context to use for object encryption. The value of this header is a base64-encoded UTF-8 string holding JSON with the encryption context key-value pairs.
	// +kubebuilder:validation:Optional
	ServerSideEncryptionContextSecretRef *v1.SecretKeySelector `json:"serverSideEncryptionContextSecretRef,omitempty" tf:"-"`

	// [string] Specifies the algorithm to use for encrypting the object. Valid value is AES256.
	// Specifies the algorithm to use to when encrypting the object (e.g., AES256).
	// +kubebuilder:validation:Optional
	ServerSideEncryptionCustomerAlgorithm *string `json:"serverSideEncryptionCustomerAlgorithm,omitempty" tf:"server_side_encryption_customer_algorithm,omitempty"`

	// [string] Specifies the 256-bit, base64-encoded encryption key to use to encrypt and decrypt your data.
	// Specifies the 256-bit, base64-encoded encryption key to use to encrypt and decrypt your data
	// +kubebuilder:validation:Optional
	ServerSideEncryptionCustomerKey *string `json:"serverSideEncryptionCustomerKey,omitempty" tf:"server_side_encryption_customer_key,omitempty"`

	// [string] Specifies the 128-bit MD5 digest of the encryption key.
	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. IONOS S3 Object Storage uses this header for a message integrity check  to ensure that the encryption key was transmitted without error
	// +kubebuilder:validation:Optional
	ServerSideEncryptionCustomerKeyMd5 *string `json:"serverSideEncryptionCustomerKeyMd5,omitempty" tf:"server_side_encryption_customer_key_md5,omitempty"`

	// [string] The path to the file to upload.
	// The path to the file to upload
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// [string] The storage class of the object. Valid value is STANDARD. Default is STANDARD.
	// The storage class of the object. Valid value is 'STANDARD'.
	// +kubebuilder:validation:Optional
	StorageClass *string `json:"storageClass,omitempty" tf:"storage_class,omitempty"`

	// [map] The tag-set for the object.
	// The tag-set for the object
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// [string] Redirects requests for this object to another object in the same bucket or to an external URL.
	// If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. IONOS S3 Object Storage stores the value of this header in the object metadata
	// +kubebuilder:validation:Optional
	WebsiteRedirect *string `json:"websiteRedirect,omitempty" tf:"website_redirect,omitempty"`
}

// ObjectSpec defines the desired state of Object
type ObjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObjectParameters `json:"forProvider"`
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
	InitProvider ObjectInitParameters `json:"initProvider,omitempty"`
}

// ObjectStatus defines the observed state of Object.
type ObjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Object is the Schema for the Objects API. Creates and manages IonosCloud S3 Objects.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type Object struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.key) || (has(self.initProvider) && has(self.initProvider.key))",message="spec.forProvider.key is a required parameter"
	Spec   ObjectSpec   `json:"spec"`
	Status ObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectList contains a list of Objects
type ObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Object `json:"items"`
}

// Repository type metadata.
var (
	Object_Kind             = "Object"
	Object_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Object_Kind}.String()
	Object_KindAPIVersion   = Object_Kind + "." + CRDGroupVersion.String()
	Object_GroupVersionKind = CRDGroupVersion.WithKind(Object_Kind)
)

func init() {
	SchemeBuilder.Register(&Object{}, &ObjectList{})
}