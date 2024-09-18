package compute

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_datacenter", func(r *config.Resource) {
		r.ShortGroup = "compute"
	})

	p.AddResourceConfigurator("ionoscloud_firewall", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["server_id"] = config.Reference{
			TerraformName: "ionoscloud_server",
		}

		r.References["nic_id"] = config.Reference{
			TerraformName: "ionoscloud_nic",
		}

		r.References["source_ip"] = config.Reference{
			TerraformName: "ionoscloud_ipblock",
		}

		r.References["target_ip"] = config.Reference{
			TerraformName: "ionoscloud_ipblock",
		}
	})

	p.AddResourceConfigurator("ionoscloud_lan", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["pcc"] = config.Reference{
			TerraformName: "ionoscloud_private_crossconnect",
		}
	})

	p.AddResourceConfigurator("ionoscloud_nic", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["server_id"] = config.Reference{
			TerraformName: "ionoscloud_server",
		}
		r.References["lan"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
	})

	p.AddResourceConfigurator("ionoscloud_server", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["nic.lan"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
	})

	p.AddResourceConfigurator("ionoscloud_volume", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["server_id"] = config.Reference{
			TerraformName: "ionoscloud_server",
		}
	})

	p.AddResourceConfigurator("ionoscloud_user", func(r *config.Resource) {
		r.ShortGroup = "compute"
	})
	p.AddResourceConfigurator("ionoscloud_group", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["user_ids"] = config.Reference{
			TerraformName: "ionoscloud_user",
		}
	})

	p.AddResourceConfigurator("ionoscloud_ipblock", func(r *config.Resource) {
		r.ShortGroup = "compute"
	})

	p.AddResourceConfigurator("ionoscloud_private_crossconnect", func(r *config.Resource) {
		r.ShortGroup = "compute"
	})

	p.AddResourceConfigurator("ionoscloud_snapshot", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["volume_id"] = config.Reference{
			TerraformName: "ionoscloud_volume",
		}
	})

	p.AddResourceConfigurator("ionoscloud_loadbalancer", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
	})

	p.AddResourceConfigurator("ionoscloud_ipfailover", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["nicuuid"] = config.Reference{
			TerraformName: "ionoscloud_server",
			Extractor:     "github.com/ionos-cloud/provider-upjet-ionoscloud/config/common.ServerPrimaryNIC()",
		}
		r.References["ip"] = config.Reference{
			TerraformName: "ionoscloud_ipblock",
			Extractor:     "github.com/ionos-cloud/provider-upjet-ionoscloud/config/common.FirstIPBlockIP()",
		}
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["lan"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
	})

	p.AddResourceConfigurator("ionoscloud_cube_server", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.Kind = "CubeServer"

		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["nic.lan"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
	})

	p.AddResourceConfigurator("ionoscloud_vcpu_server", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.Kind = "VCPUServer"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["nic.lan"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
	})

	p.AddResourceConfigurator("ionoscloud_server_boot_device_selection", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["server_id"] = config.Reference{
			TerraformName: "ionoscloud_server",
		}

		r.References["boot_device_id"] = config.Reference{
			TerraformName: "ionoscloud_volume",
		}
	})

	p.AddResourceConfigurator("ionoscloud_share", func(r *config.Resource) {
		r.ShortGroup = "compute"
		r.References["group_id"] = config.Reference{
			TerraformName: "ionoscloud_group",
		}

		r.References["resource_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
	})
}
