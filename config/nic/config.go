package nic

import "github.com/crossplane/upjet/pkg/config"

// Configure configures the nic resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_nic", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["server_id"] = config.Reference{
			TerraformName: "ionoscloud_server",
		}
		r.References["lan"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
	})
}
