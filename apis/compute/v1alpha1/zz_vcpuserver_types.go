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

type VCPUServerInitParameters struct {

	// [string] The availability zone in which the server should exist. E.g: AUTO, ZONE_1, ZONE_2. This property is immutable.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// DEPRECATED Please refer to ionoscloud_server_boot_device_selection [string] The associated boot drive, if any. Must be the UUID of a bootable CDROM image that can be retrieved using the ionoscloud_image data source.
	// The associated boot drive, if any. Must be the UUID of a bootable CDROM image that can be retrieved using the ionoscloud_image data source.
	BootCdrom *string `json:"bootCdrom,omitempty" tf:"boot_cdrom,omitempty"`

	// [string] The image or snapshot UUID / name. May also be an image alias. It is required if licence_type is not provided.
	BootImage *string `json:"bootImage,omitempty" tf:"boot_image,omitempty"`

	// [integer] Number of server CPU cores.
	Cores *float64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// [string] The ID of a Virtual Data Center.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// (Computed) The associated firewall rules.
	FirewallruleIds []*string `json:"firewallruleIds,omitempty" tf:"firewallrule_ids,omitempty"`

	// [string] The name, ID or alias of the image. May also be a snapshot ID. It is required if licence_type is not provided. Attribute is immutable.
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// [string] The password for the image.
	ImagePasswordSecretRef *v1.SecretKeySelector `json:"imagePasswordSecretRef,omitempty" tf:"-"`

	// A label can be seen as an object with only two required fields: key and value, both of the string type. Please check the example presented above to see how a label can be used in the plan. A server can have multiple labels.
	Label []VCPUServerLabelInitParameters `json:"label,omitempty" tf:"label,omitempty"`

	// [string] The name of the server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// See the Nic section.
	Nic *VCPUServerNicInitParameters `json:"nic,omitempty" tf:"nic,omitempty"`

	// [integer] The amount of memory for the server in MB.
	RAM *float64 `json:"ram,omitempty" tf:"ram,omitempty"`

	// [list] Immutable List of absolute or relative paths to files containing public SSH key that will be injected into IonosCloud provided Linux images. Also accepts ssh keys directly. Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key. This field may only be set in creation requests. When reading, it always returns null. SSH keys are only supported if a public Linux image is used for the volume creation. Does not support ~ expansion to homedir in the given path.
	// Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key. This field may only be set in creation requests. When reading, it always returns null. SSH keys are only supported if a public Linux image is used for the volume creation.
	SSHKeys []*string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`

	// Sets the power state of the vcpu server. Possible values: `RUNNING` or `SHUTOFF`.
	VMState *string `json:"vmState,omitempty" tf:"vm_state,omitempty"`

	// See the Volume section.
	Volume *VCPUServerVolumeInitParameters `json:"volume,omitempty" tf:"volume,omitempty"`
}

type VCPUServerLabelInitParameters struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type VCPUServerLabelObservation struct {
	Key *string `json:"key,omitempty" tf:"key,omitempty"`

	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type VCPUServerLabelParameters struct {

	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type VCPUServerNicFirewallInitParameters struct {
	IcmpCode *string `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	IcmpType *string `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// [string] The name of the server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	PortRangeEnd *float64 `json:"portRangeEnd,omitempty" tf:"port_range_end,omitempty"`

	PortRangeStart *float64 `json:"portRangeStart,omitempty" tf:"port_range_start,omitempty"`

	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	SourceIP *string `json:"sourceIp,omitempty" tf:"source_ip,omitempty"`

	SourceMac *string `json:"sourceMac,omitempty" tf:"source_mac,omitempty"`

	TargetIP *string `json:"targetIp,omitempty" tf:"target_ip,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type VCPUServerNicFirewallObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IcmpCode *string `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	IcmpType *string `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// [string] The name of the server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	PortRangeEnd *float64 `json:"portRangeEnd,omitempty" tf:"port_range_end,omitempty"`

	PortRangeStart *float64 `json:"portRangeStart,omitempty" tf:"port_range_start,omitempty"`

	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	SourceIP *string `json:"sourceIp,omitempty" tf:"source_ip,omitempty"`

	SourceMac *string `json:"sourceMac,omitempty" tf:"source_mac,omitempty"`

	TargetIP *string `json:"targetIp,omitempty" tf:"target_ip,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type VCPUServerNicFirewallParameters struct {

	// +kubebuilder:validation:Optional
	IcmpCode *string `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`

	// +kubebuilder:validation:Optional
	IcmpType *string `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`

	// [string] The name of the server.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PortRangeEnd *float64 `json:"portRangeEnd,omitempty" tf:"port_range_end,omitempty"`

	// +kubebuilder:validation:Optional
	PortRangeStart *float64 `json:"portRangeStart,omitempty" tf:"port_range_start,omitempty"`

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	SourceIP *string `json:"sourceIp,omitempty" tf:"source_ip,omitempty"`

	// +kubebuilder:validation:Optional
	SourceMac *string `json:"sourceMac,omitempty" tf:"source_mac,omitempty"`

	// +kubebuilder:validation:Optional
	TargetIP *string `json:"targetIp,omitempty" tf:"target_ip,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type VCPUServerNicInitParameters struct {
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	Dhcpv6 *bool `json:"dhcpv6,omitempty" tf:"dhcpv6,omitempty"`

	// Allows to define firewall rules inline in the server. See the Firewall section.
	// Firewall rules created in the server resource. The rules can also be created as separate resources outside the server resource
	Firewall []VCPUServerNicFirewallInitParameters `json:"firewall,omitempty" tf:"firewall,omitempty"`

	FirewallActive *bool `json:"firewallActive,omitempty" tf:"firewall_active,omitempty"`

	FirewallType *string `json:"firewallType,omitempty" tf:"firewall_type,omitempty"`

	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	IPv6Ips []*string `json:"ipv6Ips,omitempty" tf:"ipv6_ips,omitempty"`

	// Collection of IP addresses assigned to a nic. Explicitly assigned public IPs need to come from reserved IP blocks, Passing value null or empty array will assign an IP address automatically.
	Ips []*string `json:"ips,omitempty" tf:"ips,omitempty"`

	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Lan
	Lan *float64 `json:"lan,omitempty" tf:"lan,omitempty"`

	// Reference to a Lan in compute to populate lan.
	// +kubebuilder:validation:Optional
	LanRef *v1.Reference `json:"lanRef,omitempty" tf:"-"`

	// Selector for a Lan in compute to populate lan.
	// +kubebuilder:validation:Optional
	LanSelector *v1.Selector `json:"lanSelector,omitempty" tf:"-"`

	// [string] The name of the server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type VCPUServerNicObservation struct {
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	DeviceNumber *float64 `json:"deviceNumber,omitempty" tf:"device_number,omitempty"`

	Dhcpv6 *bool `json:"dhcpv6,omitempty" tf:"dhcpv6,omitempty"`

	// Allows to define firewall rules inline in the server. See the Firewall section.
	// Firewall rules created in the server resource. The rules can also be created as separate resources outside the server resource
	Firewall []VCPUServerNicFirewallObservation `json:"firewall,omitempty" tf:"firewall,omitempty"`

	FirewallActive *bool `json:"firewallActive,omitempty" tf:"firewall_active,omitempty"`

	FirewallType *string `json:"firewallType,omitempty" tf:"firewall_type,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	IPv6Ips []*string `json:"ipv6Ips,omitempty" tf:"ipv6_ips,omitempty"`

	// Collection of IP addresses assigned to a nic. Explicitly assigned public IPs need to come from reserved IP blocks, Passing value null or empty array will assign an IP address automatically.
	Ips []*string `json:"ips,omitempty" tf:"ips,omitempty"`

	Lan *float64 `json:"lan,omitempty" tf:"lan,omitempty"`

	Mac *string `json:"mac,omitempty" tf:"mac,omitempty"`

	// [string] The name of the server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	PciSlot *float64 `json:"pciSlot,omitempty" tf:"pci_slot,omitempty"`
}

type VCPUServerNicParameters struct {

	// +kubebuilder:validation:Optional
	DHCP *bool `json:"dhcp,omitempty" tf:"dhcp,omitempty"`

	// +kubebuilder:validation:Optional
	Dhcpv6 *bool `json:"dhcpv6,omitempty" tf:"dhcpv6,omitempty"`

	// Allows to define firewall rules inline in the server. See the Firewall section.
	// Firewall rules created in the server resource. The rules can also be created as separate resources outside the server resource
	// +kubebuilder:validation:Optional
	Firewall []VCPUServerNicFirewallParameters `json:"firewall,omitempty" tf:"firewall,omitempty"`

	// +kubebuilder:validation:Optional
	FirewallActive *bool `json:"firewallActive,omitempty" tf:"firewall_active,omitempty"`

	// +kubebuilder:validation:Optional
	FirewallType *string `json:"firewallType,omitempty" tf:"firewall_type,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6CidrBlock *string `json:"ipv6CidrBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`

	// +kubebuilder:validation:Optional
	IPv6Ips []*string `json:"ipv6Ips,omitempty" tf:"ipv6_ips,omitempty"`

	// Collection of IP addresses assigned to a nic. Explicitly assigned public IPs need to come from reserved IP blocks, Passing value null or empty array will assign an IP address automatically.
	// +kubebuilder:validation:Optional
	Ips []*string `json:"ips,omitempty" tf:"ips,omitempty"`

	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Lan
	// +kubebuilder:validation:Optional
	Lan *float64 `json:"lan,omitempty" tf:"lan,omitempty"`

	// Reference to a Lan in compute to populate lan.
	// +kubebuilder:validation:Optional
	LanRef *v1.Reference `json:"lanRef,omitempty" tf:"-"`

	// Selector for a Lan in compute to populate lan.
	// +kubebuilder:validation:Optional
	LanSelector *v1.Selector `json:"lanSelector,omitempty" tf:"-"`

	// [string] The name of the server.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type VCPUServerObservation struct {

	// [string] The availability zone in which the server should exist. E.g: AUTO, ZONE_1, ZONE_2. This property is immutable.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// DEPRECATED Please refer to ionoscloud_server_boot_device_selection [string] The associated boot drive, if any. Must be the UUID of a bootable CDROM image that can be retrieved using the ionoscloud_image data source.
	// The associated boot drive, if any. Must be the UUID of a bootable CDROM image that can be retrieved using the ionoscloud_image data source.
	BootCdrom *string `json:"bootCdrom,omitempty" tf:"boot_cdrom,omitempty"`

	// [string] The image or snapshot UUID / name. May also be an image alias. It is required if licence_type is not provided.
	BootImage *string `json:"bootImage,omitempty" tf:"boot_image,omitempty"`

	// (Computed) The associated boot volume.
	BootVolume *string `json:"bootVolume,omitempty" tf:"boot_volume,omitempty"`

	CPUFamily *string `json:"cpuFamily,omitempty" tf:"cpu_family,omitempty"`

	// [integer] Number of server CPU cores.
	Cores *float64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// [string] The ID of a Virtual Data Center.
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// (Computed) The associated firewall rule.
	FirewallruleID *string `json:"firewallruleId,omitempty" tf:"firewallrule_id,omitempty"`

	// (Computed) The associated firewall rules.
	FirewallruleIds []*string `json:"firewallruleIds,omitempty" tf:"firewallrule_ids,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// [string] The name, ID or alias of the image. May also be a snapshot ID. It is required if licence_type is not provided. Attribute is immutable.
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// (Computed) A list with the IDs for the volumes that are defined inside the server resource.
	// A list that contains the IDs for the volumes defined inside the server resource.
	InlineVolumeIds []*string `json:"inlineVolumeIds,omitempty" tf:"inline_volume_ids,omitempty"`

	// A label can be seen as an object with only two required fields: key and value, both of the string type. Please check the example presented above to see how a label can be used in the plan. A server can have multiple labels.
	Label []VCPUServerLabelObservation `json:"label,omitempty" tf:"label,omitempty"`

	// [string] The name of the server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// See the Nic section.
	Nic *VCPUServerNicObservation `json:"nic,omitempty" tf:"nic,omitempty"`

	// (Computed) The associated IP address.
	PrimaryIP *string `json:"primaryIp,omitempty" tf:"primary_ip,omitempty"`

	// (Computed) The associated NIC.
	// Id of the primary network interface
	PrimaryNic *string `json:"primaryNic,omitempty" tf:"primary_nic,omitempty"`

	// [integer] The amount of memory for the server in MB.
	RAM *float64 `json:"ram,omitempty" tf:"ram,omitempty"`

	// [list] Immutable List of absolute or relative paths to files containing public SSH key that will be injected into IonosCloud provided Linux images. Also accepts ssh keys directly. Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key. This field may only be set in creation requests. When reading, it always returns null. SSH keys are only supported if a public Linux image is used for the volume creation. Does not support ~ expansion to homedir in the given path.
	// Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key. This field may only be set in creation requests. When reading, it always returns null. SSH keys are only supported if a public Linux image is used for the volume creation.
	SSHKeys []*string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Sets the power state of the vcpu server. Possible values: `RUNNING` or `SHUTOFF`.
	VMState *string `json:"vmState,omitempty" tf:"vm_state,omitempty"`

	// See the Volume section.
	Volume *VCPUServerVolumeObservation `json:"volume,omitempty" tf:"volume,omitempty"`
}

type VCPUServerParameters struct {

	// [string] The availability zone in which the server should exist. E.g: AUTO, ZONE_1, ZONE_2. This property is immutable.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// DEPRECATED Please refer to ionoscloud_server_boot_device_selection [string] The associated boot drive, if any. Must be the UUID of a bootable CDROM image that can be retrieved using the ionoscloud_image data source.
	// The associated boot drive, if any. Must be the UUID of a bootable CDROM image that can be retrieved using the ionoscloud_image data source.
	// +kubebuilder:validation:Optional
	BootCdrom *string `json:"bootCdrom,omitempty" tf:"boot_cdrom,omitempty"`

	// [string] The image or snapshot UUID / name. May also be an image alias. It is required if licence_type is not provided.
	// +kubebuilder:validation:Optional
	BootImage *string `json:"bootImage,omitempty" tf:"boot_image,omitempty"`

	// [integer] Number of server CPU cores.
	// +kubebuilder:validation:Optional
	Cores *float64 `json:"cores,omitempty" tf:"cores,omitempty"`

	// [string] The ID of a Virtual Data Center.
	// +crossplane:generate:reference:type=github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1.Datacenter
	// +kubebuilder:validation:Optional
	DatacenterID *string `json:"datacenterId,omitempty" tf:"datacenter_id,omitempty"`

	// Reference to a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDRef *v1.Reference `json:"datacenterIdRef,omitempty" tf:"-"`

	// Selector for a Datacenter in compute to populate datacenterId.
	// +kubebuilder:validation:Optional
	DatacenterIDSelector *v1.Selector `json:"datacenterIdSelector,omitempty" tf:"-"`

	// (Computed) The associated firewall rules.
	// +kubebuilder:validation:Optional
	FirewallruleIds []*string `json:"firewallruleIds,omitempty" tf:"firewallrule_ids,omitempty"`

	// [string] The name, ID or alias of the image. May also be a snapshot ID. It is required if licence_type is not provided. Attribute is immutable.
	// +kubebuilder:validation:Optional
	ImageName *string `json:"imageName,omitempty" tf:"image_name,omitempty"`

	// [string] The password for the image.
	// +kubebuilder:validation:Optional
	ImagePasswordSecretRef *v1.SecretKeySelector `json:"imagePasswordSecretRef,omitempty" tf:"-"`

	// A label can be seen as an object with only two required fields: key and value, both of the string type. Please check the example presented above to see how a label can be used in the plan. A server can have multiple labels.
	// +kubebuilder:validation:Optional
	Label []VCPUServerLabelParameters `json:"label,omitempty" tf:"label,omitempty"`

	// [string] The name of the server.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// See the Nic section.
	// +kubebuilder:validation:Optional
	Nic *VCPUServerNicParameters `json:"nic,omitempty" tf:"nic,omitempty"`

	// [integer] The amount of memory for the server in MB.
	// +kubebuilder:validation:Optional
	RAM *float64 `json:"ram,omitempty" tf:"ram,omitempty"`

	// [list] Immutable List of absolute or relative paths to files containing public SSH key that will be injected into IonosCloud provided Linux images. Also accepts ssh keys directly. Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key. This field may only be set in creation requests. When reading, it always returns null. SSH keys are only supported if a public Linux image is used for the volume creation. Does not support ~ expansion to homedir in the given path.
	// Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key. This field may only be set in creation requests. When reading, it always returns null. SSH keys are only supported if a public Linux image is used for the volume creation.
	// +kubebuilder:validation:Optional
	SSHKeys []*string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`

	// Sets the power state of the vcpu server. Possible values: `RUNNING` or `SHUTOFF`.
	// +kubebuilder:validation:Optional
	VMState *string `json:"vmState,omitempty" tf:"vm_state,omitempty"`

	// See the Volume section.
	// +kubebuilder:validation:Optional
	Volume *VCPUServerVolumeParameters `json:"volume,omitempty" tf:"volume,omitempty"`
}

type VCPUServerVolumeInitParameters struct {

	// [string] The availability zone in which the server should exist. E.g: AUTO, ZONE_1, ZONE_2. This property is immutable.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The uuid of the Backup Unit that user has access to. The property is immutable and is only allowed to be set on a new volume creation. It is mandatory to provide either 'public image' or 'imageAlias' in conjunction with this property.
	BackupUnitID *string `json:"backupUnitId,omitempty" tf:"backup_unit_id,omitempty"`

	Bus *string `json:"bus,omitempty" tf:"bus,omitempty"`

	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// [string] Sets the OS type of the server.
	LicenceType *string `json:"licenceType,omitempty" tf:"licence_type,omitempty"`

	// [string] The name of the server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The size of the volume in GB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The cloud-init configuration for the volume as base64 encoded string. The property is immutable and is only allowed to be set on a new volume creation. It is mandatory to provide either 'public image' or 'imageAlias' that has cloud-init compatibility in conjunction with this property.
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`
}

type VCPUServerVolumeObservation struct {

	// [string] The availability zone in which the server should exist. E.g: AUTO, ZONE_1, ZONE_2. This property is immutable.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The uuid of the Backup Unit that user has access to. The property is immutable and is only allowed to be set on a new volume creation. It is mandatory to provide either 'public image' or 'imageAlias' in conjunction with this property.
	BackupUnitID *string `json:"backupUnitId,omitempty" tf:"backup_unit_id,omitempty"`

	// The UUID of the attached server.
	BootServer *string `json:"bootServer,omitempty" tf:"boot_server,omitempty"`

	Bus *string `json:"bus,omitempty" tf:"bus,omitempty"`

	CPUHotPlug *bool `json:"cpuHotPlug,omitempty" tf:"cpu_hot_plug,omitempty"`

	DeviceNumber *float64 `json:"deviceNumber,omitempty" tf:"device_number,omitempty"`

	DiscVirtioHotPlug *bool `json:"discVirtioHotPlug,omitempty" tf:"disc_virtio_hot_plug,omitempty"`

	DiscVirtioHotUnplug *bool `json:"discVirtioHotUnplug,omitempty" tf:"disc_virtio_hot_unplug,omitempty"`

	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// [string] Sets the OS type of the server.
	LicenceType *string `json:"licenceType,omitempty" tf:"licence_type,omitempty"`

	// [string] The name of the server.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	NicHotPlug *bool `json:"nicHotPlug,omitempty" tf:"nic_hot_plug,omitempty"`

	NicHotUnplug *bool `json:"nicHotUnplug,omitempty" tf:"nic_hot_unplug,omitempty"`

	PciSlot *float64 `json:"pciSlot,omitempty" tf:"pci_slot,omitempty"`

	RAMHotPlug *bool `json:"ramHotPlug,omitempty" tf:"ram_hot_plug,omitempty"`

	// The size of the volume in GB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The cloud-init configuration for the volume as base64 encoded string. The property is immutable and is only allowed to be set on a new volume creation. It is mandatory to provide either 'public image' or 'imageAlias' that has cloud-init compatibility in conjunction with this property.
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`
}

type VCPUServerVolumeParameters struct {

	// [string] The availability zone in which the server should exist. E.g: AUTO, ZONE_1, ZONE_2. This property is immutable.
	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The uuid of the Backup Unit that user has access to. The property is immutable and is only allowed to be set on a new volume creation. It is mandatory to provide either 'public image' or 'imageAlias' in conjunction with this property.
	// +kubebuilder:validation:Optional
	BackupUnitID *string `json:"backupUnitId,omitempty" tf:"backup_unit_id,omitempty"`

	// +kubebuilder:validation:Optional
	Bus *string `json:"bus,omitempty" tf:"bus,omitempty"`

	// +kubebuilder:validation:Optional
	DiskType *string `json:"diskType" tf:"disk_type,omitempty"`

	// [string] Sets the OS type of the server.
	// +kubebuilder:validation:Optional
	LicenceType *string `json:"licenceType,omitempty" tf:"licence_type,omitempty"`

	// [string] The name of the server.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The size of the volume in GB.
	// +kubebuilder:validation:Optional
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The cloud-init configuration for the volume as base64 encoded string. The property is immutable and is only allowed to be set on a new volume creation. It is mandatory to provide either 'public image' or 'imageAlias' that has cloud-init compatibility in conjunction with this property.
	// +kubebuilder:validation:Optional
	UserData *string `json:"userData,omitempty" tf:"user_data,omitempty"`
}

// VCPUServerSpec defines the desired state of VCPUServer
type VCPUServerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VCPUServerParameters `json:"forProvider"`
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
	InitProvider VCPUServerInitParameters `json:"initProvider,omitempty"`
}

// VCPUServerStatus defines the observed state of VCPUServer.
type VCPUServerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VCPUServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VCPUServer is the Schema for the VCPUServers API. Creates and manages IonosCloud VCPU Server objects.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,ionos}
type VCPUServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cores) || (has(self.initProvider) && has(self.initProvider.cores))",message="spec.forProvider.cores is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.ram) || (has(self.initProvider) && has(self.initProvider.ram))",message="spec.forProvider.ram is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.volume) || (has(self.initProvider) && has(self.initProvider.volume))",message="spec.forProvider.volume is a required parameter"
	Spec   VCPUServerSpec   `json:"spec"`
	Status VCPUServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VCPUServerList contains a list of VCPUServers
type VCPUServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VCPUServer `json:"items"`
}

// Repository type metadata.
var (
	VCPUServer_Kind             = "VCPUServer"
	VCPUServer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VCPUServer_Kind}.String()
	VCPUServer_KindAPIVersion   = VCPUServer_Kind + "." + CRDGroupVersion.String()
	VCPUServer_GroupVersionKind = CRDGroupVersion.WithKind(VCPUServer_Kind)
)

func init() {
	SchemeBuilder.Register(&VCPUServer{}, &VCPUServerList{})
}
