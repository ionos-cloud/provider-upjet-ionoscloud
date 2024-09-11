/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

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
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/mongodb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/nfs"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/nic"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/postgresql"
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
func GetProvider() *ujconfig.Provider {
	// provider := tf.Provider()
	// fwProvider := tf.FWProvider()
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("ionoscloud.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithShortName("ionos"),
		ujconfig.WithFeaturesPackage("internal/features"),

		// ujconfig.WithIncludeList(CLIReconciledResourceList()),
		// ujconfig.WithTerraformProvider(provider),
		// ujconfig.WithTerraformPluginFrameworkProvider(fwProvider),
		// ujconfig.WithTerraformPluginSDKIncludeList(ExternalNameConfigured()),
		// ujconfig.WithTerraformPluginFrameworkIncludeList(TerraformPluginFrameworkResourceList()),

		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
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
		mongodb.Configure,
		postgresql.Configure,
	} {
		configure(pc)
	}
	pc.ConfigureResources()
	return pc
}
