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

type BucketCorsConfigurationInitParameters struct {

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

	// [block] A block of cors_rule as defined below.
	// A configuration for Cross-Origin Resource Sharing (CORS).
	CorsRule []CorsRuleInitParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`
}

type BucketCorsConfigurationObservation struct {

	// [string] The name of the bucket where the object will be stored.
	// The name of the bucket
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// [block] A block of cors_rule as defined below.
	// A configuration for Cross-Origin Resource Sharing (CORS).
	CorsRule []CorsRuleObservation `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`

	// [int] Container for the Contract Number of the owner
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BucketCorsConfigurationParameters struct {

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

	// [block] A block of cors_rule as defined below.
	// A configuration for Cross-Origin Resource Sharing (CORS).
	// +kubebuilder:validation:Optional
	CorsRule []CorsRuleParameters `json:"corsRule,omitempty" tf:"cors_rule,omitempty"`
}

type CorsRuleInitParameters struct {

	// [list] Specifies which headers are allowed in a preflight OPTIONS request through the Access-Control-Request-Headers header
	// Specifies which headers are allowed in a preflight OPTIONS request through the Access-Control-Request-Headers header.
	// +listType=set
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// [list] An HTTP method that you allow the origin to execute. Valid values are GET, PUT, HEAD, POST, DELETE.
	// An HTTP method that you allow the origin to execute. Valid values are GET, PUT, HEAD, POST, DELETE.
	// +listType=set
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// [list] Specifies which origins are allowed to make requests to the resource.
	// One or more origins you want customers to be able to access the bucket from.
	// +listType=set
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// [list] Specifies which headers are exposed to the browser.
	// One or more headers in the response that you want customers to be able to access from their applications.
	// +listType=set
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// [int] Container for the Contract Number of the owner
	// Container for the Contract Number of the owner.
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// [int] Specifies how long the results of a pre-flight request can be cached in seconds.
	// The time in seconds that your browser is to cache the preflight response for the specified resource.
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type CorsRuleObservation struct {

	// [list] Specifies which headers are allowed in a preflight OPTIONS request through the Access-Control-Request-Headers header
	// Specifies which headers are allowed in a preflight OPTIONS request through the Access-Control-Request-Headers header.
	// +listType=set
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// [list] An HTTP method that you allow the origin to execute. Valid values are GET, PUT, HEAD, POST, DELETE.
	// An HTTP method that you allow the origin to execute. Valid values are GET, PUT, HEAD, POST, DELETE.
	// +listType=set
	AllowedMethods []*string `json:"allowedMethods,omitempty" tf:"allowed_methods,omitempty"`

	// [list] Specifies which origins are allowed to make requests to the resource.
	// One or more origins you want customers to be able to access the bucket from.
	// +listType=set
	AllowedOrigins []*string `json:"allowedOrigins,omitempty" tf:"allowed_origins,omitempty"`

	// [list] Specifies which headers are exposed to the browser.
	// One or more headers in the response that you want customers to be able to access from their applications.
	// +listType=set
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// [int] Container for the Contract Number of the owner
	// Container for the Contract Number of the owner.
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// [int] Specifies how long the results of a pre-flight request can be cached in seconds.
	// The time in seconds that your browser is to cache the preflight response for the specified resource.
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

type CorsRuleParameters struct {

	// [list] Specifies which headers are allowed in a preflight OPTIONS request through the Access-Control-Request-Headers header
	// Specifies which headers are allowed in a preflight OPTIONS request through the Access-Control-Request-Headers header.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedHeaders []*string `json:"allowedHeaders,omitempty" tf:"allowed_headers,omitempty"`

	// [list] An HTTP method that you allow the origin to execute. Valid values are GET, PUT, HEAD, POST, DELETE.
	// An HTTP method that you allow the origin to execute. Valid values are GET, PUT, HEAD, POST, DELETE.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedMethods []*string `json:"allowedMethods" tf:"allowed_methods,omitempty"`

	// [list] Specifies which origins are allowed to make requests to the resource.
	// One or more origins you want customers to be able to access the bucket from.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// [list] Specifies which headers are exposed to the browser.
	// One or more headers in the response that you want customers to be able to access from their applications.
	// +kubebuilder:validation:Optional
	// +listType=set
	ExposeHeaders []*string `json:"exposeHeaders,omitempty" tf:"expose_headers,omitempty"`

	// [int] Container for the Contract Number of the owner
	// Container for the Contract Number of the owner.
	// +kubebuilder:validation:Optional
	ID *float64 `json:"id,omitempty" tf:"id,omitempty"`

	// [int] Specifies how long the results of a pre-flight request can be cached in seconds.
	// The time in seconds that your browser is to cache the preflight response for the specified resource.
	// +kubebuilder:validation:Optional
	MaxAgeSeconds *float64 `json:"maxAgeSeconds,omitempty" tf:"max_age_seconds,omitempty"`
}

// BucketCorsConfigurationSpec defines the desired state of BucketCorsConfiguration
type BucketCorsConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BucketCorsConfigurationParameters `json:"forProvider"`
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
	InitProvider BucketCorsConfigurationInitParameters `json:"initProvider,omitempty"`
}

// BucketCorsConfigurationStatus defines the observed state of BucketCorsConfiguration.
type BucketCorsConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BucketCorsConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// BucketCorsConfiguration is the Schema for the BucketCorsConfigurations API. Manages Buckets cors_configuration on IonosCloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type BucketCorsConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketCorsConfigurationSpec   `json:"spec"`
	Status            BucketCorsConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BucketCorsConfigurationList contains a list of BucketCorsConfigurations
type BucketCorsConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BucketCorsConfiguration `json:"items"`
}

// Repository type metadata.
var (
	BucketCorsConfiguration_Kind             = "BucketCorsConfiguration"
	BucketCorsConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BucketCorsConfiguration_Kind}.String()
	BucketCorsConfiguration_KindAPIVersion   = BucketCorsConfiguration_Kind + "." + CRDGroupVersion.String()
	BucketCorsConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(BucketCorsConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&BucketCorsConfiguration{}, &BucketCorsConfigurationList{})
}