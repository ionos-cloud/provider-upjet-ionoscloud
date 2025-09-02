// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"context"
	_ "embed"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/registry/reference"
	"github.com/crossplane/upjet/v2/pkg/schema/traverser"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced"
)

const (
	resourcePrefix = "ionoscloud"
	modulePath     = "github.com/ionos-cloud/provider-upjet-ionoscloud"
)

func addTFSingletonConversion(pc *config.Provider) {
	for name, r := range pc.Resources {
		r := r
		// nothing to do if no singleton list has been converted to
		// an embedded object
		if len(r.CRDListConversionPaths()) == 0 {
			continue
		}

		if r.ShouldUseTerraformPluginFrameworkClient() {
			continue
		}

		// the controller will be reconciling on the CRD API version
		// with the converted API (with embedded objects in place of
		// singleton lists), so we need the appropriate Terraform
		// converter in this case.
		r.TerraformConversions = []config.TerraformConversion{
			config.NewTFSingletonConversion(),
		}

		// Workaround - disable singleton list conversion for sensitive fields
		// There is a problem with expanding singleton lists wildcards
		// See issue https://github.com/crossplane/upjet/issues/436
		if hasSingletonListCredentialsSensitiveMarshallingProblems(r) {
			r.RemoveSingletonListConversion("credentials")
			r.RemoveSingletonListConversion("auth")
			r.RemoveSingletonListConversion("replica_configuration")
			r.RemoveSingletonListConversion("external_account_binding")
		}

		pc.Resources[name] = r
	}
}

// GetProviderNamespaced returns the provider configuration.
// The `generationProvider` argument specifies whether the provider
// configuration is being read for the code generation pipelines.
// In that case, we will only use the JSON schema for generating
// the CRDs.
func GetProviderNamespaced(ctx context.Context, fwProvider fwprovider.Provider, sdkProvider *schema.Provider, generationProvider bool) (*config.Provider, error) {
	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
		}
		if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
		}
		// use the JSON schema to temporarily prevent float64->int64
		// conversions in the CRD APIs.
		// We would like to convert to int64s with the next major release of
		// the provider.
		sdkProvider = p
	}

	modulePath := "github.com/ionos-cloud/provider-upjet-ionoscloud"
	resourcePrefix := "ionoscloud"
	pc := config.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, providerMetadata,
		config.WithShortName("ionos"),
		config.WithRootGroup("m.ionoscloud.io"),
		config.WithIncludeList(CLIReconciledResourceList()),
		config.WithTerraformPluginSDKIncludeList(TerraformPluginSDKResourceList()),
		config.WithTerraformPluginFrameworkIncludeList(TerraformPluginFrameworkResourceList()),
		config.WithReferenceInjectors([]config.ReferenceInjector{reference.NewInjector(modulePath)}),
		config.WithFeaturesPackage("internal/features"),
		config.WithTerraformProvider(sdkProvider),
		config.WithTerraformPluginFrameworkProvider(fwProvider),
		config.WithSchemaTraversers(&config.SingletonListEmbedder{}),
		config.WithDefaultResourceOptions(
			ResourceConfigurator(),
		),
	)

	for _, configure := range namespaced.ProviderConfiguration {
		configure(pc)
	}
	pc.ConfigureResources()
	addTFSingletonConversion(pc)
	return pc, nil
}

func hasSingletonListCredentialsSensitiveMarshallingProblems(r *config.Resource) bool {
	return r.Name == "ionoscloud_inmemorydb_replicaset" ||
		r.Name == "ionoscloud_pg_cluster" ||
		r.Name == "ionoscloud_vpn_ipsec_tunnel" ||
		r.Name == "ionoscloud_mariadb_cluster" ||
		r.Name == "ionoscloud_mongo_cluster" ||
		r.Name == "ionoscloud_autoscaling_group" ||
		r.Name == "ionoscloud_auto_certificate_provider"
}
