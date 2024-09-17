/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/registry/reference"
	"github.com/crossplane/upjet/pkg/schema/traverser"
	conversiontfjson "github.com/crossplane/upjet/pkg/types/conversion/tfjson"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	asg "github.com/ionos-cloud/provider-upjet-ionoscloud/config/alb"
	alb "github.com/ionos-cloud/provider-upjet-ionoscloud/config/asg"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/backupunit"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/certificatemanager"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/compute"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/containerregistry"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/dataplatform"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/log"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/natgateway"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/nlb"
	"github.com/pkg/errors"

	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/ionos-cloud/terraform-provider-ionoscloud/v6/xpprovider"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/apigateway"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/dns"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/inmemorydb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/ipsec"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/k8s"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/kafka"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/mariadb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/mongodb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/nfs"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/postgresql"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/s3"
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
func GetProvider(generationProvider bool) (*ujconfig.Provider, error) {
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
		backupunit.Configure,
		alb.Configure,
		certificatemanager.Configure,
		asg.Configure,
		compute.Configure,
		nlb.Configure,
		containerregistry.Configure,
		natgateway.Configure,
		dataplatform.Configure,
		log.Configure,
	} {
		configure(pc)
	}
	pc.ConfigureResources()
	addTFSingletonConversion(pc)
	return pc, nil
}

func addTFSingletonConversion(pc *config.Provider) {
	for name, r := range pc.Resources {
		r := r
		// nothing to do if no singleton list has been converted to
		// an embedded object
		if len(r.CRDListConversionPaths()) == 0 {
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
		}

		pc.Resources[name] = r
	}
}

func hasSingletonListCredentialsSensitiveMarshallingProblems(r *config.Resource) bool {
	return r.Name == "ionoscloud_inmemorydb_replicaset" ||
		r.Name == "ionoscloud_pg_cluster" ||
		r.Name == "ionoscloud_vpn_ipsec_tunnel" ||
		r.Name == "ionoscloud_mariadb_cluster" ||
		r.Name == "ionoscloud_mongo_cluster" ||
		r.Name == "ionoscloud_autoscaling_group"
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
