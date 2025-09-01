package common

import (
	"errors"

	"github.com/crossplane/crossplane-runtime/v2/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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

// ServerPrimaryNIC returns the primary NIC of a server to be used for cross-referencing
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

// AutoCertificateProviderLocation returns the location of a certificate provider to be used for cross-referencing
func AutoCertificateProviderLocation() reference.ExtractValueFn {
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

// IgnoreEmptyDiffForComputed gets a list of fields on which to ignore the diff from terraform. Computed arguments can throw this error when
// the diff is empty. This should probably be solved by the generator
func IgnoreEmptyDiffForComputed(ignoreList []string) func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
	return func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
		// skip diff customization on create
		if state == nil || state.Empty() {
			return diff, nil
		}
		if config == nil {
			return nil, errors.New("resource config cannot be nil")
		}
		// skip no diff or destroy diffs
		if diff == nil || diff.Empty() || diff.Destroy || diff.Attributes == nil {
			return diff, nil
		}

		for _, key := range ignoreList {
			if diff.Attributes[key] != nil && diff.Attributes[key].Old == "" && diff.Attributes[key].New == "" {
				delete(diff.Attributes, key)
			}
		}

		return diff, nil
	}
}
