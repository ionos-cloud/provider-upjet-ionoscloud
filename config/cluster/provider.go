package cluster

import (
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/alb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/apigateway"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/asg"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/cdn"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/certificatemanager"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/compute"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/containerregistry"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/dataplatform"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/dns"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/inmemorydb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/ipsec"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/k8s"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/kafka"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/log"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/mariadb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/mongodb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/natgateway"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/nfs"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/nlb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/objectstorage"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/postgresql"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/cluster/wireguard"
)

func init() {
	ProviderConfiguration.AddConfig(alb.Configure)
	ProviderConfiguration.AddConfig(apigateway.Configure)
	ProviderConfiguration.AddConfig(asg.Configure)
	ProviderConfiguration.AddConfig(cdn.Configure)
	ProviderConfiguration.AddConfig(certificatemanager.Configure)
	ProviderConfiguration.AddConfig(compute.Configure)
	ProviderConfiguration.AddConfig(containerregistry.Configure)
	ProviderConfiguration.AddConfig(dataplatform.Configure)
	ProviderConfiguration.AddConfig(dns.Configure)
	ProviderConfiguration.AddConfig(inmemorydb.Configure)
	ProviderConfiguration.AddConfig(ipsec.Configure)
	ProviderConfiguration.AddConfig(k8s.Configure)
	ProviderConfiguration.AddConfig(kafka.Configure)
	ProviderConfiguration.AddConfig(log.Configure)
	ProviderConfiguration.AddConfig(mariadb.Configure)
	ProviderConfiguration.AddConfig(mongodb.Configure)
	ProviderConfiguration.AddConfig(natgateway.Configure)
	ProviderConfiguration.AddConfig(nfs.Configure)
	ProviderConfiguration.AddConfig(nlb.Configure)
	ProviderConfiguration.AddConfig(objectstorage.Configure)
	ProviderConfiguration.AddConfig(postgresql.Configure)
	ProviderConfiguration.AddConfig(wireguard.Configure)
}
