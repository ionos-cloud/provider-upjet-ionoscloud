package nfs

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator(
		"ionoscloud_nfs_cluster", func(r *config.Resource) {
			r.ShortGroup = "nfs"
			r.Kind = "NFSCluster"
			r.References["connections.datacenter_id"] = config.Reference{
				TerraformName: "ionoscloud_datacenter",
			}
			r.References["connections.lan"] = config.Reference{
				TerraformName: "ionoscloud_lan",
			}
			r.ExternalName = config.IdentifierFromProvider
			r.UseAsync = true
		},
	)

	p.AddResourceConfigurator(
		"ionoscloud_nfs_share", func(r *config.Resource) {
			r.ShortGroup = "nfs"
			r.Kind = "NFSShare"
			r.References["cluster_id"] = config.Reference{
				TerraformName: "ionoscloud_nfs_cluster",
			}
			r.ExternalName = config.IdentifierFromProvider
			r.UseAsync = true
		},
	)
}
