// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/registry/reference"
	"github.com/crossplane/upjet/v2/pkg/schema/traverser"
	"github.com/pkg/errors"

	"github.com/ionos-cloud/terraform-provider-ionoscloud/v6/xpprovider"

	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/alb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/apigateway"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/asg"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/cdn"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/certificatemanager"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/compute"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/containerregistry"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/dataplatform"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/dns"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/inmemorydb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/ipsec"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/k8s"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/kafka"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/log"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/mariadb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/mongodb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/natgateway"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/nfs"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/nlb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/objectstorage"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/postgresql"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/wireguard"
)

// GetProvider returns provider configuration
func GetProvider(generationProvider bool) (*config.Provider, error) {
	fwProvider, sdkProvider := xpprovider.GetProvider()
	if generationProvider {
		p, err := getProviderSchema(providerSchema)
		if err != nil {
			return nil, errors.Wrap(err, "cannot read the Terraform SDK provider from the JSON schema for code generation")
		}

		if err := traverser.TFResourceSchema(sdkProvider.ResourcesMap).Traverse(traverser.NewMaxItemsSync(p.ResourcesMap)); err != nil {
			return nil, errors.Wrap(err, "cannot sync the MaxItems constraints between the Go schema and the JSON schema")
		}

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

	for _, configure := range []func(provider *config.Provider){
		// add custom config functions
		objectstorage.Configure,
		compute.Configure,
		apigateway.Configure,
		mariadb.Configure,
		wireguard.Configure,
		kafka.Configure,
		inmemorydb.Configure,
		ipsec.Configure,
		nfs.Configure,
		k8s.Configure,
		dns.Configure,
		mongodb.Configure,
		postgresql.Configure,
		alb.Configure,
		certificatemanager.Configure,
		asg.Configure,
		compute.Configure,
		nlb.Configure,
		containerregistry.Configure,
		natgateway.Configure,
		dataplatform.Configure,
		log.Configure,
		cdn.Configure,
	} {
		configure(pc)
	}
	pc.ConfigureResources()
	addTFSingletonConversion(pc)
	return pc, nil
}
