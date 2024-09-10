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
func (in *ConnectionsInitParameters) DeepCopyInto(out *ConnectionsInitParameters) {
	*out = *in
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
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
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
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
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
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
func (in *CredentialsInitParameters) DeepCopyInto(out *CredentialsInitParameters) {
	*out = *in
	out.PasswordSecretRef = in.PasswordSecretRef
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
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
	out.PasswordSecretRef = in.PasswordSecretRef
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
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
func (in *FromBackupInitParameters) DeepCopyInto(out *FromBackupInitParameters) {
	*out = *in
	if in.BackupID != nil {
		in, out := &in.BackupID, &out.BackupID
		*out = new(string)
		**out = **in
	}
	if in.RecoveryTargetTime != nil {
		in, out := &in.RecoveryTargetTime, &out.RecoveryTargetTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FromBackupInitParameters.
func (in *FromBackupInitParameters) DeepCopy() *FromBackupInitParameters {
	if in == nil {
		return nil
	}
	out := new(FromBackupInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FromBackupObservation) DeepCopyInto(out *FromBackupObservation) {
	*out = *in
	if in.BackupID != nil {
		in, out := &in.BackupID, &out.BackupID
		*out = new(string)
		**out = **in
	}
	if in.RecoveryTargetTime != nil {
		in, out := &in.RecoveryTargetTime, &out.RecoveryTargetTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FromBackupObservation.
func (in *FromBackupObservation) DeepCopy() *FromBackupObservation {
	if in == nil {
		return nil
	}
	out := new(FromBackupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FromBackupParameters) DeepCopyInto(out *FromBackupParameters) {
	*out = *in
	if in.BackupID != nil {
		in, out := &in.BackupID, &out.BackupID
		*out = new(string)
		**out = **in
	}
	if in.RecoveryTargetTime != nil {
		in, out := &in.RecoveryTargetTime, &out.RecoveryTargetTime
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FromBackupParameters.
func (in *FromBackupParameters) DeepCopy() *FromBackupParameters {
	if in == nil {
		return nil
	}
	out := new(FromBackupParameters)
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
func (in *PostgresqlCluster) DeepCopyInto(out *PostgresqlCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlCluster.
func (in *PostgresqlCluster) DeepCopy() *PostgresqlCluster {
	if in == nil {
		return nil
	}
	out := new(PostgresqlCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlClusterInitParameters) DeepCopyInto(out *PostgresqlClusterInitParameters) {
	*out = *in
	if in.BackupLocation != nil {
		in, out := &in.BackupLocation, &out.BackupLocation
		*out = new(string)
		**out = **in
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
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]CredentialsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.FromBackup != nil {
		in, out := &in.FromBackup, &out.FromBackup
		*out = make([]FromBackupInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.PostgresVersion != nil {
		in, out := &in.PostgresVersion, &out.PostgresVersion
		*out = new(string)
		**out = **in
	}
	if in.RAM != nil {
		in, out := &in.RAM, &out.RAM
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
	if in.SynchronizationMode != nil {
		in, out := &in.SynchronizationMode, &out.SynchronizationMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlClusterInitParameters.
func (in *PostgresqlClusterInitParameters) DeepCopy() *PostgresqlClusterInitParameters {
	if in == nil {
		return nil
	}
	out := new(PostgresqlClusterInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlClusterList) DeepCopyInto(out *PostgresqlClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgresqlCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlClusterList.
func (in *PostgresqlClusterList) DeepCopy() *PostgresqlClusterList {
	if in == nil {
		return nil
	}
	out := new(PostgresqlClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlClusterObservation) DeepCopyInto(out *PostgresqlClusterObservation) {
	*out = *in
	if in.BackupLocation != nil {
		in, out := &in.BackupLocation, &out.BackupLocation
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
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]CredentialsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DNSName != nil {
		in, out := &in.DNSName, &out.DNSName
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.FromBackup != nil {
		in, out := &in.FromBackup, &out.FromBackup
		*out = make([]FromBackupObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.PostgresVersion != nil {
		in, out := &in.PostgresVersion, &out.PostgresVersion
		*out = new(string)
		**out = **in
	}
	if in.RAM != nil {
		in, out := &in.RAM, &out.RAM
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
	if in.SynchronizationMode != nil {
		in, out := &in.SynchronizationMode, &out.SynchronizationMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlClusterObservation.
func (in *PostgresqlClusterObservation) DeepCopy() *PostgresqlClusterObservation {
	if in == nil {
		return nil
	}
	out := new(PostgresqlClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlClusterParameters) DeepCopyInto(out *PostgresqlClusterParameters) {
	*out = *in
	if in.BackupLocation != nil {
		in, out := &in.BackupLocation, &out.BackupLocation
		*out = new(string)
		**out = **in
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
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = make([]CredentialsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.FromBackup != nil {
		in, out := &in.FromBackup, &out.FromBackup
		*out = make([]FromBackupParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
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
	if in.PostgresVersion != nil {
		in, out := &in.PostgresVersion, &out.PostgresVersion
		*out = new(string)
		**out = **in
	}
	if in.RAM != nil {
		in, out := &in.RAM, &out.RAM
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
	if in.SynchronizationMode != nil {
		in, out := &in.SynchronizationMode, &out.SynchronizationMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlClusterParameters.
func (in *PostgresqlClusterParameters) DeepCopy() *PostgresqlClusterParameters {
	if in == nil {
		return nil
	}
	out := new(PostgresqlClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlClusterSpec) DeepCopyInto(out *PostgresqlClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlClusterSpec.
func (in *PostgresqlClusterSpec) DeepCopy() *PostgresqlClusterSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresqlClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlClusterStatus) DeepCopyInto(out *PostgresqlClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlClusterStatus.
func (in *PostgresqlClusterStatus) DeepCopy() *PostgresqlClusterStatus {
	if in == nil {
		return nil
	}
	out := new(PostgresqlClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlDatabase) DeepCopyInto(out *PostgresqlDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlDatabase.
func (in *PostgresqlDatabase) DeepCopy() *PostgresqlDatabase {
	if in == nil {
		return nil
	}
	out := new(PostgresqlDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlDatabaseInitParameters) DeepCopyInto(out *PostgresqlDatabaseInitParameters) {
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlDatabaseInitParameters.
func (in *PostgresqlDatabaseInitParameters) DeepCopy() *PostgresqlDatabaseInitParameters {
	if in == nil {
		return nil
	}
	out := new(PostgresqlDatabaseInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlDatabaseList) DeepCopyInto(out *PostgresqlDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgresqlDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlDatabaseList.
func (in *PostgresqlDatabaseList) DeepCopy() *PostgresqlDatabaseList {
	if in == nil {
		return nil
	}
	out := new(PostgresqlDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlDatabaseObservation) DeepCopyInto(out *PostgresqlDatabaseObservation) {
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlDatabaseObservation.
func (in *PostgresqlDatabaseObservation) DeepCopy() *PostgresqlDatabaseObservation {
	if in == nil {
		return nil
	}
	out := new(PostgresqlDatabaseObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlDatabaseParameters) DeepCopyInto(out *PostgresqlDatabaseParameters) {
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlDatabaseParameters.
func (in *PostgresqlDatabaseParameters) DeepCopy() *PostgresqlDatabaseParameters {
	if in == nil {
		return nil
	}
	out := new(PostgresqlDatabaseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlDatabaseSpec) DeepCopyInto(out *PostgresqlDatabaseSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlDatabaseSpec.
func (in *PostgresqlDatabaseSpec) DeepCopy() *PostgresqlDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresqlDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlDatabaseStatus) DeepCopyInto(out *PostgresqlDatabaseStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlDatabaseStatus.
func (in *PostgresqlDatabaseStatus) DeepCopy() *PostgresqlDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(PostgresqlDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUser) DeepCopyInto(out *PostgresqlUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUser.
func (in *PostgresqlUser) DeepCopy() *PostgresqlUser {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUserInitParameters) DeepCopyInto(out *PostgresqlUserInitParameters) {
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
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUserInitParameters.
func (in *PostgresqlUserInitParameters) DeepCopy() *PostgresqlUserInitParameters {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUserInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUserList) DeepCopyInto(out *PostgresqlUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PostgresqlUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUserList.
func (in *PostgresqlUserList) DeepCopy() *PostgresqlUserList {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PostgresqlUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUserObservation) DeepCopyInto(out *PostgresqlUserObservation) {
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
	if in.IsSystemUser != nil {
		in, out := &in.IsSystemUser, &out.IsSystemUser
		*out = new(bool)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUserObservation.
func (in *PostgresqlUserObservation) DeepCopy() *PostgresqlUserObservation {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUserObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUserParameters) DeepCopyInto(out *PostgresqlUserParameters) {
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
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUserParameters.
func (in *PostgresqlUserParameters) DeepCopy() *PostgresqlUserParameters {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUserParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUserSpec) DeepCopyInto(out *PostgresqlUserSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUserSpec.
func (in *PostgresqlUserSpec) DeepCopy() *PostgresqlUserSpec {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresqlUserStatus) DeepCopyInto(out *PostgresqlUserStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresqlUserStatus.
func (in *PostgresqlUserStatus) DeepCopy() *PostgresqlUserStatus {
	if in == nil {
		return nil
	}
	out := new(PostgresqlUserStatus)
	in.DeepCopyInto(out)
	return out
}
