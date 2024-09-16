package common

import (
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

// FirstIPBlockIP returns the first IP of an ipblock to assign to a wireguard gateway
func FirstIPBlockIP() reference.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		obj, err := fieldpath.PaveObject(mr)
		if err != nil {
			return ""
		}

		result, err := obj.GetString("status.atProvider.ips[0]")
		if err != nil {
			return ""
		}

		return result
	}
}

// DatacenterLocation returns the location of a datacenter to be used for cross-referencing
func DatacenterLocation() reference.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		obj, err := fieldpath.PaveObject(mr)
		if err != nil {
			return ""
		}

		result, err := obj.GetString("status.atProvider.location")
		if err != nil {
			return ""
		}

		return result
	}
}

func ServerPrimaryNIC() reference.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		obj, err := fieldpath.PaveObject(mr)
		if err != nil {
			return ""
		}

		result, err := obj.GetString("status.atProvider.primaryNic")
		if err != nil {
			return ""
		}

		return result
	}
}
