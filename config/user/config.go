package user

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_user", func(r *config.Resource) {
		r.ShortGroup = "compute"
	})
	p.AddResourceConfigurator("ionoscloud_group", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["user_ids"] = config.Reference{
			TerraformName: "ionoscloud_user",
		}
	})
}
