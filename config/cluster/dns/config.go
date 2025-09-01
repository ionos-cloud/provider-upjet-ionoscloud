package dns

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator(
		"ionoscloud_dns_zone", func(r *config.Resource) {
			r.ShortGroup = "dns"
			r.Kind = "DnsZone"
		},
	)
	p.AddResourceConfigurator(
		"ionoscloud_dns_record", func(r *config.Resource) {
			r.ShortGroup = "dns"
			r.Kind = "DnsRecord"
			r.References["zone_id"] = config.Reference{
				TerraformName: "ionoscloud_dns_zone",
			}
		},
	)
}
