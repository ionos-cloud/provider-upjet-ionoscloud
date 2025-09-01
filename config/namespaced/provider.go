package namespaced

import (
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/alb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/apigateway"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/asg"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/cdn"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/certificatemanager"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/compute"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/containerregistry"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/dataplatform"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/dns"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/inmemorydb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/ipsec"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/k8s"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/kafka"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/log"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/mariadb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/mongodb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/natgateway"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/nfs"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/nlb"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/objectstorage"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/postgresql"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config/namespaced/wireguard"
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
