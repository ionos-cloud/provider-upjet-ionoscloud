package common

import (
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1"
)

// FirstIPBlockIP returns the first IP of an ipblock to assign to a wireguard gateway
func FirstIPBlockIP() reference.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(*v1alpha1.Ipblock)
		if !ok {
			return ""
		}
		if tr.Status.AtProvider.ID == nil || tr.Status.AtProvider.Ips == nil {
			return ""
		}
		return *tr.Status.AtProvider.Ips[0]
	}
}

// DatacenterLocation returns the location of a datacenter to be used for cross-referencing
func DatacenterLocation() reference.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(*v1alpha1.Datacenter)
		if !ok {
			return ""
		}
		if tr.Status.AtProvider.ID == nil || tr.Status.AtProvider.Location == nil {
			return ""
		}
		return *tr.Status.AtProvider.Location
	}
}
