package asg

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_autoscaling_group", func(r *config.Resource) {
		r.ShortGroup = "asg"
		r.Kind = "AutoscalingGroup"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["nic[0].target_group"] = config.Reference{
			TerraformName: "ionoscloud_target_group",
		}

		r.References["nic[*].lan"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
	})
}
