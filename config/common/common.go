package common

import (
	"errors"
	"strings"

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

// IgnoreDiffForFields removes specific fields from Terraform diffs. This is
// useful for immutable attributes that may appear as provider-side drift and
// should not trigger update planning.
// todo: use this as a last way to fix a diff problem, as it can hide real issues if used excessively.
// Always try to solve diff problems by fixxing the issue in terraform, or in the generator first, and use this only if there is no other way to fix a diff problem.
func IgnoreDiffForFields(ignoreList []string) func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
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
			delete(diff.Attributes, key)
		}

		return diff, nil
	}
}

// IgnoreDiffForUnconfiguredComputedField ignores diff entries for a computed field (e.g. "ips")
// only when that field was NOT explicitly set in the resource's configuration. Some computed
// attributes reflect provider-managed state that the API can legitimately change on its own over
// time (e.g. a DHCP-auto-assigned public IP) — once late-init copies the first observed value
// into spec, continuous reconciliation treats it as a desired value and perpetually tries to
// force reality back to that now-stale snapshot, which the API may not honor, causing an endless
// update loop. When the field IS explicitly configured, its diff is left untouched so an
// intentionally desired value (e.g. a genuinely reserved IP) is still enforced normally.
func IgnoreDiffForUnconfiguredComputedField(field string) func(diff *terraform.InstanceDiff, state *terraform.InstanceState, config *terraform.ResourceConfig) (*terraform.InstanceDiff, error) {
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

		// GetRaw only succeeds against the user-supplied configuration, never a computed
		// placeholder, so this is false whenever the field was left entirely unset.
		if _, configured := config.GetRaw(field); configured {
			return diff, nil
		}

		prefix := field + "."
		for key := range diff.Attributes {
			if key == field || strings.HasPrefix(key, prefix) {
				delete(diff.Attributes, key)
			}
		}

		return diff, nil
	}
}
