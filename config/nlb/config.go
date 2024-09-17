package nlb

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_networkloadbalancer", func(r *config.Resource) {
		r.ShortGroup = "nlb"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["listener_lan"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
		r.References["target_lan"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
	})

	p.AddResourceConfigurator("ionoscloud_networkloadbalancer_forwardingrule", func(r *config.Resource) {
		r.ShortGroup = "nlb"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["networkloadbalancer_id"] = config.Reference{
			TerraformName: "ionoscloud_networkloadbalancer",
		}
	})
}
