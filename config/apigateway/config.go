package apigateway

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_apigateway", func(r *config.Resource) {
		r.ShortGroup = "apigateway"

	})
}
