package postgresql

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator(
		"ionoscloud_pg_cluster", func(r *config.Resource) {
			r.ShortGroup = "postgresql"
			r.Kind = "PostgresqlCluster"
			r.References["connections.datacenter_id"] = config.Reference{
				TerraformName: "ionoscloud_datacenter",
			}
			r.References["connections.lan_id"] = config.Reference{
				TerraformName: "ionoscloud_lan",
			}
			r.References["location"] = config.Reference{
				TerraformName: "ionoscloud_datacenter",
				Extractor:     "github.com/ionos-cloud/provider-upjet-ionoscloud/config/common.DatacenterLocation()",
			}
			r.UseAsync = true
		},
	)
	p.AddResourceConfigurator(
		"ionoscloud_pg_user", func(r *config.Resource) {
			r.ShortGroup = "postgresql"
			r.Kind = "PostgresqlUser"
			r.References["cluster_id"] = config.Reference{
				TerraformName: "ionoscloud_pg_cluster",
			}
		},
	)
	p.AddResourceConfigurator(
		"ionoscloud_pg_database", func(r *config.Resource) {
			r.ShortGroup = "postgresql"
			r.Kind = "PostgresqlDatabase"
			r.References["cluster_id"] = config.Reference{
				TerraformName: "ionoscloud_pg_cluster",
			}
		},
	)
}
