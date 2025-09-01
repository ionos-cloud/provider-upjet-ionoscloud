// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	loadbalancer "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/alb/loadbalancer"
	loadbalancerforwardingrule "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/alb/loadbalancerforwardingrule"
	targetgroup "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/alb/targetgroup"
	apigateway "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/apigateway/apigateway"
	route "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/apigateway/route"
	autoscalinggroup "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/asg/autoscalinggroup"
	distribution "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/cdn/distribution"
	autocertificate "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/certificatemanager/autocertificate"
	autocertificateprovider "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/certificatemanager/autocertificateprovider"
	certificate "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/certificatemanager/certificate"
	bootdeviceselection "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/bootdeviceselection"
	crossconnect "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/crossconnect"
	cubeserver "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/cubeserver"
	datacenter "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/datacenter"
	firewall "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/firewall"
	group "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/group"
	ipblock "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/ipblock"
	ipfailover "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/ipfailover"
	lan "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/lan"
	loadbalancercompute "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/loadbalancer"
	nic "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/nic"
	server "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/server"
	share "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/share"
	snapshot "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/snapshot"
	unit "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/unit"
	user "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/user"
	vcpuserver "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/vcpuserver"
	volume "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/compute/volume"
	registry "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/containerregistry/registry"
	registrytoken "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/containerregistry/registrytoken"
	cluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/dataplatform/cluster"
	nodepool "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/dataplatform/nodepool"
	mariadbcluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/dbaas/mariadbcluster"
	dnsrecord "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/dns/dnsrecord"
	dnszone "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/dns/dnszone"
	inmemorydbreplicaset "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/inmemorydb/inmemorydbreplicaset"
	clusterk8s "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/k8s/cluster"
	nodepoolk8s "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/k8s/nodepool"
	kafka "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/kafka/kafka"
	kafkatopic "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/kafka/kafkatopic"
	pipeline "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/log/pipeline"
	mongodbcluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/mongodb/mongodbcluster"
	mongodbuser "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/mongodb/mongodbuser"
	natgateway "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/natgateway/natgateway"
	rule "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/natgateway/rule"
	nfscluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/nfs/nfscluster"
	nfsshare "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/nfs/nfsshare"
	forwardingrule "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/nlb/forwardingrule"
	networkloadbalancer "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/nlb/networkloadbalancer"
	bucket "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/objectstorage/bucket"
	bucketcorsconfiguration "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/objectstorage/bucketcorsconfiguration"
	bucketlifecycleconfiguration "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/objectstorage/bucketlifecycleconfiguration"
	bucketobjectlockconfiguration "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/objectstorage/bucketobjectlockconfiguration"
	bucketpolicy "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/objectstorage/bucketpolicy"
	bucketpublicaccessblock "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/objectstorage/bucketpublicaccessblock"
	bucketserversideencryptionconfiguration "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/objectstorage/bucketserversideencryptionconfiguration"
	bucketversioning "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/objectstorage/bucketversioning"
	bucketwebsiteconfiguration "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/objectstorage/bucketwebsiteconfiguration"
	key "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/objectstorage/key"
	object "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/objectstorage/object"
	objectcopy "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/objectstorage/objectcopy"
	storageaccesskey "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/objectstorage/storageaccesskey"
	postgresqlcluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/postgresql/postgresqlcluster"
	postgresqldatabase "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/postgresql/postgresqldatabase"
	postgresqluser "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/postgresql/postgresqluser"
	providerconfig "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/providerconfig"
	vpnipsecgateway "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/vpnipsec/vpnipsecgateway"
	vpnipsectunnel "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/vpnipsec/vpnipsectunnel"
	vpnwireguardgateway "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/vpnwireguard/vpnwireguardgateway"
	vpnwireguardpeer "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced/vpnwireguard/vpnwireguardpeer"
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
		providerconfig.Setup,
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
		providerconfig.SetupGated,
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
