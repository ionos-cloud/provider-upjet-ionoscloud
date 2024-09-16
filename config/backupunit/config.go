package backupunit

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_backup_unit", func(r *config.Resource) {
		r.ShortGroup = "backupunit"
	})
}
