package mariadb

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_mariadb_cluster", func(r *config.Resource) {
		r.ShortGroup = "dbaas.mariadb"
		r.Kind = "MariadbCluster"
		r.References["connections.datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["connections.lan_id"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
		r.UseAsync = true
	})
}
