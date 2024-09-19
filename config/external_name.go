/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Version is the version of the resources to be reconciled.
const Version = "v1alpha1"

// TerraformPluginSDKExternalNameConfigs contains all external name configurations
// belonging to Terraform Plugin SDKv2 resources to be reconciled
// under the no-fork architecture for this provider.
var TerraformPluginSDKExternalNameConfigs = map[string]config.ExternalName{
	"ionoscloud_datacenter":            config.IdentifierFromProvider,
	"ionoscloud_private_crossconnect":  config.IdentifierFromProvider,
	"ionoscloud_lan":                   config.IdentifierFromProvider,
	"ionoscloud_apigateway":            config.IdentifierFromProvider,
	"ionoscloud_apigateway_route":      config.IdentifierFromProvider,
	"ionoscloud_mariadb_cluster":       config.IdentifierFromProvider,
	"ionoscloud_ipblock":               config.IdentifierFromProvider,
	"ionoscloud_vpn_wireguard_gateway": config.IdentifierFromProvider,
	"ionoscloud_kafka_cluster":         config.IdentifierFromProvider,
	"ionoscloud_inmemorydb_replicaset": config.IdentifierFromProvider,
	"ionoscloud_vpn_wireguard_peer":    config.IdentifierFromProvider,
	"ionoscloud_kafka_cluster_topic":   config.IdentifierFromProvider,
	"ionoscloud_server":                config.IdentifierFromProvider,
	"ionoscloud_volume":                config.IdentifierFromProvider,
	"ionoscloud_vpn_ipsec_gateway":     config.IdentifierFromProvider,
	"ionoscloud_vpn_ipsec_tunnel":      config.IdentifierFromProvider,
	"ionoscloud_nfs_cluster":           config.IdentifierFromProvider,
	"ionoscloud_nfs_share":             config.IdentifierFromProvider,
	"ionoscloud_k8s_cluster":           config.IdentifierFromProvider,
	"ionoscloud_k8s_node_pool":         config.IdentifierFromProvider,
	"ionoscloud_user":                  config.IdentifierFromProvider,
	"ionoscloud_group":                 config.IdentifierFromProvider,
	"ionoscloud_nic":                   config.IdentifierFromProvider,
	"ionoscloud_s3_key":                config.IdentifierFromProvider,
	"ionoscloud_dns_zone":              config.IdentifierFromProvider,
	"ionoscloud_dns_record":            config.IdentifierFromProvider,
	"ionoscloud_mongo_cluster":         config.IdentifierFromProvider,
	"ionoscloud_mongo_user":            config.IdentifierFromProvider,
	"ionoscloud_pg_cluster":            config.IdentifierFromProvider,
	"ionoscloud_pg_user":               config.IdentifierFromProvider,
	"ionoscloud_pg_database":           config.IdentifierFromProvider,
}

// TerraformPluginFrameworkExternalNameConfigs will be used for plugin configured resources. not used yet
var TerraformPluginFrameworkExternalNameConfigs = map[string]config.ExternalName{
	"ionoscloud_s3_bucket":                                      config.NameAsIdentifier,
	"ionoscloud_s3_bucket_policy":                               config.IdentifierFromProvider,
	"ionoscloud_s3_bucket_versioning":                           config.IdentifierFromProvider,
	"ionoscloud_s3_bucket_cors_configuration":                   config.IdentifierFromProvider,
	"ionoscloud_s3_bucket_lifecycle_configuration":              config.IdentifierFromProvider,
	"ionoscloud_s3_bucket_public_access_block":                  config.IdentifierFromProvider,
	"ionoscloud_s3_bucket_website_configuration":                config.IdentifierFromProvider,
	"ionoscloud_s3_bucket_object_lock_configuration":            config.IdentifierFromProvider,
	"ionoscloud_s3_bucket_server_side_encryption_configuration": config.IdentifierFromProvider,
	"ionoscloud_s3_object":                                      config.IdentifierFromProvider,
	"ionoscloud_s3_object_copy":                                 config.IdentifierFromProvider,
}

// ResourceConfigurator applies all external name configs listed in
// the table TerraformPluginSDKExternalNameConfigs,
// CLIReconciledExternalNameConfigs, and
// TerraformPluginFrameworkExternalNameConfigs and sets the version of
// those resources to v1beta1.
func ResourceConfigurator() config.ResourceOption {
	return func(r *config.Resource) {
		// If an external name is configured for multiple architectures,
		// Terraform Plugin Framework takes precedence over Terraform
		// Plugin SDKv2, which takes precedence over CLI architecture.
		e, configured := TerraformPluginFrameworkExternalNameConfigs[r.Name]
		if !configured {
			e, configured = TerraformPluginSDKExternalNameConfigs[r.Name]
			if !configured {
				e, configured = CLIReconciledExternalNameConfigs[r.Name]
			}
		}
		if !configured {
			return
		}
		r.Version = Version
		r.ExternalName = e
	}
}

// TerraformPluginSDKResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the no-fork architecture.
func TerraformPluginSDKResourceList() []string {
	l := make([]string, len(TerraformPluginSDKExternalNameConfigs))
	i := 0
	for name := range TerraformPluginSDKExternalNameConfigs {
		// Expected format is regex, and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

// TerraformPluginFrameworkResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the TF Plugin Framework architecture.
func TerraformPluginFrameworkResourceList() []string {
	l := make([]string, len(TerraformPluginFrameworkExternalNameConfigs))
	i := 0
	for name := range TerraformPluginFrameworkExternalNameConfigs {
		// Expected format is regex, and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

// CLIReconciledResourceList returns the list of resources that have external
// name configured in ExternalNameConfigs table and to be reconciled under
// the TF CLI based architecture. Not used yet
func CLIReconciledResourceList() []string {
	l := make([]string, len(CLIReconciledExternalNameConfigs))
	i := 0
	for name := range CLIReconciledExternalNameConfigs {
		// Expected format is regex, and we'd like to have exact matches.
		l[i] = name + "$"
		i++
	}
	return l
}

// CLIReconciledExternalNameConfigs will be used for CLI reconciled resources. not used yet
var CLIReconciledExternalNameConfigs = map[string]config.ExternalName{}
