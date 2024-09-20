package certificatemanager

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_certificate", func(r *config.Resource) {
		r.Kind = "Certificate"
		r.ShortGroup = "certificatemanager"
	})

	p.AddResourceConfigurator("ionoscloud_auto_certificate_provider", func(r *config.Resource) {
		r.Kind = "AutoCertificateProvider"
		r.ShortGroup = "certificatemanager"
	})

	p.AddResourceConfigurator("ionoscloud_auto_certificate", func(r *config.Resource) {
		r.ShortGroup = "certificatemanager"
		r.Kind = "AutoCertificate"
		r.References["provider_id"] = config.Reference{
			TerraformName: "ionoscloud_auto_certificate_provider",
		}

		r.References["location"] = config.Reference{
			TerraformName: "ionoscloud_auto_certificate_provider",
			Extractor:     "github.com/ionos-cloud/provider-upjet-ionoscloud/config/common.AutoCertificateProviderLocation()",
		}
	})
}
