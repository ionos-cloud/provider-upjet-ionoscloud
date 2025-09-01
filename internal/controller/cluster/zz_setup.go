// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	loadbalancer "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/alb/loadbalancer"
	loadbalancerforwardingrule "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/alb/loadbalancerforwardingrule"
	targetgroup "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/alb/targetgroup"
	apigateway "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/apigateway/apigateway"
	route "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/apigateway/route"
	autoscalinggroup "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/asg/autoscalinggroup"
	distribution "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/cdn/distribution"
	autocertificate "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/certificatemanager/autocertificate"
	autocertificateprovider "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/certificatemanager/autocertificateprovider"
	certificate "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/certificatemanager/certificate"
	bootdeviceselection "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/bootdeviceselection"
	crossconnect "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/crossconnect"
	cubeserver "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/cubeserver"
	datacenter "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/datacenter"
	firewall "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/firewall"
	group "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/group"
	ipblock "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/ipblock"
	ipfailover "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/ipfailover"
	lan "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/lan"
	loadbalancercompute "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/loadbalancer"
	nic "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/nic"
	server "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/server"
	share "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/share"
	snapshot "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/snapshot"
	unit "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/unit"
	user "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/user"
	vcpuserver "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/vcpuserver"
	volume "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/compute/volume"
	registry "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/containerregistry/registry"
	registrytoken "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/containerregistry/registrytoken"
	cluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/dataplatform/cluster"
	nodepool "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/dataplatform/nodepool"
	mariadbcluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/dbaas/mariadbcluster"
	dnsrecord "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/dns/dnsrecord"
	dnszone "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/dns/dnszone"
	inmemorydbreplicaset "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/inmemorydb/inmemorydbreplicaset"
	clusterk8s "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/k8s/cluster"
	nodepoolk8s "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/k8s/nodepool"
	kafka "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/kafka/kafka"
	kafkatopic "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/kafka/kafkatopic"
	pipeline "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/log/pipeline"
	mongodbcluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/mongodb/mongodbcluster"
	mongodbuser "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/mongodb/mongodbuser"
	natgateway "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/natgateway/natgateway"
	rule "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/natgateway/rule"
	nfscluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/nfs/nfscluster"
	nfsshare "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/nfs/nfsshare"
	forwardingrule "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/nlb/forwardingrule"
	networkloadbalancer "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/nlb/networkloadbalancer"
	bucket "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/objectstorage/bucket"
	bucketcorsconfiguration "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/objectstorage/bucketcorsconfiguration"
	bucketlifecycleconfiguration "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/objectstorage/bucketlifecycleconfiguration"
	bucketobjectlockconfiguration "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/objectstorage/bucketobjectlockconfiguration"
	bucketpolicy "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/objectstorage/bucketpolicy"
	bucketpublicaccessblock "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/objectstorage/bucketpublicaccessblock"
	bucketserversideencryptionconfiguration "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/objectstorage/bucketserversideencryptionconfiguration"
	bucketversioning "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/objectstorage/bucketversioning"
	bucketwebsiteconfiguration "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/objectstorage/bucketwebsiteconfiguration"
	key "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/objectstorage/key"
	object "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/objectstorage/object"
	objectcopy "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/objectstorage/objectcopy"
	storageaccesskey "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/objectstorage/storageaccesskey"
	postgresqlcluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/postgresql/postgresqlcluster"
	postgresqldatabase "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/postgresql/postgresqldatabase"
	postgresqluser "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/postgresql/postgresqluser"
	vpnipsecgateway "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/vpnipsec/vpnipsecgateway"
	vpnipsectunnel "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/vpnipsec/vpnipsectunnel"
	vpnwireguardgateway "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/vpnwireguard/vpnwireguardgateway"
	vpnwireguardpeer "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster/vpnwireguard/vpnwireguardpeer"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		loadbalancer.Setup,
		loadbalancerforwardingrule.Setup,
		targetgroup.Setup,
		apigateway.Setup,
		route.Setup,
		autoscalinggroup.Setup,
		distribution.Setup,
		autocertificate.Setup,
		autocertificateprovider.Setup,
		certificate.Setup,
		bootdeviceselection.Setup,
		crossconnect.Setup,
		cubeserver.Setup,
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
		unit.Setup,
		user.Setup,
		vcpuserver.Setup,
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
		storageaccesskey.Setup,
		postgresqlcluster.Setup,
		postgresqldatabase.Setup,
		postgresqluser.Setup,
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

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		loadbalancer.SetupGated,
		loadbalancerforwardingrule.SetupGated,
		targetgroup.SetupGated,
		apigateway.SetupGated,
		route.SetupGated,
		autoscalinggroup.SetupGated,
		distribution.SetupGated,
		autocertificate.SetupGated,
		autocertificateprovider.SetupGated,
		certificate.SetupGated,
		bootdeviceselection.SetupGated,
		crossconnect.SetupGated,
		cubeserver.SetupGated,
		datacenter.SetupGated,
		firewall.SetupGated,
		group.SetupGated,
		ipblock.SetupGated,
		ipfailover.SetupGated,
		lan.SetupGated,
		loadbalancercompute.SetupGated,
		nic.SetupGated,
		server.SetupGated,
		share.SetupGated,
		snapshot.SetupGated,
		unit.SetupGated,
		user.SetupGated,
		vcpuserver.SetupGated,
		volume.SetupGated,
		registry.SetupGated,
		registrytoken.SetupGated,
		cluster.SetupGated,
		nodepool.SetupGated,
		mariadbcluster.SetupGated,
		dnsrecord.SetupGated,
		dnszone.SetupGated,
		inmemorydbreplicaset.SetupGated,
		clusterk8s.SetupGated,
		nodepoolk8s.SetupGated,
		kafka.SetupGated,
		kafkatopic.SetupGated,
		pipeline.SetupGated,
		mongodbcluster.SetupGated,
		mongodbuser.SetupGated,
		natgateway.SetupGated,
		rule.SetupGated,
		nfscluster.SetupGated,
		nfsshare.SetupGated,
		forwardingrule.SetupGated,
		networkloadbalancer.SetupGated,
		bucket.SetupGated,
		bucketcorsconfiguration.SetupGated,
		bucketlifecycleconfiguration.SetupGated,
		bucketobjectlockconfiguration.SetupGated,
		bucketpolicy.SetupGated,
		bucketpublicaccessblock.SetupGated,
		bucketserversideencryptionconfiguration.SetupGated,
		bucketversioning.SetupGated,
		bucketwebsiteconfiguration.SetupGated,
		key.SetupGated,
		object.SetupGated,
		objectcopy.SetupGated,
		storageaccesskey.SetupGated,
		postgresqlcluster.SetupGated,
		postgresqldatabase.SetupGated,
		postgresqluser.SetupGated,
		vpnipsecgateway.SetupGated,
		vpnipsectunnel.SetupGated,
		vpnwireguardgateway.SetupGated,
		vpnwireguardpeer.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
