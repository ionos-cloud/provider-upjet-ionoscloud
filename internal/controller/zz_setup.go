// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	apigateway "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/apigateway/apigateway"
	route "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/apigateway/route"
	crossconnect "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/crossconnect"
	datacenter "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/datacenter"
	group "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/group"
	ipblock "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/ipblock"
	lan "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/lan"
	nic "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/nic"
	server "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/server"
	user "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/user"
	volume "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/volume"
	mariadbcluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/dbaas/mariadbcluster"
	dnsrecord "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/dns/dnsrecord"
	dnszone "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/dns/dnszone"
	inmemorydbreplicaset "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/inmemorydb/inmemorydbreplicaset"
	cluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/k8s/cluster"
	nodepool "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/k8s/nodepool"
	kafka "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/kafka/kafka"
	kafkatopic "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/kafka/kafkatopic"
	mongodbcluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/mongodb/mongodbcluster"
	mongodbuser "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/mongodb/mongodbuser"
	nfscluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/nfs/nfscluster"
	nfsshare "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/nfs/nfsshare"
	postgresqlcluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/postgresql/postgresqlcluster"
	postgresqldatabase "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/postgresql/postgresqldatabase"
	postgresqluser "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/postgresql/postgresqluser"
	providerconfig "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/providerconfig"
	bucket "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/s3/bucket"
	key "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/s3/key"
	vpnipsecgateway "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/vpnipsec/vpnipsecgateway"
	vpnipsectunnel "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/vpnipsec/vpnipsectunnel"
	vpnwireguardgateway "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/vpnwireguard/vpnwireguardgateway"
	vpnwireguardpeer "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/vpnwireguard/vpnwireguardpeer"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		apigateway.Setup,
		route.Setup,
		crossconnect.Setup,
		datacenter.Setup,
		group.Setup,
		ipblock.Setup,
		lan.Setup,
		nic.Setup,
		server.Setup,
		user.Setup,
		volume.Setup,
		mariadbcluster.Setup,
		dnsrecord.Setup,
		dnszone.Setup,
		inmemorydbreplicaset.Setup,
		cluster.Setup,
		nodepool.Setup,
		kafka.Setup,
		kafkatopic.Setup,
		mongodbcluster.Setup,
		mongodbuser.Setup,
		nfscluster.Setup,
		nfsshare.Setup,
		postgresqlcluster.Setup,
		postgresqldatabase.Setup,
		postgresqluser.Setup,
		providerconfig.Setup,
		bucket.Setup,
		key.Setup,
		vpnipsecgateway.Setup,
		vpnipsectunnel.Setup,
		vpnwireguardgateway.Setup,
		vpnwireguardpeer.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
