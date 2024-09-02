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
func (in *AuthInitParameters) DeepCopyInto(out *AuthInitParameters) {
	*out = *in
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(string)
		**out = **in
	}
	if in.PskKeySecretRef != nil {
		in, out := &in.PskKeySecretRef, &out.PskKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthInitParameters.
func (in *AuthInitParameters) DeepCopy() *AuthInitParameters {
	if in == nil {
		return nil
	}
	out := new(AuthInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthObservation) DeepCopyInto(out *AuthObservation) {
	*out = *in
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthObservation.
func (in *AuthObservation) DeepCopy() *AuthObservation {
	if in == nil {
		return nil
	}
	out := new(AuthObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthParameters) DeepCopyInto(out *AuthParameters) {
	*out = *in
	if in.Method != nil {
		in, out := &in.Method, &out.Method
		*out = new(string)
		**out = **in
	}
	if in.PskKeySecretRef != nil {
		in, out := &in.PskKeySecretRef, &out.PskKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthParameters.
func (in *AuthParameters) DeepCopy() *AuthParameters {
	if in == nil {
		return nil
	}
	out := new(AuthParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionsInitParameters) DeepCopyInto(out *ConnectionsInitParameters) {
	*out = *in
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
	if in.IPv4Cidr != nil {
		in, out := &in.IPv4Cidr, &out.IPv4Cidr
		*out = new(string)
		**out = **in
	}
	if in.IPv6Cidr != nil {
		in, out := &in.IPv6Cidr, &out.IPv6Cidr
		*out = new(string)
		**out = **in
	}
	if in.LanID != nil {
		in, out := &in.LanID, &out.LanID
		*out = new(string)
		**out = **in
	}
	if in.LanIDRef != nil {
		in, out := &in.LanIDRef, &out.LanIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.LanIDSelector != nil {
		in, out := &in.LanIDSelector, &out.LanIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionsInitParameters.
func (in *ConnectionsInitParameters) DeepCopy() *ConnectionsInitParameters {
	if in == nil {
		return nil
	}
	out := new(ConnectionsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionsObservation) DeepCopyInto(out *ConnectionsObservation) {
	*out = *in
	if in.DatacenterID != nil {
		in, out := &in.DatacenterID, &out.DatacenterID
		*out = new(string)
		**out = **in
	}
	if in.IPv4Cidr != nil {
		in, out := &in.IPv4Cidr, &out.IPv4Cidr
		*out = new(string)
		**out = **in
	}
	if in.IPv6Cidr != nil {
		in, out := &in.IPv6Cidr, &out.IPv6Cidr
		*out = new(string)
		**out = **in
	}
	if in.LanID != nil {
		in, out := &in.LanID, &out.LanID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionsObservation.
func (in *ConnectionsObservation) DeepCopy() *ConnectionsObservation {
	if in == nil {
		return nil
	}
	out := new(ConnectionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionsParameters) DeepCopyInto(out *ConnectionsParameters) {
	*out = *in
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
	if in.IPv4Cidr != nil {
		in, out := &in.IPv4Cidr, &out.IPv4Cidr
		*out = new(string)
		**out = **in
	}
	if in.IPv6Cidr != nil {
		in, out := &in.IPv6Cidr, &out.IPv6Cidr
		*out = new(string)
		**out = **in
	}
	if in.LanID != nil {
		in, out := &in.LanID, &out.LanID
		*out = new(string)
		**out = **in
	}
	if in.LanIDRef != nil {
		in, out := &in.LanIDRef, &out.LanIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.LanIDSelector != nil {
		in, out := &in.LanIDSelector, &out.LanIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionsParameters.
func (in *ConnectionsParameters) DeepCopy() *ConnectionsParameters {
	if in == nil {
		return nil
	}
	out := new(ConnectionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EspInitParameters) DeepCopyInto(out *EspInitParameters) {
	*out = *in
	if in.DiffieHellmanGroup != nil {
		in, out := &in.DiffieHellmanGroup, &out.DiffieHellmanGroup
		*out = new(string)
		**out = **in
	}
	if in.EncryptionAlgorithm != nil {
		in, out := &in.EncryptionAlgorithm, &out.EncryptionAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.IntegrityAlgorithm != nil {
		in, out := &in.IntegrityAlgorithm, &out.IntegrityAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Lifetime != nil {
		in, out := &in.Lifetime, &out.Lifetime
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EspInitParameters.
func (in *EspInitParameters) DeepCopy() *EspInitParameters {
	if in == nil {
		return nil
	}
	out := new(EspInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EspObservation) DeepCopyInto(out *EspObservation) {
	*out = *in
	if in.DiffieHellmanGroup != nil {
		in, out := &in.DiffieHellmanGroup, &out.DiffieHellmanGroup
		*out = new(string)
		**out = **in
	}
	if in.EncryptionAlgorithm != nil {
		in, out := &in.EncryptionAlgorithm, &out.EncryptionAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.IntegrityAlgorithm != nil {
		in, out := &in.IntegrityAlgorithm, &out.IntegrityAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Lifetime != nil {
		in, out := &in.Lifetime, &out.Lifetime
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EspObservation.
func (in *EspObservation) DeepCopy() *EspObservation {
	if in == nil {
		return nil
	}
	out := new(EspObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EspParameters) DeepCopyInto(out *EspParameters) {
	*out = *in
	if in.DiffieHellmanGroup != nil {
		in, out := &in.DiffieHellmanGroup, &out.DiffieHellmanGroup
		*out = new(string)
		**out = **in
	}
	if in.EncryptionAlgorithm != nil {
		in, out := &in.EncryptionAlgorithm, &out.EncryptionAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.IntegrityAlgorithm != nil {
		in, out := &in.IntegrityAlgorithm, &out.IntegrityAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Lifetime != nil {
		in, out := &in.Lifetime, &out.Lifetime
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EspParameters.
func (in *EspParameters) DeepCopy() *EspParameters {
	if in == nil {
		return nil
	}
	out := new(EspParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IkeInitParameters) DeepCopyInto(out *IkeInitParameters) {
	*out = *in
	if in.DiffieHellmanGroup != nil {
		in, out := &in.DiffieHellmanGroup, &out.DiffieHellmanGroup
		*out = new(string)
		**out = **in
	}
	if in.EncryptionAlgorithm != nil {
		in, out := &in.EncryptionAlgorithm, &out.EncryptionAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.IntegrityAlgorithm != nil {
		in, out := &in.IntegrityAlgorithm, &out.IntegrityAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Lifetime != nil {
		in, out := &in.Lifetime, &out.Lifetime
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IkeInitParameters.
func (in *IkeInitParameters) DeepCopy() *IkeInitParameters {
	if in == nil {
		return nil
	}
	out := new(IkeInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IkeObservation) DeepCopyInto(out *IkeObservation) {
	*out = *in
	if in.DiffieHellmanGroup != nil {
		in, out := &in.DiffieHellmanGroup, &out.DiffieHellmanGroup
		*out = new(string)
		**out = **in
	}
	if in.EncryptionAlgorithm != nil {
		in, out := &in.EncryptionAlgorithm, &out.EncryptionAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.IntegrityAlgorithm != nil {
		in, out := &in.IntegrityAlgorithm, &out.IntegrityAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Lifetime != nil {
		in, out := &in.Lifetime, &out.Lifetime
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IkeObservation.
func (in *IkeObservation) DeepCopy() *IkeObservation {
	if in == nil {
		return nil
	}
	out := new(IkeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IkeParameters) DeepCopyInto(out *IkeParameters) {
	*out = *in
	if in.DiffieHellmanGroup != nil {
		in, out := &in.DiffieHellmanGroup, &out.DiffieHellmanGroup
		*out = new(string)
		**out = **in
	}
	if in.EncryptionAlgorithm != nil {
		in, out := &in.EncryptionAlgorithm, &out.EncryptionAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.IntegrityAlgorithm != nil {
		in, out := &in.IntegrityAlgorithm, &out.IntegrityAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.Lifetime != nil {
		in, out := &in.Lifetime, &out.Lifetime
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IkeParameters.
func (in *IkeParameters) DeepCopy() *IkeParameters {
	if in == nil {
		return nil
	}
	out := new(IkeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnIpsecGateway) DeepCopyInto(out *VpnIpsecGateway) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnIpsecGateway.
func (in *VpnIpsecGateway) DeepCopy() *VpnIpsecGateway {
	if in == nil {
		return nil
	}
	out := new(VpnIpsecGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnIpsecGateway) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnIpsecGatewayInitParameters) DeepCopyInto(out *VpnIpsecGatewayInitParameters) {
	*out = *in
	if in.Connections != nil {
		in, out := &in.Connections, &out.Connections
		*out = make([]ConnectionsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.GatewayIP != nil {
		in, out := &in.GatewayIP, &out.GatewayIP
		*out = new(string)
		**out = **in
	}
	if in.GatewayIPRef != nil {
		in, out := &in.GatewayIPRef, &out.GatewayIPRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.GatewayIPSelector != nil {
		in, out := &in.GatewayIPSelector, &out.GatewayIPSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnIpsecGatewayInitParameters.
func (in *VpnIpsecGatewayInitParameters) DeepCopy() *VpnIpsecGatewayInitParameters {
	if in == nil {
		return nil
	}
	out := new(VpnIpsecGatewayInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnIpsecGatewayList) DeepCopyInto(out *VpnIpsecGatewayList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpnIpsecGateway, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnIpsecGatewayList.
func (in *VpnIpsecGatewayList) DeepCopy() *VpnIpsecGatewayList {
	if in == nil {
		return nil
	}
	out := new(VpnIpsecGatewayList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnIpsecGatewayList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnIpsecGatewayObservation) DeepCopyInto(out *VpnIpsecGatewayObservation) {
	*out = *in
	if in.Connections != nil {
		in, out := &in.Connections, &out.Connections
		*out = make([]ConnectionsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.GatewayIP != nil {
		in, out := &in.GatewayIP, &out.GatewayIP
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnIpsecGatewayObservation.
func (in *VpnIpsecGatewayObservation) DeepCopy() *VpnIpsecGatewayObservation {
	if in == nil {
		return nil
	}
	out := new(VpnIpsecGatewayObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnIpsecGatewayParameters) DeepCopyInto(out *VpnIpsecGatewayParameters) {
	*out = *in
	if in.Connections != nil {
		in, out := &in.Connections, &out.Connections
		*out = make([]ConnectionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.GatewayIP != nil {
		in, out := &in.GatewayIP, &out.GatewayIP
		*out = new(string)
		**out = **in
	}
	if in.GatewayIPRef != nil {
		in, out := &in.GatewayIPRef, &out.GatewayIPRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.GatewayIPSelector != nil {
		in, out := &in.GatewayIPSelector, &out.GatewayIPSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnIpsecGatewayParameters.
func (in *VpnIpsecGatewayParameters) DeepCopy() *VpnIpsecGatewayParameters {
	if in == nil {
		return nil
	}
	out := new(VpnIpsecGatewayParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnIpsecGatewaySpec) DeepCopyInto(out *VpnIpsecGatewaySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnIpsecGatewaySpec.
func (in *VpnIpsecGatewaySpec) DeepCopy() *VpnIpsecGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(VpnIpsecGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnIpsecGatewayStatus) DeepCopyInto(out *VpnIpsecGatewayStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnIpsecGatewayStatus.
func (in *VpnIpsecGatewayStatus) DeepCopy() *VpnIpsecGatewayStatus {
	if in == nil {
		return nil
	}
	out := new(VpnIpsecGatewayStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnIpsecTunnel) DeepCopyInto(out *VpnIpsecTunnel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnIpsecTunnel.
func (in *VpnIpsecTunnel) DeepCopy() *VpnIpsecTunnel {
	if in == nil {
		return nil
	}
	out := new(VpnIpsecTunnel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnIpsecTunnel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnIpsecTunnelInitParameters) DeepCopyInto(out *VpnIpsecTunnelInitParameters) {
	*out = *in
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = make([]AuthInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CloudNetworkCidrs != nil {
		in, out := &in.CloudNetworkCidrs, &out.CloudNetworkCidrs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Esp != nil {
		in, out := &in.Esp, &out.Esp
		*out = make([]EspInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GatewayID != nil {
		in, out := &in.GatewayID, &out.GatewayID
		*out = new(string)
		**out = **in
	}
	if in.GatewayIDRef != nil {
		in, out := &in.GatewayIDRef, &out.GatewayIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.GatewayIDSelector != nil {
		in, out := &in.GatewayIDSelector, &out.GatewayIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Ike != nil {
		in, out := &in.Ike, &out.Ike
		*out = make([]IkeInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PeerNetworkCidrs != nil {
		in, out := &in.PeerNetworkCidrs, &out.PeerNetworkCidrs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RemoteHost != nil {
		in, out := &in.RemoteHost, &out.RemoteHost
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnIpsecTunnelInitParameters.
func (in *VpnIpsecTunnelInitParameters) DeepCopy() *VpnIpsecTunnelInitParameters {
	if in == nil {
		return nil
	}
	out := new(VpnIpsecTunnelInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnIpsecTunnelList) DeepCopyInto(out *VpnIpsecTunnelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VpnIpsecTunnel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnIpsecTunnelList.
func (in *VpnIpsecTunnelList) DeepCopy() *VpnIpsecTunnelList {
	if in == nil {
		return nil
	}
	out := new(VpnIpsecTunnelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VpnIpsecTunnelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnIpsecTunnelObservation) DeepCopyInto(out *VpnIpsecTunnelObservation) {
	*out = *in
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = make([]AuthObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CloudNetworkCidrs != nil {
		in, out := &in.CloudNetworkCidrs, &out.CloudNetworkCidrs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Esp != nil {
		in, out := &in.Esp, &out.Esp
		*out = make([]EspObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GatewayID != nil {
		in, out := &in.GatewayID, &out.GatewayID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Ike != nil {
		in, out := &in.Ike, &out.Ike
		*out = make([]IkeObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PeerNetworkCidrs != nil {
		in, out := &in.PeerNetworkCidrs, &out.PeerNetworkCidrs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RemoteHost != nil {
		in, out := &in.RemoteHost, &out.RemoteHost
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnIpsecTunnelObservation.
func (in *VpnIpsecTunnelObservation) DeepCopy() *VpnIpsecTunnelObservation {
	if in == nil {
		return nil
	}
	out := new(VpnIpsecTunnelObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnIpsecTunnelParameters) DeepCopyInto(out *VpnIpsecTunnelParameters) {
	*out = *in
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		*out = make([]AuthParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CloudNetworkCidrs != nil {
		in, out := &in.CloudNetworkCidrs, &out.CloudNetworkCidrs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Esp != nil {
		in, out := &in.Esp, &out.Esp
		*out = make([]EspParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GatewayID != nil {
		in, out := &in.GatewayID, &out.GatewayID
		*out = new(string)
		**out = **in
	}
	if in.GatewayIDRef != nil {
		in, out := &in.GatewayIDRef, &out.GatewayIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.GatewayIDSelector != nil {
		in, out := &in.GatewayIDSelector, &out.GatewayIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Ike != nil {
		in, out := &in.Ike, &out.Ike
		*out = make([]IkeParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PeerNetworkCidrs != nil {
		in, out := &in.PeerNetworkCidrs, &out.PeerNetworkCidrs
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RemoteHost != nil {
		in, out := &in.RemoteHost, &out.RemoteHost
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnIpsecTunnelParameters.
func (in *VpnIpsecTunnelParameters) DeepCopy() *VpnIpsecTunnelParameters {
	if in == nil {
		return nil
	}
	out := new(VpnIpsecTunnelParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnIpsecTunnelSpec) DeepCopyInto(out *VpnIpsecTunnelSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnIpsecTunnelSpec.
func (in *VpnIpsecTunnelSpec) DeepCopy() *VpnIpsecTunnelSpec {
	if in == nil {
		return nil
	}
	out := new(VpnIpsecTunnelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VpnIpsecTunnelStatus) DeepCopyInto(out *VpnIpsecTunnelStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VpnIpsecTunnelStatus.
func (in *VpnIpsecTunnelStatus) DeepCopy() *VpnIpsecTunnelStatus {
	if in == nil {
		return nil
	}
	out := new(VpnIpsecTunnelStatus)
	in.DeepCopyInto(out)
	return out
}