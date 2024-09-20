package apigateway

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_apigateway", func(r *config.Resource) {
		r.ShortGroup = "apigateway"
	})

	p.AddResourceConfigurator("ionoscloud_apigateway_route", func(r *config.Resource) {
		r.ShortGroup = "apigateway"
		r.References["gateway_id"] = config.Reference{
			TerraformName: "ionoscloud_apigateway",
		}
	})
}
