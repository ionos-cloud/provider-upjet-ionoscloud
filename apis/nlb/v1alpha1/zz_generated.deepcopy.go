//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowlogInitParameters) DeepCopyInto(out *FlowlogInitParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Direction != nil {
		in, out := &in.Direction, &out.Direction
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowlogInitParameters.
func (in *FlowlogInitParameters) DeepCopy() *FlowlogInitParameters {
	if in == nil {
		return nil
	}
	out := new(FlowlogInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowlogObservation) DeepCopyInto(out *FlowlogObservation) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Direction != nil {
		in, out := &in.Direction, &out.Direction
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowlogObservation.
func (in *FlowlogObservation) DeepCopy() *FlowlogObservation {
	if in == nil {
		return nil
	}
	out := new(FlowlogObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowlogParameters) DeepCopyInto(out *FlowlogParameters) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.Direction != nil {
		in, out := &in.Direction, &out.Direction
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowlogParameters.
func (in *FlowlogParameters) DeepCopy() *FlowlogParameters {
	if in == nil {
		return nil
	}
	out := new(FlowlogParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Forwardingrule) DeepCopyInto(out *Forwardingrule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Forwardingrule.
func (in *Forwardingrule) DeepCopy() *Forwardingrule {
	if in == nil {
		return nil
	}
	out := new(Forwardingrule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Forwardingrule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingruleInitParameters) DeepCopyInto(out *ForwardingruleInitParameters) {
	*out = *in
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(string)
		**out = **in
	}
	if in.DatacenterID != nil {
		in, out := &in.DatacenterID, &out.DatacenterID
		*out = new(string)
		**out = **in
	}
	if in.DatacenterIDRef != nil {
		in, out := &in.DatacenterIDRef, &out.DatacenterIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DatacenterIDSelector != nil {
		in, out := &in.DatacenterIDSelector, &out.DatacenterIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(HealthCheckInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.ListenerIP != nil {
		in, out := &in.ListenerIP, &out.ListenerIP
		*out = new(string)
		**out = **in
	}
	if in.ListenerPort != nil {
		in, out := &in.ListenerPort, &out.ListenerPort
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkloadbalancerID != nil {
		in, out := &in.NetworkloadbalancerID, &out.NetworkloadbalancerID
		*out = new(string)
		**out = **in
	}
	if in.NetworkloadbalancerIDRef != nil {
		in, out := &in.NetworkloadbalancerIDRef, &out.NetworkloadbalancerIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkloadbalancerIDSelector != nil {
		in, out := &in.NetworkloadbalancerIDSelector, &out.NetworkloadbalancerIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]TargetsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingruleInitParameters.
func (in *ForwardingruleInitParameters) DeepCopy() *ForwardingruleInitParameters {
	if in == nil {
		return nil
	}
	out := new(ForwardingruleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingruleList) DeepCopyInto(out *ForwardingruleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Forwardingrule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingruleList.
func (in *ForwardingruleList) DeepCopy() *ForwardingruleList {
	if in == nil {
		return nil
	}
	out := new(ForwardingruleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ForwardingruleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingruleObservation) DeepCopyInto(out *ForwardingruleObservation) {
	*out = *in
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(string)
		**out = **in
	}
	if in.DatacenterID != nil {
		in, out := &in.DatacenterID, &out.DatacenterID
		*out = new(string)
		**out = **in
	}
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(HealthCheckObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ListenerIP != nil {
		in, out := &in.ListenerIP, &out.ListenerIP
		*out = new(string)
		**out = **in
	}
	if in.ListenerPort != nil {
		in, out := &in.ListenerPort, &out.ListenerPort
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkloadbalancerID != nil {
		in, out := &in.NetworkloadbalancerID, &out.NetworkloadbalancerID
		*out = new(string)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]TargetsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingruleObservation.
func (in *ForwardingruleObservation) DeepCopy() *ForwardingruleObservation {
	if in == nil {
		return nil
	}
	out := new(ForwardingruleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingruleParameters) DeepCopyInto(out *ForwardingruleParameters) {
	*out = *in
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(string)
		**out = **in
	}
	if in.DatacenterID != nil {
		in, out := &in.DatacenterID, &out.DatacenterID
		*out = new(string)
		**out = **in
	}
	if in.DatacenterIDRef != nil {
		in, out := &in.DatacenterIDRef, &out.DatacenterIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DatacenterIDSelector != nil {
		in, out := &in.DatacenterIDSelector, &out.DatacenterIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(HealthCheckParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.ListenerIP != nil {
		in, out := &in.ListenerIP, &out.ListenerIP
		*out = new(string)
		**out = **in
	}
	if in.ListenerPort != nil {
		in, out := &in.ListenerPort, &out.ListenerPort
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NetworkloadbalancerID != nil {
		in, out := &in.NetworkloadbalancerID, &out.NetworkloadbalancerID
		*out = new(string)
		**out = **in
	}
	if in.NetworkloadbalancerIDRef != nil {
		in, out := &in.NetworkloadbalancerIDRef, &out.NetworkloadbalancerIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NetworkloadbalancerIDSelector != nil {
		in, out := &in.NetworkloadbalancerIDSelector, &out.NetworkloadbalancerIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]TargetsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingruleParameters.
func (in *ForwardingruleParameters) DeepCopy() *ForwardingruleParameters {
	if in == nil {
		return nil
	}
	out := new(ForwardingruleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingruleSpec) DeepCopyInto(out *ForwardingruleSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingruleSpec.
func (in *ForwardingruleSpec) DeepCopy() *ForwardingruleSpec {
	if in == nil {
		return nil
	}
	out := new(ForwardingruleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForwardingruleStatus) DeepCopyInto(out *ForwardingruleStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForwardingruleStatus.
func (in *ForwardingruleStatus) DeepCopy() *ForwardingruleStatus {
	if in == nil {
		return nil
	}
	out := new(ForwardingruleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckInitParameters) DeepCopyInto(out *HealthCheckInitParameters) {
	*out = *in
	if in.ClientTimeout != nil {
		in, out := &in.ClientTimeout, &out.ClientTimeout
		*out = new(float64)
		**out = **in
	}
	if in.ConnectTimeout != nil {
		in, out := &in.ConnectTimeout, &out.ConnectTimeout
		*out = new(float64)
		**out = **in
	}
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		*out = new(float64)
		**out = **in
	}
	if in.TargetTimeout != nil {
		in, out := &in.TargetTimeout, &out.TargetTimeout
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckInitParameters.
func (in *HealthCheckInitParameters) DeepCopy() *HealthCheckInitParameters {
	if in == nil {
		return nil
	}
	out := new(HealthCheckInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckObservation) DeepCopyInto(out *HealthCheckObservation) {
	*out = *in
	if in.ClientTimeout != nil {
		in, out := &in.ClientTimeout, &out.ClientTimeout
		*out = new(float64)
		**out = **in
	}
	if in.ConnectTimeout != nil {
		in, out := &in.ConnectTimeout, &out.ConnectTimeout
		*out = new(float64)
		**out = **in
	}
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		*out = new(float64)
		**out = **in
	}
	if in.TargetTimeout != nil {
		in, out := &in.TargetTimeout, &out.TargetTimeout
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckObservation.
func (in *HealthCheckObservation) DeepCopy() *HealthCheckObservation {
	if in == nil {
		return nil
	}
	out := new(HealthCheckObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheckParameters) DeepCopyInto(out *HealthCheckParameters) {
	*out = *in
	if in.ClientTimeout != nil {
		in, out := &in.ClientTimeout, &out.ClientTimeout
		*out = new(float64)
		**out = **in
	}
	if in.ConnectTimeout != nil {
		in, out := &in.ConnectTimeout, &out.ConnectTimeout
		*out = new(float64)
		**out = **in
	}
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		*out = new(float64)
		**out = **in
	}
	if in.TargetTimeout != nil {
		in, out := &in.TargetTimeout, &out.TargetTimeout
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheckParameters.
func (in *HealthCheckParameters) DeepCopy() *HealthCheckParameters {
	if in == nil {
		return nil
	}
	out := new(HealthCheckParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Networkloadbalancer) DeepCopyInto(out *Networkloadbalancer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Networkloadbalancer.
func (in *Networkloadbalancer) DeepCopy() *Networkloadbalancer {
	if in == nil {
		return nil
	}
	out := new(Networkloadbalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Networkloadbalancer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkloadbalancerInitParameters) DeepCopyInto(out *NetworkloadbalancerInitParameters) {
	*out = *in
	if in.CentralLogging != nil {
		in, out := &in.CentralLogging, &out.CentralLogging
		*out = new(bool)
		**out = **in
	}
	if in.DatacenterID != nil {
		in, out := &in.DatacenterID, &out.DatacenterID
		*out = new(string)
		**out = **in
	}
	if in.DatacenterIDRef != nil {
		in, out := &in.DatacenterIDRef, &out.DatacenterIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DatacenterIDSelector != nil {
		in, out := &in.DatacenterIDSelector, &out.DatacenterIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Flowlog != nil {
		in, out := &in.Flowlog, &out.Flowlog
		*out = new(FlowlogInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Ips != nil {
		in, out := &in.Ips, &out.Ips
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.LBPrivateIps != nil {
		in, out := &in.LBPrivateIps, &out.LBPrivateIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ListenerLan != nil {
		in, out := &in.ListenerLan, &out.ListenerLan
		*out = new(float64)
		**out = **in
	}
	if in.ListenerLanRef != nil {
		in, out := &in.ListenerLanRef, &out.ListenerLanRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ListenerLanSelector != nil {
		in, out := &in.ListenerLanSelector, &out.ListenerLanSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.LoggingFormat != nil {
		in, out := &in.LoggingFormat, &out.LoggingFormat
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TargetLan != nil {
		in, out := &in.TargetLan, &out.TargetLan
		*out = new(float64)
		**out = **in
	}
	if in.TargetLanRef != nil {
		in, out := &in.TargetLanRef, &out.TargetLanRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetLanSelector != nil {
		in, out := &in.TargetLanSelector, &out.TargetLanSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkloadbalancerInitParameters.
func (in *NetworkloadbalancerInitParameters) DeepCopy() *NetworkloadbalancerInitParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkloadbalancerInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkloadbalancerList) DeepCopyInto(out *NetworkloadbalancerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Networkloadbalancer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkloadbalancerList.
func (in *NetworkloadbalancerList) DeepCopy() *NetworkloadbalancerList {
	if in == nil {
		return nil
	}
	out := new(NetworkloadbalancerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkloadbalancerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkloadbalancerObservation) DeepCopyInto(out *NetworkloadbalancerObservation) {
	*out = *in
	if in.CentralLogging != nil {
		in, out := &in.CentralLogging, &out.CentralLogging
		*out = new(bool)
		**out = **in
	}
	if in.DatacenterID != nil {
		in, out := &in.DatacenterID, &out.DatacenterID
		*out = new(string)
		**out = **in
	}
	if in.Flowlog != nil {
		in, out := &in.Flowlog, &out.Flowlog
		*out = new(FlowlogObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Ips != nil {
		in, out := &in.Ips, &out.Ips
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.LBPrivateIps != nil {
		in, out := &in.LBPrivateIps, &out.LBPrivateIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ListenerLan != nil {
		in, out := &in.ListenerLan, &out.ListenerLan
		*out = new(float64)
		**out = **in
	}
	if in.LoggingFormat != nil {
		in, out := &in.LoggingFormat, &out.LoggingFormat
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TargetLan != nil {
		in, out := &in.TargetLan, &out.TargetLan
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkloadbalancerObservation.
func (in *NetworkloadbalancerObservation) DeepCopy() *NetworkloadbalancerObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkloadbalancerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkloadbalancerParameters) DeepCopyInto(out *NetworkloadbalancerParameters) {
	*out = *in
	if in.CentralLogging != nil {
		in, out := &in.CentralLogging, &out.CentralLogging
		*out = new(bool)
		**out = **in
	}
	if in.DatacenterID != nil {
		in, out := &in.DatacenterID, &out.DatacenterID
		*out = new(string)
		**out = **in
	}
	if in.DatacenterIDRef != nil {
		in, out := &in.DatacenterIDRef, &out.DatacenterIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DatacenterIDSelector != nil {
		in, out := &in.DatacenterIDSelector, &out.DatacenterIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Flowlog != nil {
		in, out := &in.Flowlog, &out.Flowlog
		*out = new(FlowlogParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Ips != nil {
		in, out := &in.Ips, &out.Ips
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.LBPrivateIps != nil {
		in, out := &in.LBPrivateIps, &out.LBPrivateIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ListenerLan != nil {
		in, out := &in.ListenerLan, &out.ListenerLan
		*out = new(float64)
		**out = **in
	}
	if in.ListenerLanRef != nil {
		in, out := &in.ListenerLanRef, &out.ListenerLanRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ListenerLanSelector != nil {
		in, out := &in.ListenerLanSelector, &out.ListenerLanSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.LoggingFormat != nil {
		in, out := &in.LoggingFormat, &out.LoggingFormat
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TargetLan != nil {
		in, out := &in.TargetLan, &out.TargetLan
		*out = new(float64)
		**out = **in
	}
	if in.TargetLanRef != nil {
		in, out := &in.TargetLanRef, &out.TargetLanRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetLanSelector != nil {
		in, out := &in.TargetLanSelector, &out.TargetLanSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkloadbalancerParameters.
func (in *NetworkloadbalancerParameters) DeepCopy() *NetworkloadbalancerParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkloadbalancerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkloadbalancerSpec) DeepCopyInto(out *NetworkloadbalancerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkloadbalancerSpec.
func (in *NetworkloadbalancerSpec) DeepCopy() *NetworkloadbalancerSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkloadbalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkloadbalancerStatus) DeepCopyInto(out *NetworkloadbalancerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkloadbalancerStatus.
func (in *NetworkloadbalancerStatus) DeepCopy() *NetworkloadbalancerStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkloadbalancerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetsHealthCheckInitParameters) DeepCopyInto(out *TargetsHealthCheckInitParameters) {
	*out = *in
	if in.Check != nil {
		in, out := &in.Check, &out.Check
		*out = new(bool)
		**out = **in
	}
	if in.CheckInterval != nil {
		in, out := &in.CheckInterval, &out.CheckInterval
		*out = new(float64)
		**out = **in
	}
	if in.Maintenance != nil {
		in, out := &in.Maintenance, &out.Maintenance
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetsHealthCheckInitParameters.
func (in *TargetsHealthCheckInitParameters) DeepCopy() *TargetsHealthCheckInitParameters {
	if in == nil {
		return nil
	}
	out := new(TargetsHealthCheckInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetsHealthCheckObservation) DeepCopyInto(out *TargetsHealthCheckObservation) {
	*out = *in
	if in.Check != nil {
		in, out := &in.Check, &out.Check
		*out = new(bool)
		**out = **in
	}
	if in.CheckInterval != nil {
		in, out := &in.CheckInterval, &out.CheckInterval
		*out = new(float64)
		**out = **in
	}
	if in.Maintenance != nil {
		in, out := &in.Maintenance, &out.Maintenance
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetsHealthCheckObservation.
func (in *TargetsHealthCheckObservation) DeepCopy() *TargetsHealthCheckObservation {
	if in == nil {
		return nil
	}
	out := new(TargetsHealthCheckObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetsHealthCheckParameters) DeepCopyInto(out *TargetsHealthCheckParameters) {
	*out = *in
	if in.Check != nil {
		in, out := &in.Check, &out.Check
		*out = new(bool)
		**out = **in
	}
	if in.CheckInterval != nil {
		in, out := &in.CheckInterval, &out.CheckInterval
		*out = new(float64)
		**out = **in
	}
	if in.Maintenance != nil {
		in, out := &in.Maintenance, &out.Maintenance
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetsHealthCheckParameters.
func (in *TargetsHealthCheckParameters) DeepCopy() *TargetsHealthCheckParameters {
	if in == nil {
		return nil
	}
	out := new(TargetsHealthCheckParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetsInitParameters) DeepCopyInto(out *TargetsInitParameters) {
	*out = *in
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(TargetsHealthCheckInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.ProxyProtocol != nil {
		in, out := &in.ProxyProtocol, &out.ProxyProtocol
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetsInitParameters.
func (in *TargetsInitParameters) DeepCopy() *TargetsInitParameters {
	if in == nil {
		return nil
	}
	out := new(TargetsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetsObservation) DeepCopyInto(out *TargetsObservation) {
	*out = *in
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(TargetsHealthCheckObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.ProxyProtocol != nil {
		in, out := &in.ProxyProtocol, &out.ProxyProtocol
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetsObservation.
func (in *TargetsObservation) DeepCopy() *TargetsObservation {
	if in == nil {
		return nil
	}
	out := new(TargetsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetsParameters) DeepCopyInto(out *TargetsParameters) {
	*out = *in
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(TargetsHealthCheckParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
	if in.ProxyProtocol != nil {
		in, out := &in.ProxyProtocol, &out.ProxyProtocol
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetsParameters.
func (in *TargetsParameters) DeepCopy() *TargetsParameters {
	if in == nil {
		return nil
	}
	out := new(TargetsParameters)
	in.DeepCopyInto(out)
	return out
}
