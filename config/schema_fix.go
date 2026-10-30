// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// fixRequiredComputedFields clears Computed on every field that is both Required and Computed.
//
// upjet converts Terraform JSON schema blocks with min_items == max_items == 0 to Computed
// (pkg/types/conversion/tfjson.tfJSONBlockTypeToV2Schema). For single nested blocks with a
// required child it additionally sets Required, but keeps Computed. upjet's IsObservation()
// treats "Computed && !Optional" as an observation-only field, so such blocks are generated
// into status.atProvider only and cannot be set, e.g. versioning_configuration of
// ionoscloud_s3_bucket_versioning. A field cannot be required and computed at the same time
// in the Terraform SDK v2 schema model, so Required wins.
func fixRequiredComputedFields(pc *config.Provider) {
	for _, r := range pc.Resources {
		clearRequiredComputed(r.TerraformResource.Schema)
	}
}

func clearRequiredComputed(m map[string]*schema.Schema) {
	for _, s := range m {
		if s.Required && s.Computed {
			s.Computed = false
		}
		if r, ok := s.Elem.(*schema.Resource); ok {
			clearRequiredComputed(r.Schema)
		}
	}
}
