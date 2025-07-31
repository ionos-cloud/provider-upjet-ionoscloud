package alb

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_application_loadbalancer", func(r *config.Resource) {
		r.ShortGroup = "alb"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["listener_lan"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
		r.References["target_lan"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("ionoscloud_application_loadbalancer_forwardingrule", func(r *config.Resource) {
		r.ShortGroup = "alb"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["application_loadbalancer_id"] = config.Reference{
			TerraformName: "ionoscloud_application_loadbalancer",
		}
	})

	p.AddResourceConfigurator("ionoscloud_target_group", func(r *config.Resource) {
		r.Kind = "TargetGroup"
		r.ShortGroup = "alb"
		r.TerraformCustomDiff = common.IgnoreEmptyDiffForComputed([]string{"http_health_check.#", "health_check.#", "target.#"})
	})
}
