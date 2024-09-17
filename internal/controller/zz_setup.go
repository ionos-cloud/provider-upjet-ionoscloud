// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	loadbalancer "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/alb/loadbalancer"
	loadbalancerforwardingrule "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/alb/loadbalancerforwardingrule"
	apigateway "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/apigateway/apigateway"
	route "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/apigateway/route"
	autoscalinggroup "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/asg/autoscalinggroup"
	targetgroup "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/asg/targetgroup"
	unit "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/backupunit/unit"
	autocertificate "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/certificatemanager/autocertificate"
	autocertificateprovider "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/certificatemanager/autocertificateprovider"
	certificate "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/certificatemanager/certificate"
	bootdeviceselection "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/bootdeviceselection"
	crossconnect "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/crossconnect"
	datacenter "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/datacenter"
	firewall "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/firewall"
	group "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/group"
	ipblock "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/ipblock"
	ipfailover "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/ipfailover"
	lan "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/lan"
	loadbalancercompute "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/loadbalancer"
	nic "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/nic"
	server "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/server"
	share "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/share"
	snapshot "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/snapshot"
	user "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/user"
	volume "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/compute/volume"
	registry "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/containerregistry/registry"
	registrytoken "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/containerregistry/registrytoken"
	cluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/dataplatform/cluster"
	nodepool "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/dataplatform/nodepool"
	mariadbcluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/dbaas/mariadbcluster"
	dnsrecord "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/dns/dnsrecord"
	dnszone "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/dns/dnszone"
	inmemorydbreplicaset "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/inmemorydb/inmemorydbreplicaset"
	clusterk8s "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/k8s/cluster"
	nodepoolk8s "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/k8s/nodepool"
	kafka "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/kafka/kafka"
	kafkatopic "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/kafka/kafkatopic"
	pipeline "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/log/pipeline"
	mongodbcluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/mongodb/mongodbcluster"
	mongodbuser "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/mongodb/mongodbuser"
	natgateway "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/natgateway/natgateway"
	rule "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/natgateway/rule"
	nfscluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/nfs/nfscluster"
	nfsshare "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/nfs/nfsshare"
	forwardingrule "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/nlb/forwardingrule"
	networkloadbalancer "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/nlb/networkloadbalancer"
	postgresqlcluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/postgresql/postgresqlcluster"
	postgresqldatabase "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/postgresql/postgresqldatabase"
	postgresqluser "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/postgresql/postgresqluser"
	providerconfig "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/providerconfig"
	bucket "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/s3/bucket"
	bucketcorsconfiguration "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/s3/bucketcorsconfiguration"
	bucketlifecycleconfiguration "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/s3/bucketlifecycleconfiguration"
	bucketobjectlockconfiguration "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/s3/bucketobjectlockconfiguration"
	bucketpolicy "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/s3/bucketpolicy"
	bucketpublicaccessblock "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/s3/bucketpublicaccessblock"
	bucketserversideencryptionconfiguration "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/s3/bucketserversideencryptionconfiguration"
	bucketversioning "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/s3/bucketversioning"
	bucketwebsiteconfiguration "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/s3/bucketwebsiteconfiguration"
	key "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/s3/key"
	object "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/s3/object"
	objectcopy "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/s3/objectcopy"
	vpnipsecgateway "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/vpnipsec/vpnipsecgateway"
	vpnipsectunnel "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/vpnipsec/vpnipsectunnel"
	vpnwireguardgateway "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/vpnwireguard/vpnwireguardgateway"
	vpnwireguardpeer "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/vpnwireguard/vpnwireguardpeer"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		loadbalancer.Setup,
		loadbalancerforwardingrule.Setup,
		apigateway.Setup,
		route.Setup,
		autoscalinggroup.Setup,
		targetgroup.Setup,
		unit.Setup,
		autocertificate.Setup,
		autocertificateprovider.Setup,
		certificate.Setup,
		bootdeviceselection.Setup,
		crossconnect.Setup,
		datacenter.Setup,
		firewall.Setup,
		group.Setup,
		ipblock.Setup,
		ipfailover.Setup,
		lan.Setup,
		loadbalancercompute.Setup,
		nic.Setup,
		server.Setup,
		share.Setup,
		snapshot.Setup,
		user.Setup,
		volume.Setup,
		registry.Setup,
		registrytoken.Setup,
		cluster.Setup,
		nodepool.Setup,
		mariadbcluster.Setup,
		dnsrecord.Setup,
		dnszone.Setup,
		inmemorydbreplicaset.Setup,
		clusterk8s.Setup,
		nodepoolk8s.Setup,
		kafka.Setup,
		kafkatopic.Setup,
		pipeline.Setup,
		mongodbcluster.Setup,
		mongodbuser.Setup,
		natgateway.Setup,
		rule.Setup,
		nfscluster.Setup,
		nfsshare.Setup,
		forwardingrule.Setup,
		networkloadbalancer.Setup,
		postgresqlcluster.Setup,
		postgresqldatabase.Setup,
		postgresqluser.Setup,
		providerconfig.Setup,
		bucket.Setup,
		bucketcorsconfiguration.Setup,
		bucketlifecycleconfiguration.Setup,
		bucketobjectlockconfiguration.Setup,
		bucketpolicy.Setup,
		bucketpublicaccessblock.Setup,
		bucketserversideencryptionconfiguration.Setup,
		bucketversioning.Setup,
		bucketwebsiteconfiguration.Setup,
		key.Setup,
		object.Setup,
		objectcopy.Setup,
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
