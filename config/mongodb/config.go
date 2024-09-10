package mongodb

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator(
		"ionoscloud_mongo_cluster", func(r *config.Resource) {
			r.ShortGroup = "mongodb"
			r.Kind = "MongodbCluster"
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
			// This is added so that ram, storage_size and cores are not set by the late initializer in the ForProvider spec field
			r.LateInitializer = config.LateInitializer{
				IgnoredFields: []string{"template_id", "ram", "storage_size", "cores"},
			}
		},
	)

	p.AddResourceConfigurator(
		"ionoscloud_mongo_user", func(r *config.Resource) {
			r.ShortGroup = "mongodb"
			r.Kind = "MongodbUser"
			r.References["cluster_id"] = config.Reference{
				TerraformName: "ionoscloud_mongo_cluster",
			}
		},
	)
}
