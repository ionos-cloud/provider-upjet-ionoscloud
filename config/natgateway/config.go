package natgateway

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_natgateway", func(r *config.Resource) {
		r.ShortGroup = "natgateway"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
	})

	p.AddResourceConfigurator("ionoscloud_natgateway_rule", func(r *config.Resource) {
		r.ShortGroup = "natgateway"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["natgateway_id"] = config.Reference{
			TerraformName: "ionoscloud_natgateway",
		}
	})
}
