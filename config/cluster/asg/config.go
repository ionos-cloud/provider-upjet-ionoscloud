package asg

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_autoscaling_group", func(r *config.Resource) {
		r.ShortGroup = "asg"
		r.Kind = "AutoscalingGroup"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["replica_configuration.nic.target_group.target_group_id"] = config.Reference{
			TerraformName: "ionoscloud_target_group",
		}

		r.References["replica_configuration.nic.lan"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
	})
}
