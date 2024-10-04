package k8s

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_k8s_cluster", func(r *config.Resource) {
		r.ShortGroup = "k8s"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["s3_buckets.name"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})
	p.AddResourceConfigurator("ionoscloud_k8s_node_pool", func(r *config.Resource) {
		r.ShortGroup = "k8s"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["k8s_cluster_id"] = config.Reference{
			TerraformName: "ionoscloud_k8s_cluster",
		}
		r.References["lans.id"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
	})
}
