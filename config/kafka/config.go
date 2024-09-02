package kafka

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_kafka_cluster", func(r *config.Resource) {
		r.ShortGroup = "kafka"
		r.Kind = "Kafka"
		r.References["connections.datacenter_id"] = config.Reference{
			TerraformName: "ionoscloud_datacenter",
		}
		r.References["connections.lan_id"] = config.Reference{
			TerraformName: "ionoscloud_lan",
		}
	})
	p.AddResourceConfigurator("ionoscloud_kafka_cluster_topic", func(r *config.Resource) {
		r.ShortGroup = "kafka"
		r.Kind = "KafkaTopic"
		r.References["cluster_id"] = config.Reference{
			TerraformName: "ionoscloud_kafka_cluster",
		}
	})
}
