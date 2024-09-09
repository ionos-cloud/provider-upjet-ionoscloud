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
	inmemorydbreplicaset "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/inmemorydb/inmemorydbreplicaset"
	cluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/k8s/cluster"
	nodepool "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/k8s/nodepool"
	kafka "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/kafka/kafka"
	kafkatopic "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/kafka/kafkatopic"
	nfscluster "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/nfs/nfscluster"
	nfsshare "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/nfs/nfsshare"
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
		inmemorydbreplicaset.Setup,
		cluster.Setup,
		nodepool.Setup,
		kafka.Setup,
		kafkatopic.Setup,
		nfscluster.Setup,
		nfsshare.Setup,
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
