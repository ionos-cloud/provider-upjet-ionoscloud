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
func (in *BackupInitParameters) DeepCopyInto(out *BackupInitParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.PointInTimeWindowHours != nil {
		in, out := &in.PointInTimeWindowHours, &out.PointInTimeWindowHours
		*out = new(float64)
		**out = **in
	}
	if in.SnapshotIntervalHours != nil {
		in, out := &in.SnapshotIntervalHours, &out.SnapshotIntervalHours
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupInitParameters.
func (in *BackupInitParameters) DeepCopy() *BackupInitParameters {
	if in == nil {
		return nil
	}
	out := new(BackupInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupObservation) DeepCopyInto(out *BackupObservation) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.PointInTimeWindowHours != nil {
		in, out := &in.PointInTimeWindowHours, &out.PointInTimeWindowHours
		*out = new(float64)
		**out = **in
	}
	if in.SnapshotIntervalHours != nil {
		in, out := &in.SnapshotIntervalHours, &out.SnapshotIntervalHours
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupObservation.
func (in *BackupObservation) DeepCopy() *BackupObservation {
	if in == nil {
		return nil
	}
	out := new(BackupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupParameters) DeepCopyInto(out *BackupParameters) {
	*out = *in
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.PointInTimeWindowHours != nil {
		in, out := &in.PointInTimeWindowHours, &out.PointInTimeWindowHours
		*out = new(float64)
		**out = **in
	}
	if in.SnapshotIntervalHours != nil {
		in, out := &in.SnapshotIntervalHours, &out.SnapshotIntervalHours
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupParameters.
func (in *BackupParameters) DeepCopy() *BackupParameters {
	if in == nil {
		return nil
	}
	out := new(BackupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BiConnectorInitParameters) DeepCopyInto(out *BiConnectorInitParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BiConnectorInitParameters.
func (in *BiConnectorInitParameters) DeepCopy() *BiConnectorInitParameters {
	if in == nil {
		return nil
	}
	out := new(BiConnectorInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BiConnectorObservation) DeepCopyInto(out *BiConnectorObservation) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BiConnectorObservation.
func (in *BiConnectorObservation) DeepCopy() *BiConnectorObservation {
	if in == nil {
		return nil
	}
	out := new(BiConnectorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BiConnectorParameters) DeepCopyInto(out *BiConnectorParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BiConnectorParameters.
func (in *BiConnectorParameters) DeepCopy() *BiConnectorParameters {
	if in == nil {
		return nil
	}
	out := new(BiConnectorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionsInitParameters) DeepCopyInto(out *ConnectionsInitParameters) {
	*out = *in
	if in.CidrList != nil {
		in, out := &in.CidrList, &out.CidrList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
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
	if in.CidrList != nil {
		in, out := &in.CidrList, &out.CidrList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DatacenterID != nil {
		in, out := &in.DatacenterID, &out.DatacenterID
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
	if in.CidrList != nil {
		in, out := &in.CidrList, &out.CidrList
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
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
func (in *MongodbCluster) DeepCopyInto(out *MongodbCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodbCluster.
func (in *MongodbCluster) DeepCopy() *MongodbCluster {
	if in == nil {
		return nil
	}
	out := new(MongodbCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongodbCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodbClusterInitParameters) DeepCopyInto(out *MongodbClusterInitParameters) {
	*out = *in
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = make([]BackupInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BiConnector != nil {
		in, out := &in.BiConnector, &out.BiConnector
		*out = make([]BiConnectorInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Connections != nil {
		in, out := &in.Connections, &out.Connections
		*out = make([]ConnectionsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Cores != nil {
		in, out := &in.Cores, &out.Cores
		*out = new(float64)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Edition != nil {
		in, out := &in.Edition, &out.Edition
		*out = new(string)
		**out = **in
	}
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = new(float64)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.LocationRef != nil {
		in, out := &in.LocationRef, &out.LocationRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.LocationSelector != nil {
		in, out := &in.LocationSelector, &out.LocationSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = make([]MaintenanceWindowInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MongodbVersion != nil {
		in, out := &in.MongodbVersion, &out.MongodbVersion
		*out = new(string)
		**out = **in
	}
	if in.RAM != nil {
		in, out := &in.RAM, &out.RAM
		*out = new(float64)
		**out = **in
	}
	if in.Shards != nil {
		in, out := &in.Shards, &out.Shards
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
	if in.TemplateID != nil {
		in, out := &in.TemplateID, &out.TemplateID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodbClusterInitParameters.
func (in *MongodbClusterInitParameters) DeepCopy() *MongodbClusterInitParameters {
	if in == nil {
		return nil
	}
	out := new(MongodbClusterInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodbClusterList) DeepCopyInto(out *MongodbClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MongodbCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodbClusterList.
func (in *MongodbClusterList) DeepCopy() *MongodbClusterList {
	if in == nil {
		return nil
	}
	out := new(MongodbClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongodbClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodbClusterObservation) DeepCopyInto(out *MongodbClusterObservation) {
	*out = *in
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = make([]BackupObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BiConnector != nil {
		in, out := &in.BiConnector, &out.BiConnector
		*out = make([]BiConnectorObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConnectionString != nil {
		in, out := &in.ConnectionString, &out.ConnectionString
		*out = new(string)
		**out = **in
	}
	if in.Connections != nil {
		in, out := &in.Connections, &out.Connections
		*out = make([]ConnectionsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Cores != nil {
		in, out := &in.Cores, &out.Cores
		*out = new(float64)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Edition != nil {
		in, out := &in.Edition, &out.Edition
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = new(float64)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = make([]MaintenanceWindowObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MongodbVersion != nil {
		in, out := &in.MongodbVersion, &out.MongodbVersion
		*out = new(string)
		**out = **in
	}
	if in.RAM != nil {
		in, out := &in.RAM, &out.RAM
		*out = new(float64)
		**out = **in
	}
	if in.Shards != nil {
		in, out := &in.Shards, &out.Shards
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
	if in.TemplateID != nil {
		in, out := &in.TemplateID, &out.TemplateID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodbClusterObservation.
func (in *MongodbClusterObservation) DeepCopy() *MongodbClusterObservation {
	if in == nil {
		return nil
	}
	out := new(MongodbClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodbClusterParameters) DeepCopyInto(out *MongodbClusterParameters) {
	*out = *in
	if in.Backup != nil {
		in, out := &in.Backup, &out.Backup
		*out = make([]BackupParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BiConnector != nil {
		in, out := &in.BiConnector, &out.BiConnector
		*out = make([]BiConnectorParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Connections != nil {
		in, out := &in.Connections, &out.Connections
		*out = make([]ConnectionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Cores != nil {
		in, out := &in.Cores, &out.Cores
		*out = new(float64)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Edition != nil {
		in, out := &in.Edition, &out.Edition
		*out = new(string)
		**out = **in
	}
	if in.Instances != nil {
		in, out := &in.Instances, &out.Instances
		*out = new(float64)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.LocationRef != nil {
		in, out := &in.LocationRef, &out.LocationRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.LocationSelector != nil {
		in, out := &in.LocationSelector, &out.LocationSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = make([]MaintenanceWindowParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MongodbVersion != nil {
		in, out := &in.MongodbVersion, &out.MongodbVersion
		*out = new(string)
		**out = **in
	}
	if in.RAM != nil {
		in, out := &in.RAM, &out.RAM
		*out = new(float64)
		**out = **in
	}
	if in.Shards != nil {
		in, out := &in.Shards, &out.Shards
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
	if in.TemplateID != nil {
		in, out := &in.TemplateID, &out.TemplateID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodbClusterParameters.
func (in *MongodbClusterParameters) DeepCopy() *MongodbClusterParameters {
	if in == nil {
		return nil
	}
	out := new(MongodbClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodbClusterSpec) DeepCopyInto(out *MongodbClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodbClusterSpec.
func (in *MongodbClusterSpec) DeepCopy() *MongodbClusterSpec {
	if in == nil {
		return nil
	}
	out := new(MongodbClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodbClusterStatus) DeepCopyInto(out *MongodbClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodbClusterStatus.
func (in *MongodbClusterStatus) DeepCopy() *MongodbClusterStatus {
	if in == nil {
		return nil
	}
	out := new(MongodbClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodbUser) DeepCopyInto(out *MongodbUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodbUser.
func (in *MongodbUser) DeepCopy() *MongodbUser {
	if in == nil {
		return nil
	}
	out := new(MongodbUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongodbUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodbUserInitParameters) DeepCopyInto(out *MongodbUserInitParameters) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ClusterIDRef != nil {
		in, out := &in.ClusterIDRef, &out.ClusterIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterIDSelector != nil {
		in, out := &in.ClusterIDSelector, &out.ClusterIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	out.PasswordSecretRef = in.PasswordSecretRef
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]RolesInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodbUserInitParameters.
func (in *MongodbUserInitParameters) DeepCopy() *MongodbUserInitParameters {
	if in == nil {
		return nil
	}
	out := new(MongodbUserInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodbUserList) DeepCopyInto(out *MongodbUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MongodbUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodbUserList.
func (in *MongodbUserList) DeepCopy() *MongodbUserList {
	if in == nil {
		return nil
	}
	out := new(MongodbUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MongodbUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodbUserObservation) DeepCopyInto(out *MongodbUserObservation) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]RolesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodbUserObservation.
func (in *MongodbUserObservation) DeepCopy() *MongodbUserObservation {
	if in == nil {
		return nil
	}
	out := new(MongodbUserObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodbUserParameters) DeepCopyInto(out *MongodbUserParameters) {
	*out = *in
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ClusterIDRef != nil {
		in, out := &in.ClusterIDRef, &out.ClusterIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterIDSelector != nil {
		in, out := &in.ClusterIDSelector, &out.ClusterIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	out.PasswordSecretRef = in.PasswordSecretRef
	if in.Roles != nil {
		in, out := &in.Roles, &out.Roles
		*out = make([]RolesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodbUserParameters.
func (in *MongodbUserParameters) DeepCopy() *MongodbUserParameters {
	if in == nil {
		return nil
	}
	out := new(MongodbUserParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodbUserSpec) DeepCopyInto(out *MongodbUserSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodbUserSpec.
func (in *MongodbUserSpec) DeepCopy() *MongodbUserSpec {
	if in == nil {
		return nil
	}
	out := new(MongodbUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodbUserStatus) DeepCopyInto(out *MongodbUserStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodbUserStatus.
func (in *MongodbUserStatus) DeepCopy() *MongodbUserStatus {
	if in == nil {
		return nil
	}
	out := new(MongodbUserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolesInitParameters) DeepCopyInto(out *RolesInitParameters) {
	*out = *in
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolesInitParameters.
func (in *RolesInitParameters) DeepCopy() *RolesInitParameters {
	if in == nil {
		return nil
	}
	out := new(RolesInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolesObservation) DeepCopyInto(out *RolesObservation) {
	*out = *in
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolesObservation.
func (in *RolesObservation) DeepCopy() *RolesObservation {
	if in == nil {
		return nil
	}
	out := new(RolesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolesParameters) DeepCopyInto(out *RolesParameters) {
	*out = *in
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolesParameters.
func (in *RolesParameters) DeepCopy() *RolesParameters {
	if in == nil {
		return nil
	}
	out := new(RolesParameters)
	in.DeepCopyInto(out)
	return out
}