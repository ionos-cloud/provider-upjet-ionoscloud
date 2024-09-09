/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"context"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/registry/reference"
	"github.com/crossplane/upjet/pkg/schema/traverser"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"

	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"github.com/ionos-cloud/terraform-provider-ionoscloud/v6/xpprovider"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/apigateway"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/apigatewayroute"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/crossconnect"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/datacenter"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/inmemorydb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/ipblock"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/ipsec"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/k8s"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/kafka"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/lan"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/mariadb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/nfs"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/nic"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/s3"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/server"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/user"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/volume"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/wireguard"
)

const (
	resourcePrefix = "ionoscloud"
	modulePath     = "github.com/ionos-cloud/provider-upjet-ionoscloud"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider(ctx context.Context) (*ujconfig.Provider, error) {
	fwProvider, sdkProvider := xpprovider.GetProvider()
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

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithShortName("ionos"),
		ujconfig.WithRootGroup("ionoscloud.io"),
		ujconfig.WithIncludeList(CLIReconciledResourceList()),
		ujconfig.WithTerraformPluginSDKIncludeList(TerraformPluginSDKResourceList()),
		ujconfig.WithTerraformPluginFrameworkIncludeList(TerraformPluginFrameworkResourceList()),
		ujconfig.WithReferenceInjectors([]config.ReferenceInjector{reference.NewInjector(modulePath)}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithTerraformProvider(sdkProvider),
		ujconfig.WithTerraformPluginFrameworkProvider(fwProvider),
		ujconfig.WithSchemaTraversers(&config.SingletonListEmbedder{}),
		ujconfig.WithDefaultResourceOptions(
			ResourceConfigurator(),
		),
	)

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		s3.Configure,
		datacenter.Configure,
		lan.Configure,
		crossconnect.Configure,
		apigateway.Configure,
		apigatewayroute.Configure,
		mariadb.Configure,
		ipblock.Configure,
		wireguard.Configure,
		kafka.Configure,
		inmemorydb.Configure,
		server.Configure,
		ipsec.Configure,
		nfs.Configure,
		volume.Configure,
		k8s.Configure,
		user.Configure,
		nic.Configure,
	} {
		configure(pc)
	}
	pc.ConfigureResources()
	return pc, nil
}

func getProviderSchema(s string) (*schema.Provider, error) {
	ps := tfjson.ProviderSchemas{}
	if err := ps.UnmarshalJSON([]byte(s)); err != nil {
		panic(err)
	}
	if len(ps.Schemas) != 1 {
		return nil, errors.Errorf("there should exactly be 1 provider schema but there are %d", len(ps.Schemas))
	}
	var rs map[string]*tfjson.Schema
	for _, v := range ps.Schemas {
		rs = v.ResourceSchemas
		break
	}
	return &schema.Provider{
		ResourcesMap: conversiontfjson.GetV2ResourceMap(rs),
	}, nil
}
