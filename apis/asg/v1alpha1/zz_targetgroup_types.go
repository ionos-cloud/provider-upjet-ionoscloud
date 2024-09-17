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

type HTTPHealthCheckInitParameters struct {

	// [string]
	MatchType *string `json:"matchType,omitempty" tf:"match_type,omitempty"`

	// [string] The method for the HTTP health check.
	// The method for the HTTP health check.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// [bool]
	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	// [string] The path (destination URL) for the HTTP health check request; the default is /.
	// The path (destination URL) for the HTTP health check request; the default is /.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// [bool]
	Regex *bool `json:"regex,omitempty" tf:"regex,omitempty"`

	// [string] The response returned by the request, depending on the match type.
	// The response returned by the request, depending on the match type.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type HTTPHealthCheckObservation struct {

	// [string]
	MatchType *string `json:"matchType,omitempty" tf:"match_type,omitempty"`

	// [string] The method for the HTTP health check.
	// The method for the HTTP health check.
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// [bool]
	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	// [string] The path (destination URL) for the HTTP health check request; the default is /.
	// The path (destination URL) for the HTTP health check request; the default is /.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// [bool]
	Regex *bool `json:"regex,omitempty" tf:"regex,omitempty"`

	// [string] The response returned by the request, depending on the match type.
	// The response returned by the request, depending on the match type.
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type HTTPHealthCheckParameters struct {

	// [string]
	// +kubebuilder:validation:Optional
	MatchType *string `json:"matchType" tf:"match_type,omitempty"`

	// [string] The method for the HTTP health check.
	// The method for the HTTP health check.
	// +kubebuilder:validation:Optional
	Method *string `json:"method,omitempty" tf:"method,omitempty"`

	// [bool]
	// +kubebuilder:validation:Optional
	Negate *bool `json:"negate,omitempty" tf:"negate,omitempty"`

	// [string] The path (destination URL) for the HTTP health check request; the default is /.
	// The path (destination URL) for the HTTP health check request; the default is /.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// [bool]
	// +kubebuilder:validation:Optional
	Regex *bool `json:"regex,omitempty" tf:"regex,omitempty"`

	// [string] The response returned by the request, depending on the match type.
	// The response returned by the request, depending on the match type.
	// +kubebuilder:validation:Optional
	Response *string `json:"response" tf:"response,omitempty"`
}

type HealthCheckInitParameters struct {

	// [int] The interval in milliseconds between consecutive health checks; default is 2000.
	// The interval in milliseconds between consecutive health checks; default is 2000.
	CheckInterval *float64 `json:"checkInterval,omitempty" tf:"check_interval,omitempty"`

	// [int] The maximum time in milliseconds to wait for a target to respond to a check. For target VMs with 'Check Interval' set, the lesser of the two  values is used once the TCP connection is established.
	// The maximum time in milliseconds to wait for a target to respond to a check. For target VMs with 'Check Interval' set, the lesser of the two  values is used once the TCP connection is established.
	CheckTimeout *float64 `json:"checkTimeout,omitempty" tf:"check_timeout,omitempty"`

	// [int] The maximum number of attempts to reconnect to a target after a connection failure. Valid range is 0 to 65535, and default is three reconnection.
	// The maximum number of attempts to reconnect to a target after a connection failure. Valid range is 0 to 65535, and default is three reconnection.
	Retries *float64 `json:"retries,omitempty" tf:"retries,omitempty"`
}

type HealthCheckObservation struct {

	// [int] The interval in milliseconds between consecutive health checks; default is 2000.
	// The interval in milliseconds between consecutive health checks; default is 2000.
	CheckInterval *float64 `json:"checkInterval,omitempty" tf:"check_interval,omitempty"`

	// [int] The maximum time in milliseconds to wait for a target to respond to a check. For target VMs with 'Check Interval' set, the lesser of the two  values is used once the TCP connection is established.
	// The maximum time in milliseconds to wait for a target to respond to a check. For target VMs with 'Check Interval' set, the lesser of the two  values is used once the TCP connection is established.
	CheckTimeout *float64 `json:"checkTimeout,omitempty" tf:"check_timeout,omitempty"`

	// [int] The maximum number of attempts to reconnect to a target after a connection failure. Valid range is 0 to 65535, and default is three reconnection.
	// The maximum number of attempts to reconnect to a target after a connection failure. Valid range is 0 to 65535, and default is three reconnection.
	Retries *float64 `json:"retries,omitempty" tf:"retries,omitempty"`
}

type HealthCheckParameters struct {

	// [int] The interval in milliseconds between consecutive health checks; default is 2000.
	// The interval in milliseconds between consecutive health checks; default is 2000.
	// +kubebuilder:validation:Optional
	CheckInterval *float64 `json:"checkInterval,omitempty" tf:"check_interval,omitempty"`

	// [int] The maximum time in milliseconds to wait for a target to respond to a check. For target VMs with 'Check Interval' set, the lesser of the two  values is used once the TCP connection is established.
	// The maximum time in milliseconds to wait for a target to respond to a check. For target VMs with 'Check Interval' set, the lesser of the two  values is used once the TCP connection is established.
	// +kubebuilder:validation:Optional
	CheckTimeout *float64 `json:"checkTimeout,omitempty" tf:"check_timeout,omitempty"`

	// [int] The maximum number of attempts to reconnect to a target after a connection failure. Valid range is 0 to 65535, and default is three reconnection.
	// The maximum number of attempts to reconnect to a target after a connection failure. Valid range is 0 to 65535, and default is three reconnection.
	// +kubebuilder:validation:Optional
	Retries *float64 `json:"retries,omitempty" tf:"retries,omitempty"`
}

type TargetGroupInitParameters_2 struct {

	// [string] Balancing algorithm.
	// Balancing algorithm.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Http health check attributes for Target Group
	// Http health check attributes for Target Group
	HTTPHealthCheck *HTTPHealthCheckInitParameters `json:"httpHealthCheck,omitempty" tf:"http_health_check,omitempty"`

	// Health check attributes for Target Group.
	// Health check attributes for Application Load Balancer forwarding rule
	HealthCheck *HealthCheckInitParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// [string] The name of the target group.
	// The name of the target group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [string] Balancing protocol.
	// Balancing protocol.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// [string] The forwarding protocol version. Value is ignored when protocol is not 'HTTP'.
	// The forwarding protocol version. Value is ignored when protocol is not 'HTTP'.
	ProtocolVersion *string `json:"protocolVersion,omitempty" tf:"protocol_version,omitempty"`

	// [list] Array of items in the collection
	// Array of items in the collection.
	Targets []TargetsInitParameters `json:"targets,omitempty" tf:"targets,omitempty"`
}

type TargetGroupObservation_2 struct {

	// [string] Balancing algorithm.
	// Balancing algorithm.
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Http health check attributes for Target Group
	// Http health check attributes for Target Group
	HTTPHealthCheck *HTTPHealthCheckObservation `json:"httpHealthCheck,omitempty" tf:"http_health_check,omitempty"`

	// Health check attributes for Target Group.
	// Health check attributes for Application Load Balancer forwarding rule
	HealthCheck *HealthCheckObservation `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [string] The name of the target group.
	// The name of the target group.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [string] Balancing protocol.
	// Balancing protocol.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// [string] The forwarding protocol version. Value is ignored when protocol is not 'HTTP'.
	// The forwarding protocol version. Value is ignored when protocol is not 'HTTP'.
	ProtocolVersion *string `json:"protocolVersion,omitempty" tf:"protocol_version,omitempty"`

	// [list] Array of items in the collection
	// Array of items in the collection.
	Targets []TargetsObservation `json:"targets,omitempty" tf:"targets,omitempty"`
}

type TargetGroupParameters_2 struct {

	// [string] Balancing algorithm.
	// Balancing algorithm.
	// +kubebuilder:validation:Optional
	Algorithm *string `json:"algorithm,omitempty" tf:"algorithm,omitempty"`

	// Http health check attributes for Target Group
	// Http health check attributes for Target Group
	// +kubebuilder:validation:Optional
	HTTPHealthCheck *HTTPHealthCheckParameters `json:"httpHealthCheck,omitempty" tf:"http_health_check,omitempty"`

	// Health check attributes for Target Group.
	// Health check attributes for Application Load Balancer forwarding rule
	// +kubebuilder:validation:Optional
	HealthCheck *HealthCheckParameters `json:"healthCheck,omitempty" tf:"health_check,omitempty"`

	// [string] The name of the target group.
	// The name of the target group.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// [string] Balancing protocol.
	// Balancing protocol.
	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// [string] The forwarding protocol version. Value is ignored when protocol is not 'HTTP'.
	// The forwarding protocol version. Value is ignored when protocol is not 'HTTP'.
	// +kubebuilder:validation:Optional
	ProtocolVersion *string `json:"protocolVersion,omitempty" tf:"protocol_version,omitempty"`

	// [list] Array of items in the collection
	// Array of items in the collection.
	// +kubebuilder:validation:Optional
	Targets []TargetsParameters `json:"targets,omitempty" tf:"targets,omitempty"`
}

type TargetsInitParameters struct {

	// [bool] Makes the target available only if it accepts periodic health check TCP connection attempts; when turned off, the target is considered always available. The health check only consists of a connection attempt to the address and port of the target. Default is True.
	// Makes the target available only if it accepts periodic health check TCP connection attempts; when turned off, the target is considered always available. The health check only consists of a connection attempt to the address and port of the target. Default is True.
	HealthCheckEnabled *bool `json:"healthCheckEnabled,omitempty" tf:"health_check_enabled,omitempty"`

	// [string] The IP of the balanced target VM.
	// The IP of the balanced target VM.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// [bool] Maintenance mode prevents the target from receiving balanced traffic.
	// Maintenance mode prevents the target from receiving balanced traffic.
	MaintenanceEnabled *bool `json:"maintenanceEnabled,omitempty" tf:"maintenance_enabled,omitempty"`

	// [int] The port of the balanced target service; valid range is 1 to 65535.
	// The port of the balanced target service; valid range is 1 to 65535.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// [string] The proxy protocol version. Accepted values are none, v1, v2, v2ssl. If unspecified, the default value of none is used.
	// Proxy protocol version
	ProxyProtocol *string `json:"proxyProtocol,omitempty" tf:"proxy_protocol,omitempty"`

	// [int] Traffic is distributed in proportion to target weight, relative to the combined weight of all targets. A target with higher weight receives a greater share of traffic. Valid range is 0 to 256 and default is 1; targets with weight of 0 do not participate in load balancing but still accept persistent connections. It is best use values in the middle of the range to leave room for later adjustments.
	// Traffic is distributed in proportion to target weight, relative to the combined weight of all targets. A target with higher weight receives a greater share of traffic. Valid range is 0 to 256 and default is 1; targets with weight of 0 do not participate in load balancing but still accept persistent connections. It is best use values in the middle of the range to leave room for later adjustments.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TargetsObservation struct {

	// [bool] Makes the target available only if it accepts periodic health check TCP connection attempts; when turned off, the target is considered always available. The health check only consists of a connection attempt to the address and port of the target. Default is True.
	// Makes the target available only if it accepts periodic health check TCP connection attempts; when turned off, the target is considered always available. The health check only consists of a connection attempt to the address and port of the target. Default is True.
	HealthCheckEnabled *bool `json:"healthCheckEnabled,omitempty" tf:"health_check_enabled,omitempty"`

	// [string] The IP of the balanced target VM.
	// The IP of the balanced target VM.
	IP *string `json:"ip,omitempty" tf:"ip,omitempty"`

	// [bool] Maintenance mode prevents the target from receiving balanced traffic.
	// Maintenance mode prevents the target from receiving balanced traffic.
	MaintenanceEnabled *bool `json:"maintenanceEnabled,omitempty" tf:"maintenance_enabled,omitempty"`

	// [int] The port of the balanced target service; valid range is 1 to 65535.
	// The port of the balanced target service; valid range is 1 to 65535.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// [string] The proxy protocol version. Accepted values are none, v1, v2, v2ssl. If unspecified, the default value of none is used.
	// Proxy protocol version
	ProxyProtocol *string `json:"proxyProtocol,omitempty" tf:"proxy_protocol,omitempty"`

	// [int] Traffic is distributed in proportion to target weight, relative to the combined weight of all targets. A target with higher weight receives a greater share of traffic. Valid range is 0 to 256 and default is 1; targets with weight of 0 do not participate in load balancing but still accept persistent connections. It is best use values in the middle of the range to leave room for later adjustments.
	// Traffic is distributed in proportion to target weight, relative to the combined weight of all targets. A target with higher weight receives a greater share of traffic. Valid range is 0 to 256 and default is 1; targets with weight of 0 do not participate in load balancing but still accept persistent connections. It is best use values in the middle of the range to leave room for later adjustments.
	Weight *float64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type TargetsParameters struct {

	// [bool] Makes the target available only if it accepts periodic health check TCP connection attempts; when turned off, the target is considered always available. The health check only consists of a connection attempt to the address and port of the target. Default is True.
	// Makes the target available only if it accepts periodic health check TCP connection attempts; when turned off, the target is considered always available. The health check only consists of a connection attempt to the address and port of the target. Default is True.
	// +kubebuilder:validation:Optional
	HealthCheckEnabled *bool `json:"healthCheckEnabled,omitempty" tf:"health_check_enabled,omitempty"`

	// [string] The IP of the balanced target VM.
	// The IP of the balanced target VM.
	// +kubebuilder:validation:Optional
	IP *string `json:"ip" tf:"ip,omitempty"`

	// [bool] Maintenance mode prevents the target from receiving balanced traffic.
	// Maintenance mode prevents the target from receiving balanced traffic.
	// +kubebuilder:validation:Optional
	MaintenanceEnabled *bool `json:"maintenanceEnabled,omitempty" tf:"maintenance_enabled,omitempty"`

	// [int] The port of the balanced target service; valid range is 1 to 65535.
	// The port of the balanced target service; valid range is 1 to 65535.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port" tf:"port,omitempty"`

	// [string] The proxy protocol version. Accepted values are none, v1, v2, v2ssl. If unspecified, the default value of none is used.
	// Proxy protocol version
	// +kubebuilder:validation:Optional
	ProxyProtocol *string `json:"proxyProtocol,omitempty" tf:"proxy_protocol,omitempty"`

	// [int] Traffic is distributed in proportion to target weight, relative to the combined weight of all targets. A target with higher weight receives a greater share of traffic. Valid range is 0 to 256 and default is 1; targets with weight of 0 do not participate in load balancing but still accept persistent connections. It is best use values in the middle of the range to leave room for later adjustments.
	// Traffic is distributed in proportion to target weight, relative to the combined weight of all targets. A target with higher weight receives a greater share of traffic. Valid range is 0 to 256 and default is 1; targets with weight of 0 do not participate in load balancing but still accept persistent connections. It is best use values in the middle of the range to leave room for later adjustments.
	// +kubebuilder:validation:Optional
	Weight *float64 `json:"weight" tf:"weight,omitempty"`
}

// TargetGroupSpec defines the desired state of TargetGroup
type TargetGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TargetGroupParameters_2 `json:"forProvider"`
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
	InitProvider TargetGroupInitParameters_2 `json:"initProvider,omitempty"`
}

// TargetGroupStatus defines the observed state of TargetGroup.
type TargetGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TargetGroupObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TargetGroup is the Schema for the TargetGroups API. Creates and manages IonosCloud Target Group.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type TargetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.algorithm) || (has(self.initProvider) && has(self.initProvider.algorithm))",message="spec.forProvider.algorithm is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocol) || (has(self.initProvider) && has(self.initProvider.protocol))",message="spec.forProvider.protocol is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.protocolVersion) || (has(self.initProvider) && has(self.initProvider.protocolVersion))",message="spec.forProvider.protocolVersion is a required parameter"
	Spec   TargetGroupSpec   `json:"spec"`
	Status TargetGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetGroupList contains a list of TargetGroups
type TargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TargetGroup `json:"items"`
}

// Repository type metadata.
var (
	TargetGroup_Kind             = "TargetGroup"
	TargetGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TargetGroup_Kind}.String()
	TargetGroup_KindAPIVersion   = TargetGroup_Kind + "." + CRDGroupVersion.String()
	TargetGroup_GroupVersionKind = CRDGroupVersion.WithKind(TargetGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&TargetGroup{}, &TargetGroupList{})
}