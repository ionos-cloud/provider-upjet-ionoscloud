// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	"context"
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/registry/reference"
	"github.com/crossplane/upjet/v2/pkg/schema/traverser"
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster"
)

// GetProvider returns provider configuration
// func GetProvider(generationProvider bool) (*config.Provider, error) {
// 	fwProvider, sdkProvider := xpprovider.GetProvider()
// 	if generationProvider {
// 		p, err := getProviderSchema(providerSchema)
// 		if err != nil {
// 			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
// 		}
//
// 		if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
// 			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
// 		}
//
// 		sdkProvider = p
// 	}
// 	modulePath := "github.com/ionos-cloud/provider-upjet-ionoscloud"
// 	resourcePrefix := "ionoscloud"
// 	pc := config.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
// 		config.WithShortName("ionos"),
// 		config.WithRootGroup("ionoscloud.io"),
// 		config.WithIncludeList(CLIReconciledResourceList()),
// 		config.WithTerraformPluginSDKIncludeList(TerraformPluginSDKResourceList()),
// 		config.WithTerraformPluginFrameworkIncludeList(TerraformPluginFrameworkResourceList()),
// 		config.WithReferenceInjectors([]config.ReferenceInjector{reference.NewInjector(modulePath)}),
// 		config.WithFeaturesPackage("internal/features"),
// 		config.WithTerraformProvider(sdkProvider),
// 		config.WithTerraformPluginFrameworkProvider(fwProvider),
// 		config.WithSchemaTraversers(&config.SingletonListEmbedder{}),
// 		config.WithDefaultResourceOptions(
// 			ResourceConfigurator(),
// 		),
// 	)
//
// 	for _, configure := range cluster.ProviderConfiguration {
// 		configure(pc)
// 	}
// 	pc.ConfigureResources()
// 	addTFSingletonConversion(pc)
// 	return pc, nil
// }

// GetProvider returns provider configuration
func GetProvider(ctx context.Context, fwProvider fwprovider.Provider, sdkProvider *schema.Provider, generationProvider bool, skipDefaultTags bool) (*config.Provider, error) {
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
	pc := config.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		config.WithShortName("ionos"),
		config.WithRootGroup("ionoscloud.io"),
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

	for _, configure := range cluster.ProviderConfiguration {
		configure(pc)
	}
	pc.ConfigureResources()
	addTFSingletonConversion(pc)
	return pc, nil
}
