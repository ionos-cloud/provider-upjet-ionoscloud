package dataplatform

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_dataplatform_cluster", func(r *config.Resource) {
		r.ShortGroup = "dataplatform"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
	})

	p.AddResourceConfigurator("ionoscloud_dataplatform_node_pool", func(r *config.Resource) {
		r.ShortGroup = "dataplatform"
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ionoscloud_dataplatform_cluster",
		}
	})
}
