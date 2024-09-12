// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1"
	common "github.com/ionos-cloud/provider-upjet-ionoscloud/config/common"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this MongodbCluster.
func (mg *MongodbCluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	if mg.Spec.ForProvider.Connections != nil {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Connections.DatacenterID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Connections.DatacenterIDRef,
			Selector:     mg.Spec.ForProvider.Connections.DatacenterIDSelector,
			To: reference.To{
				List:    &v1alpha1.DatacenterList{},
				Managed: &v1alpha1.Datacenter{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Connections.DatacenterID")
		}
		mg.Spec.ForProvider.Connections.DatacenterID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Connections.DatacenterIDRef = rsp.ResolvedReference

	}
	if mg.Spec.ForProvider.Connections != nil {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Connections.LanID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Connections.LanIDRef,
			Selector:     mg.Spec.ForProvider.Connections.LanIDSelector,
			To: reference.To{
				List:    &v1alpha1.LanList{},
				Managed: &v1alpha1.Lan{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Connections.LanID")
		}
		mg.Spec.ForProvider.Connections.LanID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Connections.LanIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Location),
		Extract:      common.DatacenterLocation(),
		Reference:    mg.Spec.ForProvider.LocationRef,
		Selector:     mg.Spec.ForProvider.LocationSelector,
		To: reference.To{
			List:    &v1alpha1.DatacenterList{},
			Managed: &v1alpha1.Datacenter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Location")
	}
	mg.Spec.ForProvider.Location = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LocationRef = rsp.ResolvedReference

	if mg.Spec.InitProvider.Connections != nil {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Connections.DatacenterID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.Connections.DatacenterIDRef,
			Selector:     mg.Spec.InitProvider.Connections.DatacenterIDSelector,
			To: reference.To{
				List:    &v1alpha1.DatacenterList{},
				Managed: &v1alpha1.Datacenter{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Connections.DatacenterID")
		}
		mg.Spec.InitProvider.Connections.DatacenterID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Connections.DatacenterIDRef = rsp.ResolvedReference

	}
	if mg.Spec.InitProvider.Connections != nil {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Connections.LanID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.Connections.LanIDRef,
			Selector:     mg.Spec.InitProvider.Connections.LanIDSelector,
			To: reference.To{
				List:    &v1alpha1.LanList{},
				Managed: &v1alpha1.Lan{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Connections.LanID")
		}
		mg.Spec.InitProvider.Connections.LanID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Connections.LanIDRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Location),
		Extract:      common.DatacenterLocation(),
		Reference:    mg.Spec.InitProvider.LocationRef,
		Selector:     mg.Spec.InitProvider.LocationSelector,
		To: reference.To{
			List:    &v1alpha1.DatacenterList{},
			Managed: &v1alpha1.Datacenter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Location")
	}
	mg.Spec.InitProvider.Location = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LocationRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MongodbUser.
func (mg *MongodbUser) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &MongodbClusterList{},
			Managed: &MongodbCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ClusterIDRef,
		Selector:     mg.Spec.InitProvider.ClusterIDSelector,
		To: reference.To{
			List:    &MongodbClusterList{},
			Managed: &MongodbCluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ClusterID")
	}
	mg.Spec.InitProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ClusterIDRef = rsp.ResolvedReference

	return nil
}
