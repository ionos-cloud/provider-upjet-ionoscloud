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
func (in *CredentialsInitParameters) DeepCopyInto(out *CredentialsInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsInitParameters.
func (in *CredentialsInitParameters) DeepCopy() *CredentialsInitParameters {
	if in == nil {
		return nil
	}
	out := new(CredentialsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsObservation) DeepCopyInto(out *CredentialsObservation) {
	*out = *in
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsObservation.
func (in *CredentialsObservation) DeepCopy() *CredentialsObservation {
	if in == nil {
		return nil
	}
	out := new(CredentialsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsParameters) DeepCopyInto(out *CredentialsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsParameters.
func (in *CredentialsParameters) DeepCopy() *CredentialsParameters {
	if in == nil {
		return nil
	}
	out := new(CredentialsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeaturesInitParameters) DeepCopyInto(out *FeaturesInitParameters) {
	*out = *in
	if in.VulnerabilityScanning != nil {
		in, out := &in.VulnerabilityScanning, &out.VulnerabilityScanning
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeaturesInitParameters.
func (in *FeaturesInitParameters) DeepCopy() *FeaturesInitParameters {
	if in == nil {
		return nil
	}
	out := new(FeaturesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeaturesObservation) DeepCopyInto(out *FeaturesObservation) {
	*out = *in
	if in.VulnerabilityScanning != nil {
		in, out := &in.VulnerabilityScanning, &out.VulnerabilityScanning
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeaturesObservation.
func (in *FeaturesObservation) DeepCopy() *FeaturesObservation {
	if in == nil {
		return nil
	}
	out := new(FeaturesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeaturesParameters) DeepCopyInto(out *FeaturesParameters) {
	*out = *in
	if in.VulnerabilityScanning != nil {
		in, out := &in.VulnerabilityScanning, &out.VulnerabilityScanning
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeaturesParameters.
func (in *FeaturesParameters) DeepCopy() *FeaturesParameters {
	if in == nil {
		return nil
	}
	out := new(FeaturesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GarbageCollectionScheduleInitParameters) DeepCopyInto(out *GarbageCollectionScheduleInitParameters) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GarbageCollectionScheduleInitParameters.
func (in *GarbageCollectionScheduleInitParameters) DeepCopy() *GarbageCollectionScheduleInitParameters {
	if in == nil {
		return nil
	}
	out := new(GarbageCollectionScheduleInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GarbageCollectionScheduleObservation) DeepCopyInto(out *GarbageCollectionScheduleObservation) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GarbageCollectionScheduleObservation.
func (in *GarbageCollectionScheduleObservation) DeepCopy() *GarbageCollectionScheduleObservation {
	if in == nil {
		return nil
	}
	out := new(GarbageCollectionScheduleObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GarbageCollectionScheduleParameters) DeepCopyInto(out *GarbageCollectionScheduleParameters) {
	*out = *in
	if in.Days != nil {
		in, out := &in.Days, &out.Days
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GarbageCollectionScheduleParameters.
func (in *GarbageCollectionScheduleParameters) DeepCopy() *GarbageCollectionScheduleParameters {
	if in == nil {
		return nil
	}
	out := new(GarbageCollectionScheduleParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Registry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryInitParameters) DeepCopyInto(out *RegistryInitParameters) {
	*out = *in
	if in.APISubnetAllowList != nil {
		in, out := &in.APISubnetAllowList, &out.APISubnetAllowList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = new(FeaturesInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.GarbageCollectionSchedule != nil {
		in, out := &in.GarbageCollectionSchedule, &out.GarbageCollectionSchedule
		*out = new(GarbageCollectionScheduleInitParameters)
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryInitParameters.
func (in *RegistryInitParameters) DeepCopy() *RegistryInitParameters {
	if in == nil {
		return nil
	}
	out := new(RegistryInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryList) DeepCopyInto(out *RegistryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Registry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryList.
func (in *RegistryList) DeepCopy() *RegistryList {
	if in == nil {
		return nil
	}
	out := new(RegistryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegistryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryObservation) DeepCopyInto(out *RegistryObservation) {
	*out = *in
	if in.APISubnetAllowList != nil {
		in, out := &in.APISubnetAllowList, &out.APISubnetAllowList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = new(FeaturesObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.GarbageCollectionSchedule != nil {
		in, out := &in.GarbageCollectionSchedule, &out.GarbageCollectionSchedule
		*out = new(GarbageCollectionScheduleObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
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
	if in.StorageUsage != nil {
		in, out := &in.StorageUsage, &out.StorageUsage
		*out = make([]StorageUsageObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryObservation.
func (in *RegistryObservation) DeepCopy() *RegistryObservation {
	if in == nil {
		return nil
	}
	out := new(RegistryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryParameters) DeepCopyInto(out *RegistryParameters) {
	*out = *in
	if in.APISubnetAllowList != nil {
		in, out := &in.APISubnetAllowList, &out.APISubnetAllowList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = new(FeaturesParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.GarbageCollectionSchedule != nil {
		in, out := &in.GarbageCollectionSchedule, &out.GarbageCollectionSchedule
		*out = new(GarbageCollectionScheduleParameters)
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryParameters.
func (in *RegistryParameters) DeepCopy() *RegistryParameters {
	if in == nil {
		return nil
	}
	out := new(RegistryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistrySpec) DeepCopyInto(out *RegistrySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistrySpec.
func (in *RegistrySpec) DeepCopy() *RegistrySpec {
	if in == nil {
		return nil
	}
	out := new(RegistrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryStatus) DeepCopyInto(out *RegistryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryStatus.
func (in *RegistryStatus) DeepCopy() *RegistryStatus {
	if in == nil {
		return nil
	}
	out := new(RegistryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryToken) DeepCopyInto(out *RegistryToken) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryToken.
func (in *RegistryToken) DeepCopy() *RegistryToken {
	if in == nil {
		return nil
	}
	out := new(RegistryToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegistryToken) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryTokenInitParameters) DeepCopyInto(out *RegistryTokenInitParameters) {
	*out = *in
	if in.ExpiryDate != nil {
		in, out := &in.ExpiryDate, &out.ExpiryDate
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(string)
		**out = **in
	}
	if in.RegistryIDRef != nil {
		in, out := &in.RegistryIDRef, &out.RegistryIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RegistryIDSelector != nil {
		in, out := &in.RegistryIDSelector, &out.RegistryIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SavePasswordToFile != nil {
		in, out := &in.SavePasswordToFile, &out.SavePasswordToFile
		*out = new(string)
		**out = **in
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]ScopesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryTokenInitParameters.
func (in *RegistryTokenInitParameters) DeepCopy() *RegistryTokenInitParameters {
	if in == nil {
		return nil
	}
	out := new(RegistryTokenInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryTokenList) DeepCopyInto(out *RegistryTokenList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RegistryToken, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryTokenList.
func (in *RegistryTokenList) DeepCopy() *RegistryTokenList {
	if in == nil {
		return nil
	}
	out := new(RegistryTokenList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegistryTokenList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryTokenObservation) DeepCopyInto(out *RegistryTokenObservation) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]CredentialsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExpiryDate != nil {
		in, out := &in.ExpiryDate, &out.ExpiryDate
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
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(string)
		**out = **in
	}
	if in.SavePasswordToFile != nil {
		in, out := &in.SavePasswordToFile, &out.SavePasswordToFile
		*out = new(string)
		**out = **in
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]ScopesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryTokenObservation.
func (in *RegistryTokenObservation) DeepCopy() *RegistryTokenObservation {
	if in == nil {
		return nil
	}
	out := new(RegistryTokenObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryTokenParameters) DeepCopyInto(out *RegistryTokenParameters) {
	*out = *in
	if in.ExpiryDate != nil {
		in, out := &in.ExpiryDate, &out.ExpiryDate
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(string)
		**out = **in
	}
	if in.RegistryIDRef != nil {
		in, out := &in.RegistryIDRef, &out.RegistryIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.RegistryIDSelector != nil {
		in, out := &in.RegistryIDSelector, &out.RegistryIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.SavePasswordToFile != nil {
		in, out := &in.SavePasswordToFile, &out.SavePasswordToFile
		*out = new(string)
		**out = **in
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]ScopesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryTokenParameters.
func (in *RegistryTokenParameters) DeepCopy() *RegistryTokenParameters {
	if in == nil {
		return nil
	}
	out := new(RegistryTokenParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryTokenSpec) DeepCopyInto(out *RegistryTokenSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryTokenSpec.
func (in *RegistryTokenSpec) DeepCopy() *RegistryTokenSpec {
	if in == nil {
		return nil
	}
	out := new(RegistryTokenSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegistryTokenStatus) DeepCopyInto(out *RegistryTokenStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegistryTokenStatus.
func (in *RegistryTokenStatus) DeepCopy() *RegistryTokenStatus {
	if in == nil {
		return nil
	}
	out := new(RegistryTokenStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScopesInitParameters) DeepCopyInto(out *ScopesInitParameters) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScopesInitParameters.
func (in *ScopesInitParameters) DeepCopy() *ScopesInitParameters {
	if in == nil {
		return nil
	}
	out := new(ScopesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScopesObservation) DeepCopyInto(out *ScopesObservation) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScopesObservation.
func (in *ScopesObservation) DeepCopy() *ScopesObservation {
	if in == nil {
		return nil
	}
	out := new(ScopesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScopesParameters) DeepCopyInto(out *ScopesParameters) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScopesParameters.
func (in *ScopesParameters) DeepCopy() *ScopesParameters {
	if in == nil {
		return nil
	}
	out := new(ScopesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageUsageInitParameters) DeepCopyInto(out *StorageUsageInitParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageUsageInitParameters.
func (in *StorageUsageInitParameters) DeepCopy() *StorageUsageInitParameters {
	if in == nil {
		return nil
	}
	out := new(StorageUsageInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageUsageObservation) DeepCopyInto(out *StorageUsageObservation) {
	*out = *in
	if in.Bytes != nil {
		in, out := &in.Bytes, &out.Bytes
		*out = new(float64)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageUsageObservation.
func (in *StorageUsageObservation) DeepCopy() *StorageUsageObservation {
	if in == nil {
		return nil
	}
	out := new(StorageUsageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageUsageParameters) DeepCopyInto(out *StorageUsageParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageUsageParameters.
func (in *StorageUsageParameters) DeepCopy() *StorageUsageParameters {
	if in == nil {
		return nil
	}
	out := new(StorageUsageParameters)
	in.DeepCopyInto(out)
	return out
}
