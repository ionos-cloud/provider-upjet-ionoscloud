// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ionosterraform "github.com/ionos-cloud/terraform-provider-ionoscloud/v6/xpprovider"
)

func collectRequiredComputed(prefix string, m map[string]*schema.Schema, out *[]string) {
	for k, s := range m {
		p := prefix + k
		if s.Required && s.Computed {
			*out = append(*out, p)
		}
		if r, ok := s.Elem.(*schema.Resource); ok {
			collectRequiredComputed(p+".", r.Schema, out)
		}
	}
}

// Required+Computed fields end up in status.atProvider only (upjet IsObservation),
// so they cannot be set. See fixRequiredComputedFields.
func TestNoRequiredComputedFields(t *testing.T) {
	fw, sdk := ionosterraform.GetProvider()
	providers := map[string]func(bool) (*config.Provider, error){
		"cluster":    func(gen bool) (*config.Provider, error) { return GetProvider(fw, sdk, gen) },
		"namespaced": func(gen bool) (*config.Provider, error) { return GetProviderNamespaced(fw, sdk, gen) },
	}
	for name, get := range providers {
		for _, gen := range []bool{true, false} {
			pc, err := get(gen)
			if err != nil {
				t.Fatalf("%s (generation=%v): %v", name, gen, err)
			}
			for rn, r := range pc.Resources {
				var hits []string
				collectRequiredComputed("", r.TerraformResource.Schema, &hits)
				for _, h := range hits {
					t.Errorf("%s (generation=%v): %s.%s is Required and Computed", name, gen, rn, h)
				}
			}
		}
	}
}

func TestS3BucketVersioningConfigurationIsSettable(t *testing.T) {
	fw, sdk := ionosterraform.GetProvider()
	pc, err := GetProviderNamespaced(fw, sdk, true)
	if err != nil {
		t.Fatal(err)
	}
	s := pc.Resources["ionoscloud_s3_bucket_versioning"].TerraformResource.Schema["versioning_configuration"]
	if s == nil || !s.Required || s.Computed {
		t.Fatalf("versioning_configuration must be Required and not Computed, got %+v", s)
	}
}
