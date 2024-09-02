package lan

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_lan", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["pcc"] = config.Reference{
			TerraformName: "ionoscloud_private_crossconnect",
		}
	})
}
