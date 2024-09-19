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
func (in *AutoScalingInitParameters) DeepCopyInto(out *AutoScalingInitParameters) {
	*out = *in
	if in.MaxNodeCount != nil {
		in, out := &in.MaxNodeCount, &out.MaxNodeCount
		*out = new(float64)
		**out = **in
	}
	if in.MinNodeCount != nil {
		in, out := &in.MinNodeCount, &out.MinNodeCount
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingInitParameters.
func (in *AutoScalingInitParameters) DeepCopy() *AutoScalingInitParameters {
	if in == nil {
		return nil
	}
	out := new(AutoScalingInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingObservation) DeepCopyInto(out *AutoScalingObservation) {
	*out = *in
	if in.MaxNodeCount != nil {
		in, out := &in.MaxNodeCount, &out.MaxNodeCount
		*out = new(float64)
		**out = **in
	}
	if in.MinNodeCount != nil {
		in, out := &in.MinNodeCount, &out.MinNodeCount
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingObservation.
func (in *AutoScalingObservation) DeepCopy() *AutoScalingObservation {
	if in == nil {
		return nil
	}
	out := new(AutoScalingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoScalingParameters) DeepCopyInto(out *AutoScalingParameters) {
	*out = *in
	if in.MaxNodeCount != nil {
		in, out := &in.MaxNodeCount, &out.MaxNodeCount
		*out = new(float64)
		**out = **in
	}
	if in.MinNodeCount != nil {
		in, out := &in.MinNodeCount, &out.MinNodeCount
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoScalingParameters.
func (in *AutoScalingParameters) DeepCopy() *AutoScalingParameters {
	if in == nil {
		return nil
	}
	out := new(AutoScalingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInitParameters) DeepCopyInto(out *ClusterInitParameters) {
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
	if in.AllowReplace != nil {
		in, out := &in.AllowReplace, &out.AllowReplace
		*out = new(bool)
		**out = **in
	}
	if in.K8SVersion != nil {
		in, out := &in.K8SVersion, &out.K8SVersion
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(MaintenanceWindowInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.NATGatewayIP != nil {
		in, out := &in.NATGatewayIP, &out.NATGatewayIP
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NodeSubnet != nil {
		in, out := &in.NodeSubnet, &out.NodeSubnet
		*out = new(string)
		**out = **in
	}
	if in.Public != nil {
		in, out := &in.Public, &out.Public
		*out = new(bool)
		**out = **in
	}
	if in.S3Buckets != nil {
		in, out := &in.S3Buckets, &out.S3Buckets
		*out = make([]S3BucketsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInitParameters.
func (in *ClusterInitParameters) DeepCopy() *ClusterInitParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterObservation) DeepCopyInto(out *ClusterObservation) {
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
	if in.AllowReplace != nil {
		in, out := &in.AllowReplace, &out.AllowReplace
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.K8SVersion != nil {
		in, out := &in.K8SVersion, &out.K8SVersion
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(MaintenanceWindowObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.NATGatewayIP != nil {
		in, out := &in.NATGatewayIP, &out.NATGatewayIP
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NodeSubnet != nil {
		in, out := &in.NodeSubnet, &out.NodeSubnet
		*out = new(string)
		**out = **in
	}
	if in.Public != nil {
		in, out := &in.Public, &out.Public
		*out = new(bool)
		**out = **in
	}
	if in.S3Buckets != nil {
		in, out := &in.S3Buckets, &out.S3Buckets
		*out = make([]S3BucketsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ViableNodePoolVersions != nil {
		in, out := &in.ViableNodePoolVersions, &out.ViableNodePoolVersions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterObservation.
func (in *ClusterObservation) DeepCopy() *ClusterObservation {
	if in == nil {
		return nil
	}
	out := new(ClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterParameters) DeepCopyInto(out *ClusterParameters) {
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
	if in.AllowReplace != nil {
		in, out := &in.AllowReplace, &out.AllowReplace
		*out = new(bool)
		**out = **in
	}
	if in.K8SVersion != nil {
		in, out := &in.K8SVersion, &out.K8SVersion
		*out = new(string)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(MaintenanceWindowParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.NATGatewayIP != nil {
		in, out := &in.NATGatewayIP, &out.NATGatewayIP
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NodeSubnet != nil {
		in, out := &in.NodeSubnet, &out.NodeSubnet
		*out = new(string)
		**out = **in
	}
	if in.Public != nil {
		in, out := &in.Public, &out.Public
		*out = new(bool)
		**out = **in
	}
	if in.S3Buckets != nil {
		in, out := &in.S3Buckets, &out.S3Buckets
		*out = make([]S3BucketsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterParameters.
func (in *ClusterParameters) DeepCopy() *ClusterParameters {
	if in == nil {
		return nil
	}
	out := new(ClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LansInitParameters) DeepCopyInto(out *LansInitParameters) {
	*out = *in
	if in.DHCP != nil {
		in, out := &in.DHCP, &out.DHCP
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(float64)
		**out = **in
	}
	if in.IDRef != nil {
		in, out := &in.IDRef, &out.IDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.IDSelector != nil {
		in, out := &in.IDSelector, &out.IDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]RoutesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LansInitParameters.
func (in *LansInitParameters) DeepCopy() *LansInitParameters {
	if in == nil {
		return nil
	}
	out := new(LansInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LansObservation) DeepCopyInto(out *LansObservation) {
	*out = *in
	if in.DHCP != nil {
		in, out := &in.DHCP, &out.DHCP
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(float64)
		**out = **in
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]RoutesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LansObservation.
func (in *LansObservation) DeepCopy() *LansObservation {
	if in == nil {
		return nil
	}
	out := new(LansObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LansParameters) DeepCopyInto(out *LansParameters) {
	*out = *in
	if in.DHCP != nil {
		in, out := &in.DHCP, &out.DHCP
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(float64)
		**out = **in
	}
	if in.IDRef != nil {
		in, out := &in.IDRef, &out.IDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.IDSelector != nil {
		in, out := &in.IDSelector, &out.IDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Routes != nil {
		in, out := &in.Routes, &out.Routes
		*out = make([]RoutesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LansParameters.
func (in *LansParameters) DeepCopy() *LansParameters {
	if in == nil {
		return nil
	}
	out := new(LansParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceWindowInitParameters) DeepCopyInto(out *MaintenanceWindowInitParameters) {
	*out = *in
	if in.DayOfTheWeek != nil {
		in, out := &in.DayOfTheWeek, &out.DayOfTheWeek
		*out = new(string)
		**out = **in
	}
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceWindowInitParameters.
func (in *MaintenanceWindowInitParameters) DeepCopy() *MaintenanceWindowInitParameters {
	if in == nil {
		return nil
	}
	out := new(MaintenanceWindowInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceWindowObservation) DeepCopyInto(out *MaintenanceWindowObservation) {
	*out = *in
	if in.DayOfTheWeek != nil {
		in, out := &in.DayOfTheWeek, &out.DayOfTheWeek
		*out = new(string)
		**out = **in
	}
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceWindowObservation.
func (in *MaintenanceWindowObservation) DeepCopy() *MaintenanceWindowObservation {
	if in == nil {
		return nil
	}
	out := new(MaintenanceWindowObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaintenanceWindowParameters) DeepCopyInto(out *MaintenanceWindowParameters) {
	*out = *in
	if in.DayOfTheWeek != nil {
		in, out := &in.DayOfTheWeek, &out.DayOfTheWeek
		*out = new(string)
		**out = **in
	}
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaintenanceWindowParameters.
func (in *MaintenanceWindowParameters) DeepCopy() *MaintenanceWindowParameters {
	if in == nil {
		return nil
	}
	out := new(MaintenanceWindowParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePool) DeepCopyInto(out *NodePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePool.
func (in *NodePool) DeepCopy() *NodePool {
	if in == nil {
		return nil
	}
	out := new(NodePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolInitParameters) DeepCopyInto(out *NodePoolInitParameters) {
	*out = *in
	if in.AllowReplace != nil {
		in, out := &in.AllowReplace, &out.AllowReplace
		*out = new(bool)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.AutoScaling != nil {
		in, out := &in.AutoScaling, &out.AutoScaling
		*out = new(AutoScalingInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.CPUFamily != nil {
		in, out := &in.CPUFamily, &out.CPUFamily
		*out = new(string)
		**out = **in
	}
	if in.CoresCount != nil {
		in, out := &in.CoresCount, &out.CoresCount
		*out = new(float64)
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
	if in.K8SClusterID != nil {
		in, out := &in.K8SClusterID, &out.K8SClusterID
		*out = new(string)
		**out = **in
	}
	if in.K8SClusterIDRef != nil {
		in, out := &in.K8SClusterIDRef, &out.K8SClusterIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.K8SClusterIDSelector != nil {
		in, out := &in.K8SClusterIDSelector, &out.K8SClusterIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.K8SVersion != nil {
		in, out := &in.K8SVersion, &out.K8SVersion
		*out = new(string)
		**out = **in
	}
	if in.K8SVersionRef != nil {
		in, out := &in.K8SVersionRef, &out.K8SVersionRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.K8SVersionSelector != nil {
		in, out := &in.K8SVersionSelector, &out.K8SVersionSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Lans != nil {
		in, out := &in.Lans, &out.Lans
		*out = make([]LansInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(NodePoolMaintenanceWindowInitParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(float64)
		**out = **in
	}
	if in.PublicIps != nil {
		in, out := &in.PublicIps, &out.PublicIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RAMSize != nil {
		in, out := &in.RAMSize, &out.RAMSize
		*out = new(float64)
		**out = **in
	}
	if in.StorageSize != nil {
		in, out := &in.StorageSize, &out.StorageSize
		*out = new(float64)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolInitParameters.
func (in *NodePoolInitParameters) DeepCopy() *NodePoolInitParameters {
	if in == nil {
		return nil
	}
	out := new(NodePoolInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolList) DeepCopyInto(out *NodePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolList.
func (in *NodePoolList) DeepCopy() *NodePoolList {
	if in == nil {
		return nil
	}
	out := new(NodePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolMaintenanceWindowInitParameters) DeepCopyInto(out *NodePoolMaintenanceWindowInitParameters) {
	*out = *in
	if in.DayOfTheWeek != nil {
		in, out := &in.DayOfTheWeek, &out.DayOfTheWeek
		*out = new(string)
		**out = **in
	}
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolMaintenanceWindowInitParameters.
func (in *NodePoolMaintenanceWindowInitParameters) DeepCopy() *NodePoolMaintenanceWindowInitParameters {
	if in == nil {
		return nil
	}
	out := new(NodePoolMaintenanceWindowInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolMaintenanceWindowObservation) DeepCopyInto(out *NodePoolMaintenanceWindowObservation) {
	*out = *in
	if in.DayOfTheWeek != nil {
		in, out := &in.DayOfTheWeek, &out.DayOfTheWeek
		*out = new(string)
		**out = **in
	}
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolMaintenanceWindowObservation.
func (in *NodePoolMaintenanceWindowObservation) DeepCopy() *NodePoolMaintenanceWindowObservation {
	if in == nil {
		return nil
	}
	out := new(NodePoolMaintenanceWindowObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolMaintenanceWindowParameters) DeepCopyInto(out *NodePoolMaintenanceWindowParameters) {
	*out = *in
	if in.DayOfTheWeek != nil {
		in, out := &in.DayOfTheWeek, &out.DayOfTheWeek
		*out = new(string)
		**out = **in
	}
	if in.Time != nil {
		in, out := &in.Time, &out.Time
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolMaintenanceWindowParameters.
func (in *NodePoolMaintenanceWindowParameters) DeepCopy() *NodePoolMaintenanceWindowParameters {
	if in == nil {
		return nil
	}
	out := new(NodePoolMaintenanceWindowParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolObservation) DeepCopyInto(out *NodePoolObservation) {
	*out = *in
	if in.AllowReplace != nil {
		in, out := &in.AllowReplace, &out.AllowReplace
		*out = new(bool)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.AutoScaling != nil {
		in, out := &in.AutoScaling, &out.AutoScaling
		*out = new(AutoScalingObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.CPUFamily != nil {
		in, out := &in.CPUFamily, &out.CPUFamily
		*out = new(string)
		**out = **in
	}
	if in.CoresCount != nil {
		in, out := &in.CoresCount, &out.CoresCount
		*out = new(float64)
		**out = **in
	}
	if in.DatacenterID != nil {
		in, out := &in.DatacenterID, &out.DatacenterID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.K8SClusterID != nil {
		in, out := &in.K8SClusterID, &out.K8SClusterID
		*out = new(string)
		**out = **in
	}
	if in.K8SVersion != nil {
		in, out := &in.K8SVersion, &out.K8SVersion
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Lans != nil {
		in, out := &in.Lans, &out.Lans
		*out = make([]LansObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(NodePoolMaintenanceWindowObservation)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(float64)
		**out = **in
	}
	if in.PublicIps != nil {
		in, out := &in.PublicIps, &out.PublicIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RAMSize != nil {
		in, out := &in.RAMSize, &out.RAMSize
		*out = new(float64)
		**out = **in
	}
	if in.StorageSize != nil {
		in, out := &in.StorageSize, &out.StorageSize
		*out = new(float64)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolObservation.
func (in *NodePoolObservation) DeepCopy() *NodePoolObservation {
	if in == nil {
		return nil
	}
	out := new(NodePoolObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolParameters) DeepCopyInto(out *NodePoolParameters) {
	*out = *in
	if in.AllowReplace != nil {
		in, out := &in.AllowReplace, &out.AllowReplace
		*out = new(bool)
		**out = **in
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.AutoScaling != nil {
		in, out := &in.AutoScaling, &out.AutoScaling
		*out = new(AutoScalingParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.CPUFamily != nil {
		in, out := &in.CPUFamily, &out.CPUFamily
		*out = new(string)
		**out = **in
	}
	if in.CoresCount != nil {
		in, out := &in.CoresCount, &out.CoresCount
		*out = new(float64)
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
	if in.K8SClusterID != nil {
		in, out := &in.K8SClusterID, &out.K8SClusterID
		*out = new(string)
		**out = **in
	}
	if in.K8SClusterIDRef != nil {
		in, out := &in.K8SClusterIDRef, &out.K8SClusterIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.K8SClusterIDSelector != nil {
		in, out := &in.K8SClusterIDSelector, &out.K8SClusterIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.K8SVersion != nil {
		in, out := &in.K8SVersion, &out.K8SVersion
		*out = new(string)
		**out = **in
	}
	if in.K8SVersionRef != nil {
		in, out := &in.K8SVersionRef, &out.K8SVersionRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.K8SVersionSelector != nil {
		in, out := &in.K8SVersionSelector, &out.K8SVersionSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Lans != nil {
		in, out := &in.Lans, &out.Lans
		*out = make([]LansParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(NodePoolMaintenanceWindowParameters)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(float64)
		**out = **in
	}
	if in.PublicIps != nil {
		in, out := &in.PublicIps, &out.PublicIps
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.RAMSize != nil {
		in, out := &in.RAMSize, &out.RAMSize
		*out = new(float64)
		**out = **in
	}
	if in.StorageSize != nil {
		in, out := &in.StorageSize, &out.StorageSize
		*out = new(float64)
		**out = **in
	}
	if in.StorageType != nil {
		in, out := &in.StorageType, &out.StorageType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolParameters.
func (in *NodePoolParameters) DeepCopy() *NodePoolParameters {
	if in == nil {
		return nil
	}
	out := new(NodePoolParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolSpec) DeepCopyInto(out *NodePoolSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolSpec.
func (in *NodePoolSpec) DeepCopy() *NodePoolSpec {
	if in == nil {
		return nil
	}
	out := new(NodePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodePoolStatus) DeepCopyInto(out *NodePoolStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodePoolStatus.
func (in *NodePoolStatus) DeepCopy() *NodePoolStatus {
	if in == nil {
		return nil
	}
	out := new(NodePoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutesInitParameters) DeepCopyInto(out *RoutesInitParameters) {
	*out = *in
	if in.GatewayIP != nil {
		in, out := &in.GatewayIP, &out.GatewayIP
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutesInitParameters.
func (in *RoutesInitParameters) DeepCopy() *RoutesInitParameters {
	if in == nil {
		return nil
	}
	out := new(RoutesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutesObservation) DeepCopyInto(out *RoutesObservation) {
	*out = *in
	if in.GatewayIP != nil {
		in, out := &in.GatewayIP, &out.GatewayIP
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutesObservation.
func (in *RoutesObservation) DeepCopy() *RoutesObservation {
	if in == nil {
		return nil
	}
	out := new(RoutesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoutesParameters) DeepCopyInto(out *RoutesParameters) {
	*out = *in
	if in.GatewayIP != nil {
		in, out := &in.GatewayIP, &out.GatewayIP
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoutesParameters.
func (in *RoutesParameters) DeepCopy() *RoutesParameters {
	if in == nil {
		return nil
	}
	out := new(RoutesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketsInitParameters) DeepCopyInto(out *S3BucketsInitParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NameRef != nil {
		in, out := &in.NameRef, &out.NameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NameSelector != nil {
		in, out := &in.NameSelector, &out.NameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketsInitParameters.
func (in *S3BucketsInitParameters) DeepCopy() *S3BucketsInitParameters {
	if in == nil {
		return nil
	}
	out := new(S3BucketsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketsObservation) DeepCopyInto(out *S3BucketsObservation) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketsObservation.
func (in *S3BucketsObservation) DeepCopy() *S3BucketsObservation {
	if in == nil {
		return nil
	}
	out := new(S3BucketsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketsParameters) DeepCopyInto(out *S3BucketsParameters) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NameRef != nil {
		in, out := &in.NameRef, &out.NameRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.NameSelector != nil {
		in, out := &in.NameSelector, &out.NameSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketsParameters.
func (in *S3BucketsParameters) DeepCopy() *S3BucketsParameters {
	if in == nil {
		return nil
	}
	out := new(S3BucketsParameters)
	in.DeepCopyInto(out)
	return out
}
