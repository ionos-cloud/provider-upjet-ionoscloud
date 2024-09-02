package wireguard

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_vpn_wireguard_gateway", func(r *config.Resource) {
		r.ShortGroup = "vpnwireguard"
		r.Kind = "VpnWireguardGateway"
		r.References["gateway_ip"] = config.Reference{
			TerraformName: "ionoscloud_ipblock",
			Extractor:     "github.com/ionos-cloud/provider-ionoscloud/config/common.FirstIPBlockIP()",
		}
		r.References["connections.datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["connections.lan_id"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
	})
	p.AddResourceConfigurator("ionoscloud_vpn_wireguard_peer", func(r *config.Resource) {
		r.ShortGroup = "vpnwireguard"
		r.Kind = "VpnWireguardPeer"
		r.References["gateway_id"] = config.Reference{
			TerraformName: "ionoscloud_vpn_wireguard_gateway",
		}
	})
}
