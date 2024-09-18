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

type ConditionsInitParameters struct {

	// Matching rule for the HTTP rule condition attribute; mandatory for HEADER, PATH, QUERY, METHOD, HOST, and COOKIE types; must be null when type is SOURCE_IP.
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// Must be null when type is PATH, METHOD, HOST, or SOURCE_IP. Key can only be set when type is COOKIES, HEADER, or QUERY.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Specifies whether the condition is negated or not; the default is False.
	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	// Type of the HTTP rule condition.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Mandatory for conditions CONTAINS, EQUALS, MATCHES, STARTS_WITH, ENDS_WITH; must be null when condition is EXISTS; should be a valid CIDR if provided and if type is SOURCE_IP.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConditionsObservation struct {

	// Matching rule for the HTTP rule condition attribute; mandatory for HEADER, PATH, QUERY, METHOD, HOST, and COOKIE types; must be null when type is SOURCE_IP.
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// Must be null when type is PATH, METHOD, HOST, or SOURCE_IP. Key can only be set when type is COOKIES, HEADER, or QUERY.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Specifies whether the condition is negated or not; the default is False.
	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	// Type of the HTTP rule condition.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Mandatory for conditions CONTAINS, EQUALS, MATCHES, STARTS_WITH, ENDS_WITH; must be null when condition is EXISTS; should be a valid CIDR if provided and if type is SOURCE_IP.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConditionsParameters struct {

	// Matching rule for the HTTP rule condition attribute; mandatory for HEADER, PATH, QUERY, METHOD, HOST, and COOKIE types; must be null when type is SOURCE_IP.
	// +kubebuilder:validation:Optional
	Condition *string `json:"condition,omitempty" tf:"condition,omitempty"`

	// Must be null when type is PATH, METHOD, HOST, or SOURCE_IP. Key can only be set when type is COOKIES, HEADER, or QUERY.
	// +kubebuilder:validation:Optional
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	// Specifies whether the condition is negated or not; the default is False.
	// +kubebuilder:validation:Optional
	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	// Type of the HTTP rule condition.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`

	// Mandatory for conditions CONTAINS, EQUALS, MATCHES, STARTS_WITH, ENDS_WITH; must be null when condition is EXISTS; should be a valid CIDR if provided and if type is SOURCE_IP.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type HTTPRulesInitParameters struct {

	// An array of items in the collection.The action is only performed if each and every condition is met; if no conditions are set, the rule will always be performed.
	Conditions []ConditionsInitParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// Valid only for STATIC actions.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Default is false; valid only for REDIRECT actions.
	DropQuery *bool `json:"dropQuery,omitempty" tf:"drop_query,omitempty"`

	// The location for redirecting; mandatory and valid only for REDIRECT actions.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The unique name of the Application Load Balancer HTTP rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The response message of the request; mandatory for STATIC actions.
	ResponseMessage *string `json:"responseMessage,omitempty" tf:"response_message,omitempty"`

	// Valid only for REDIRECT and STATIC actions. For REDIRECT actions, default is 301 and possible values are 301, 302, 303, 307, and 308. For STATIC actions, default is 503 and valid range is 200 to 599.
	StatusCode *float64 `json:"statusCode,omitempty" tf:"status_code,omitempty"`

	// The ID of the target group; mandatory and only valid for FORWARD actions.
	TargetGroup *string `json:"targetGroup,omitempty" tf:"target_group,omitempty"`

	// Type of the HTTP rule.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HTTPRulesObservation struct {

	// An array of items in the collection.The action is only performed if each and every condition is met; if no conditions are set, the rule will always be performed.
	Conditions []ConditionsObservation `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// Valid only for STATIC actions.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Default is false; valid only for REDIRECT actions.
	DropQuery *bool `json:"dropQuery,omitempty" tf:"drop_query,omitempty"`

	// The location for redirecting; mandatory and valid only for REDIRECT actions.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The unique name of the Application Load Balancer HTTP rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The response message of the request; mandatory for STATIC actions.
	ResponseMessage *string `json:"responseMessage,omitempty" tf:"response_message,omitempty"`

	// Valid only for REDIRECT and STATIC actions. For REDIRECT actions, default is 301 and possible values are 301, 302, 303, 307, and 308. For STATIC actions, default is 503 and valid range is 200 to 599.
	StatusCode *float64 `json:"statusCode,omitempty" tf:"status_code,omitempty"`

	// The ID of the target group; mandatory and only valid for FORWARD actions.
	TargetGroup *string `json:"targetGroup,omitempty" tf:"target_group,omitempty"`

	// Type of the HTTP rule.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type HTTPRulesParameters struct {

	// An array of items in the collection.The action is only performed if each and every condition is met; if no conditions are set, the rule will always be performed.
	// +kubebuilder:validation:Optional
	Conditions []ConditionsParameters `json:"conditions,omitempty" tf:"conditions,omitempty"`

	// Valid only for STATIC actions.
	// +kubebuilder:validation:Optional
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`

	// Default is false; valid only for REDIRECT actions.
	// +kubebuilder:validation:Optional
	DropQuery *bool `json:"dropQuery,omitempty" tf:"drop_query,omitempty"`

	// The location for redirecting; mandatory and valid only for REDIRECT actions.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The unique name of the Application Load Balancer HTTP rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The response message of the request; mandatory for STATIC actions.
	// +kubebuilder:validation:Optional
	ResponseMessage *string `json:"responseMessage,omitempty" tf:"response_message,omitempty"`

	// Valid only for REDIRECT and STATIC actions. For REDIRECT actions, default is 301 and possible values are 301, 302, 303, 307, and 308. For STATIC actions, default is 503 and valid range is 200 to 599.
	// +kubebuilder:validation:Optional
	StatusCode *float64 `json:"statusCode,omitempty" tf:"status_code,omitempty"`

	// The ID of the target group; mandatory and only valid for FORWARD actions.
	// +kubebuilder:validation:Optional
	TargetGroup *string `json:"targetGroup,omitempty" tf:"target_group,omitempty"`

	// Type of the HTTP rule.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type LoadbalancerForwardingruleInitParameters struct {

	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/alb/v1alpha1.Loadbalancer
	ApplicationLoadbalancerID *string `json:"applicationLoadbalancerId,omitempty" tf:"application_loadbalancer_id,omitempty"`

	// Reference to a Loadbalancer in alb to populate applicationLoadbalancerId.
	// +kubebuilder:validation:Optional
	ApplicationLoadbalancerIDRef *v1.Reference `json:"applicationLoadbalancerIdRef,omitempty" tf:"-"`

	// Selector for a Loadbalancer in alb to populate applicationLoadbalancerId.
	// +kubebuilder:validation:Optional
	ApplicationLoadbalancerIDSelector *v1.Selector `json:"applicationLoadbalancerIdSelector,omitempty" tf:"-"`

	// The maximum time in milliseconds to wait for the client to acknowledge or send data; default is 50,000 (50 seconds).
	ClientTimeout *float64 `json:"clientTimeout,omitempty" tf:"client_timeout,omitempty"`

	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// Array of items in that collection
	HTTPRules []HTTPRulesInitParameters `json:"httpRules,omitempty" tf:"http_rules,omitempty"`

	// Listening (inbound) IP.
	ListenerIP *string `json:"listenerIp,omitempty" tf:"listener_ip,omitempty"`

	// Listening (inbound) port number; valid range is 1 to 65535.
	ListenerPort *float64 `json:"listenerPort,omitempty" tf:"listener_port,omitempty"`

	// The name of the Application Load Balancer forwarding rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Balancing protocol.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Array of items in the collection.
	// +listType=set
	ServerCertificates []*string `json:"serverCertificates,omitempty" tf:"server_certificates,omitempty"`
}

type LoadbalancerForwardingruleObservation struct {
	ApplicationLoadbalancerID *string `json:"applicationLoadbalancerId,omitempty" tf:"application_loadbalancer_id,omitempty"`

	// The maximum time in milliseconds to wait for the client to acknowledge or send data; default is 50,000 (50 seconds).
	ClientTimeout *float64 `json:"clientTimeout,omitempty" tf:"client_timeout,omitempty"`

	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Array of items in that collection
	HTTPRules []HTTPRulesObservation `json:"httpRules,omitempty" tf:"http_rules,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Listening (inbound) IP.
	ListenerIP *string `json:"listenerIp,omitempty" tf:"listener_ip,omitempty"`

	// Listening (inbound) port number; valid range is 1 to 65535.
	ListenerPort *float64 `json:"listenerPort,omitempty" tf:"listener_port,omitempty"`

	// The name of the Application Load Balancer forwarding rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Balancing protocol.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Array of items in the collection.
	// +listType=set
	ServerCertificates []*string `json:"serverCertificates,omitempty" tf:"server_certificates,omitempty"`
}

type LoadbalancerForwardingruleParameters struct {

	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/alb/v1alpha1.Loadbalancer
	// +kubebuilder:validation:Optional
	ApplicationLoadbalancerID *string `json:"applicationLoadbalancerId,omitempty" tf:"application_loadbalancer_id,omitempty"`

	// Reference to a Loadbalancer in alb to populate applicationLoadbalancerId.
	// +kubebuilder:validation:Optional
	ApplicationLoadbalancerIDRef *v1.Reference `json:"applicationLoadbalancerIdRef,omitempty" tf:"-"`

	// Selector for a Loadbalancer in alb to populate applicationLoadbalancerId.
	// +kubebuilder:validation:Optional
	ApplicationLoadbalancerIDSelector *v1.Selector `json:"applicationLoadbalancerIdSelector,omitempty" tf:"-"`

	// The maximum time in milliseconds to wait for the client to acknowledge or send data; default is 50,000 (50 seconds).
	// +kubebuilder:validation:Optional
	ClientTimeout *float64 `json:"clientTimeout,omitempty" tf:"client_timeout,omitempty"`

	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	// +kubebuilder:validation:Optional
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// Array of items in that collection
	// +kubebuilder:validation:Optional
	HTTPRules []HTTPRulesParameters `json:"httpRules,omitempty" tf:"http_rules,omitempty"`

	// Listening (inbound) IP.
	// +kubebuilder:validation:Optional
	ListenerIP *string `json:"listenerIp,omitempty" tf:"listener_ip,omitempty"`

	// Listening (inbound) port number; valid range is 1 to 65535.
	// +kubebuilder:validation:Optional
	ListenerPort *float64 `json:"listenerPort,omitempty" tf:"listener_port,omitempty"`

	// The name of the Application Load Balancer forwarding rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Balancing protocol.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// Array of items in the collection.
	// +kubebuilder:validation:Optional
	// +listType=set
	ServerCertificates []*string `json:"serverCertificates,omitempty" tf:"server_certificates,omitempty"`
}

// LoadbalancerForwardingruleSpec defines the desired state of LoadbalancerForwardingrule
type LoadbalancerForwardingruleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LoadbalancerForwardingruleParameters `json:"forProvider"`
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
	InitProvider LoadbalancerForwardingruleInitParameters `json:"initProvider,omitempty"`
}

// LoadbalancerForwardingruleStatus defines the observed state of LoadbalancerForwardingrule.
type LoadbalancerForwardingruleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LoadbalancerForwardingruleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LoadbalancerForwardingrule is the Schema for the LoadbalancerForwardingrules API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type LoadbalancerForwardingrule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.listenerIp) || (has(self.initProvider) && has(self.initProvider.listenerIp))",message="spec.forProvider.listenerIp is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.listenerPort) || (has(self.initProvider) && has(self.initProvider.listenerPort))",message="spec.forProvider.listenerPort is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	Spec   LoadbalancerForwardingruleSpec   `json:"spec"`
	Status LoadbalancerForwardingruleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadbalancerForwardingruleList contains a list of LoadbalancerForwardingrules
type LoadbalancerForwardingruleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadbalancerForwardingrule `json:"items"`
}

// Repository type metadata.
var (
	LoadbalancerForwardingrule_Kind             = "LoadbalancerForwardingrule"
	LoadbalancerForwardingrule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LoadbalancerForwardingrule_Kind}.String()
	LoadbalancerForwardingrule_KindAPIVersion   = LoadbalancerForwardingrule_Kind + "." + CRDGroupVersion.String()
	LoadbalancerForwardingrule_GroupVersionKind = CRDGroupVersion.WithKind(LoadbalancerForwardingrule_Kind)
)

func init() {
	SchemeBuilder.Register(&LoadbalancerForwardingrule{}, &LoadbalancerForwardingruleList{})
}