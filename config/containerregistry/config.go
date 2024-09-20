package containerregistry

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_container_registry", func(r *config.Resource) {
		r.ShortGroup = "containerregistry"
	})

	p.AddResourceConfigurator("ionoscloud_container_registry_token", func(r *config.Resource) {
		r.ShortGroup = "containerregistry"
		r.References["registry_id"] = config.Reference{
			TerraformName: "ionoscloud_container_registry",
		}
	})
}
