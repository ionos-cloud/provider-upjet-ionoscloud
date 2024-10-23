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

type DistributionInitParameters struct {

	// [string] The ID of the certificate to use for the distribution. You can create certificates with the certificate resource.
	// The ID of the certificate to use for the distribution.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/certificatemanager/v1alpha1.Certificate
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// Reference to a Certificate in certificatemanager to populate certificateId.
	// +kubebuilder:validation:Optional
	CertificateIDRef *v1.Reference `json:"certificateIdRef,omitempty" tf:"-"`

	// Selector for a Certificate in certificatemanager to populate certificateId.
	// +kubebuilder:validation:Optional
	CertificateIDSelector *v1.Selector `json:"certificateIdSelector,omitempty" tf:"-"`

	// [string] The domain of the distribution.
	// The domain of the distribution.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// [list] The routing rules for the distribution.
	// The routing rules for the distribution.
	RoutingRules []RoutingRulesInitParameters `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

type DistributionObservation struct {

	// [string] The ID of the certificate to use for the distribution. You can create certificates with the certificate resource.
	// The ID of the certificate to use for the distribution.
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// [string] The domain of the distribution.
	// The domain of the distribution.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// IP of the distribution, it has to be included on the domain DNS Zone as A record.
	// IP of the distribution, it has to be included on the domain DNS Zone as A record.
	PublicEndpointV4 *string `json:"publicEndpointV4,omitempty" tf:"public_endpoint_v4,omitempty"`

	// IP of the distribution, it has to be included on the domain DNS Zone as AAAA record.
	// IP of the distribution, it has to be included on the domain DNS Zone as AAAA record.
	PublicEndpointV6 *string `json:"publicEndpointV6,omitempty" tf:"public_endpoint_v6,omitempty"`

	// Unique resource indentifier.
	// Unique name of the resource.
	ResourceUrn *string `json:"resourceUrn,omitempty" tf:"resource_urn,omitempty"`

	// [list] The routing rules for the distribution.
	// The routing rules for the distribution.
	RoutingRules []RoutingRulesObservation `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

type DistributionParameters struct {

	// [string] The ID of the certificate to use for the distribution. You can create certificates with the certificate resource.
	// The ID of the certificate to use for the distribution.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/certificatemanager/v1alpha1.Certificate
	// +kubebuilder:validation:Optional
	CertificateID *string `json:"certificateId,omitempty" tf:"certificate_id,omitempty"`

	// Reference to a Certificate in certificatemanager to populate certificateId.
	// +kubebuilder:validation:Optional
	CertificateIDRef *v1.Reference `json:"certificateIdRef,omitempty" tf:"-"`

	// Selector for a Certificate in certificatemanager to populate certificateId.
	// +kubebuilder:validation:Optional
	CertificateIDSelector *v1.Selector `json:"certificateIdSelector,omitempty" tf:"-"`

	// [string] The domain of the distribution.
	// The domain of the distribution.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// [list] The routing rules for the distribution.
	// The routing rules for the distribution.
	// +kubebuilder:validation:Optional
	RoutingRules []RoutingRulesParameters `json:"routingRules,omitempty" tf:"routing_rules,omitempty"`
}

type GeoRestrictionsInitParameters struct {

	// [string] List of allowed countries
	AllowList []*string `json:"allowList,omitempty" tf:"allow_list,omitempty"`

	// [string] List of blocked countries
	BlockList []*string `json:"blockList,omitempty" tf:"block_list,omitempty"`
}

type GeoRestrictionsObservation struct {

	// [string] List of allowed countries
	AllowList []*string `json:"allowList,omitempty" tf:"allow_list,omitempty"`

	// [string] List of blocked countries
	BlockList []*string `json:"blockList,omitempty" tf:"block_list,omitempty"`
}

type GeoRestrictionsParameters struct {

	// [string] List of allowed countries
	// +kubebuilder:validation:Optional
	AllowList []*string `json:"allowList,omitempty" tf:"allow_list,omitempty"`

	// [string] List of blocked countries
	// +kubebuilder:validation:Optional
	BlockList []*string `json:"blockList,omitempty" tf:"block_list,omitempty"`
}

type RoutingRulesInitParameters struct {

	// [string] The prefix of the routing rule.
	// The prefix of the routing rule.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// [string] The scheme of the routing rule.
	// The scheme of the routing rule.
	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`

	// [map] - A map of properties for the rule
	Upstream *UpstreamInitParameters `json:"upstream,omitempty" tf:"upstream,omitempty"`
}

type RoutingRulesObservation struct {

	// [string] The prefix of the routing rule.
	// The prefix of the routing rule.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// [string] The scheme of the routing rule.
	// The scheme of the routing rule.
	Scheme *string `json:"scheme,omitempty" tf:"scheme,omitempty"`

	// [map] - A map of properties for the rule
	Upstream *UpstreamObservation `json:"upstream,omitempty" tf:"upstream,omitempty"`
}

type RoutingRulesParameters struct {

	// [string] The prefix of the routing rule.
	// The prefix of the routing rule.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix" tf:"prefix,omitempty"`

	// [string] The scheme of the routing rule.
	// The scheme of the routing rule.
	// +kubebuilder:validation:Optional
	Scheme *string `json:"scheme" tf:"scheme,omitempty"`

	// [map] - A map of properties for the rule
	// +kubebuilder:validation:Optional
	Upstream *UpstreamParameters `json:"upstream" tf:"upstream,omitempty"`
}

type UpstreamInitParameters struct {

	// [bool] Enable or disable caching. If enabled, the CDN will cache the responses from the upstream host. Subsequent requests for the same resource will be served from the cache.
	// Enable or disable caching. If enabled, the CDN will cache the responses from the upstream host. Subsequent requests for the same resource will be served from the cache.
	Caching *bool `json:"caching,omitempty" tf:"caching,omitempty"`

	// [map] - A map of geo_restrictions
	GeoRestrictions *GeoRestrictionsInitParameters `json:"geoRestrictions,omitempty" tf:"geo_restrictions,omitempty"`

	// [string] The upstream host that handles the requests if not already cached. This host will be protected by the WAF if the option is enabled.
	// The upstream host that handles the requests if not already cached. This host will be protected by the WAF if the option is enabled.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// [string] Rate limit class that will be applied to limit the number of incoming requests per IP.
	// Rate limit class that will be applied to limit the number of incoming requests per IP.
	RateLimitClass *string `json:"rateLimitClass,omitempty" tf:"rate_limit_class,omitempty"`

	// [string] The SNI (Server Name Indication) mode of the upstream. It supports two modes: 1) distribution: for outgoing connections to the upstream host, the CDN requires the upstream host to present a valid certificate that matches the configured domain of the CDN distribution; 2) origin: for outgoing connections to the upstream host, the CDN requires the upstream host to present a valid certificate that matches the configured upstream/origin hostname.
	// The SNI (Server Name Indication) mode of the upstream host. It supports two modes: 'distribution' and 'origin', for more information about these modes please check the resource docs.
	SniMode *string `json:"sniMode,omitempty" tf:"sni_mode,omitempty"`

	// [bool] Enable or disable WAF to protect the upstream host.
	// Enable or disable WAF to protect the upstream host.
	Waf *bool `json:"waf,omitempty" tf:"waf,omitempty"`
}

type UpstreamObservation struct {

	// [bool] Enable or disable caching. If enabled, the CDN will cache the responses from the upstream host. Subsequent requests for the same resource will be served from the cache.
	// Enable or disable caching. If enabled, the CDN will cache the responses from the upstream host. Subsequent requests for the same resource will be served from the cache.
	Caching *bool `json:"caching,omitempty" tf:"caching,omitempty"`

	// [map] - A map of geo_restrictions
	GeoRestrictions *GeoRestrictionsObservation `json:"geoRestrictions,omitempty" tf:"geo_restrictions,omitempty"`

	// [string] The upstream host that handles the requests if not already cached. This host will be protected by the WAF if the option is enabled.
	// The upstream host that handles the requests if not already cached. This host will be protected by the WAF if the option is enabled.
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// [string] Rate limit class that will be applied to limit the number of incoming requests per IP.
	// Rate limit class that will be applied to limit the number of incoming requests per IP.
	RateLimitClass *string `json:"rateLimitClass,omitempty" tf:"rate_limit_class,omitempty"`

	// [string] The SNI (Server Name Indication) mode of the upstream. It supports two modes: 1) distribution: for outgoing connections to the upstream host, the CDN requires the upstream host to present a valid certificate that matches the configured domain of the CDN distribution; 2) origin: for outgoing connections to the upstream host, the CDN requires the upstream host to present a valid certificate that matches the configured upstream/origin hostname.
	// The SNI (Server Name Indication) mode of the upstream host. It supports two modes: 'distribution' and 'origin', for more information about these modes please check the resource docs.
	SniMode *string `json:"sniMode,omitempty" tf:"sni_mode,omitempty"`

	// [bool] Enable or disable WAF to protect the upstream host.
	// Enable or disable WAF to protect the upstream host.
	Waf *bool `json:"waf,omitempty" tf:"waf,omitempty"`
}

type UpstreamParameters struct {

	// [bool] Enable or disable caching. If enabled, the CDN will cache the responses from the upstream host. Subsequent requests for the same resource will be served from the cache.
	// Enable or disable caching. If enabled, the CDN will cache the responses from the upstream host. Subsequent requests for the same resource will be served from the cache.
	// +kubebuilder:validation:Optional
	Caching *bool `json:"caching" tf:"caching,omitempty"`

	// [map] - A map of geo_restrictions
	// +kubebuilder:validation:Optional
	GeoRestrictions *GeoRestrictionsParameters `json:"geoRestrictions,omitempty" tf:"geo_restrictions,omitempty"`

	// [string] The upstream host that handles the requests if not already cached. This host will be protected by the WAF if the option is enabled.
	// The upstream host that handles the requests if not already cached. This host will be protected by the WAF if the option is enabled.
	// +kubebuilder:validation:Optional
	Host *string `json:"host" tf:"host,omitempty"`

	// [string] Rate limit class that will be applied to limit the number of incoming requests per IP.
	// Rate limit class that will be applied to limit the number of incoming requests per IP.
	// +kubebuilder:validation:Optional
	RateLimitClass *string `json:"rateLimitClass" tf:"rate_limit_class,omitempty"`

	// [string] The SNI (Server Name Indication) mode of the upstream. It supports two modes: 1) distribution: for outgoing connections to the upstream host, the CDN requires the upstream host to present a valid certificate that matches the configured domain of the CDN distribution; 2) origin: for outgoing connections to the upstream host, the CDN requires the upstream host to present a valid certificate that matches the configured upstream/origin hostname.
	// The SNI (Server Name Indication) mode of the upstream host. It supports two modes: 'distribution' and 'origin', for more information about these modes please check the resource docs.
	// +kubebuilder:validation:Optional
	SniMode *string `json:"sniMode" tf:"sni_mode,omitempty"`

	// [bool] Enable or disable WAF to protect the upstream host.
	// Enable or disable WAF to protect the upstream host.
	// +kubebuilder:validation:Optional
	Waf *bool `json:"waf" tf:"waf,omitempty"`
}

// DistributionSpec defines the desired state of Distribution
type DistributionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DistributionParameters `json:"forProvider"`
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
	InitProvider DistributionInitParameters `json:"initProvider,omitempty"`
}

// DistributionStatus defines the observed state of Distribution.
type DistributionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DistributionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Distribution is the Schema for the Distributions API. Creates and manages IonosCloud CDN Distributions.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type Distribution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.domain) || (has(self.initProvider) && has(self.initProvider.domain))",message="spec.forProvider.domain is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.routingRules) || (has(self.initProvider) && has(self.initProvider.routingRules))",message="spec.forProvider.routingRules is a required parameter"
	Spec   DistributionSpec   `json:"spec"`
	Status DistributionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DistributionList contains a list of Distributions
type DistributionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Distribution `json:"items"`
}

// Repository type metadata.
var (
	Distribution_Kind             = "Distribution"
	Distribution_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Distribution_Kind}.String()
	Distribution_KindAPIVersion   = Distribution_Kind + "." + CRDGroupVersion.String()
	Distribution_GroupVersionKind = CRDGroupVersion.WithKind(Distribution_Kind)
)

func init() {
	SchemeBuilder.Register(&Distribution{}, &DistributionList{})
}
