/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"ionoscloud_datacenter":            config.IdentifierFromProvider,
	"ionoscloud_private_crossconnect":  config.IdentifierFromProvider,
	"ionoscloud_lan":                   config.IdentifierFromProvider,
	"ionoscloud_apigateway":            config.IdentifierFromProvider,
	"ionoscloud_apigateway_route":      config.IdentifierFromProvider,
	"ionoscloud_s3_bucket":             config.NameAsIdentifier,
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
}

// TerraformPluginFrameworkExternalNameConfigs will be used for plugin configured resources. not used yet
var TerraformPluginFrameworkExternalNameConfigs = map[string]config.ExternalName{
	"ionoscloud_s3_bucket": config.NameAsIdentifier,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
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

// TerraformPluginFrameworkResourceList returns the list of plugin resources. unused at the moment
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
