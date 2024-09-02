package s3

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the s3bucket resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_s3_bucket", func(r *config.Resource) {
		r.ShortGroup = "s3"
	})
	p.AddResourceConfigurator("ionoscloud_s3_key", func(r *config.Resource) {
		r.ShortGroup = "s3"
		r.References["user_id"] = config.Reference{
			TerraformName: "ionoscloud_user",
		}
	})
}
