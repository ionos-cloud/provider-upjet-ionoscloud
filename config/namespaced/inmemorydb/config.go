package inmemorydb

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_inmemorydb_replicaset", func(r *config.Resource) {
		r.ShortGroup = "inmemorydb.replicaset"
		r.Kind = "InMemoryDBReplicaset"
		r.References["connections.datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["connections.lan_id"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("ionoscloud_inmemorydb_cluster_v2", func(r *config.Resource) {
		r.ShortGroup = "inmemorydbv2"
		r.Kind = "InMemoryDBCluster"
		r.References["connections.datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["connections.lan_id"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
		r.UseAsync = true
	})
}
