package log

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_logging_pipeline", func(r *config.Resource) {
		r.ShortGroup = "log"
	})
}
