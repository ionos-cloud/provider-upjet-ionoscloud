package cdn

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_cdn_distribution", func(r *config.Resource) {
		r.ShortGroup = "cdn"
		r.References["certificate_id"] = config.Reference{
			TerraformName: "ionoscloud_certificate",
		}
	})
}
