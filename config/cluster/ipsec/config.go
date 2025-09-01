package ipsec

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator(
		"ionoscloud_vpn_ipsec_gateway", func(r *config.Resource) {
			r.ShortGroup = "vpnipsec"
			r.Kind = "VpnIpsecGateway"
			r.References["gateway_ip"] = config.Reference{
				TerraformName: "ionoscloud_ipblock",
				Extractor:     "github.com/ionos-cloud/provider-upjet-ionoscloud/config/common.FirstIPBlockIP()",
			}
			r.References["connections.datacenter_id"] = config.Reference{
				TerraformName: "ionoscloud_datacenter",
			}
			r.References["connections.lan_id"] = config.Reference{
				TerraformName: "ionoscloud_lan",
			}
			r.UseAsync = true
		},
	)

	p.AddResourceConfigurator(
		"ionoscloud_vpn_ipsec_tunnel", func(r *config.Resource) {
			r.ShortGroup = "vpnipsec"
			r.Kind = "VpnIpsecTunnel"
			r.References["gateway_id"] = config.Reference{
				TerraformName: "ionoscloud_vpn_ipsec_gateway",
			}
		},
	)
}
