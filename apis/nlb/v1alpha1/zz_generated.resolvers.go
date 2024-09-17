// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/ionos-cloud/provider-upjet-ionoscloud/apis/compute/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Forwardingrule.
func (mg *Forwardingrule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatacenterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatacenterIDRef,
		Selector:     mg.Spec.ForProvider.DatacenterIDSelector,
		To: reference.To{
			List:    &v1alpha1.DatacenterList{},
			Managed: &v1alpha1.Datacenter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatacenterID")
	}
	mg.Spec.ForProvider.DatacenterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatacenterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkloadbalancerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkloadbalancerIDRef,
		Selector:     mg.Spec.ForProvider.NetworkloadbalancerIDSelector,
		To: reference.To{
			List:    &NetworkloadbalancerList{},
			Managed: &Networkloadbalancer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkloadbalancerID")
	}
	mg.Spec.ForProvider.NetworkloadbalancerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkloadbalancerIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DatacenterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.DatacenterIDRef,
		Selector:     mg.Spec.InitProvider.DatacenterIDSelector,
		To: reference.To{
			List:    &v1alpha1.DatacenterList{},
			Managed: &v1alpha1.Datacenter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DatacenterID")
	}
	mg.Spec.InitProvider.DatacenterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DatacenterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.NetworkloadbalancerID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.NetworkloadbalancerIDRef,
		Selector:     mg.Spec.InitProvider.NetworkloadbalancerIDSelector,
		To: reference.To{
			List:    &NetworkloadbalancerList{},
			Managed: &Networkloadbalancer{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.NetworkloadbalancerID")
	}
	mg.Spec.InitProvider.NetworkloadbalancerID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NetworkloadbalancerIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Networkloadbalancer.
func (mg *Networkloadbalancer) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DatacenterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.DatacenterIDRef,
		Selector:     mg.Spec.ForProvider.DatacenterIDSelector,
		To: reference.To{
			List:    &v1alpha1.DatacenterList{},
			Managed: &v1alpha1.Datacenter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.DatacenterID")
	}
	mg.Spec.ForProvider.DatacenterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.DatacenterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.ListenerLan),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ListenerLanRef,
		Selector:     mg.Spec.ForProvider.ListenerLanSelector,
		To: reference.To{
			List:    &v1alpha1.LanList{},
			Managed: &v1alpha1.Lan{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ListenerLan")
	}
	mg.Spec.ForProvider.ListenerLan = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ListenerLanRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.ForProvider.TargetLan),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TargetLanRef,
		Selector:     mg.Spec.ForProvider.TargetLanSelector,
		To: reference.To{
			List:    &v1alpha1.LanList{},
			Managed: &v1alpha1.Lan{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TargetLan")
	}
	mg.Spec.ForProvider.TargetLan = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetLanRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DatacenterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.DatacenterIDRef,
		Selector:     mg.Spec.InitProvider.DatacenterIDSelector,
		To: reference.To{
			List:    &v1alpha1.DatacenterList{},
			Managed: &v1alpha1.Datacenter{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.DatacenterID")
	}
	mg.Spec.InitProvider.DatacenterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.DatacenterIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.InitProvider.ListenerLan),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ListenerLanRef,
		Selector:     mg.Spec.InitProvider.ListenerLanSelector,
		To: reference.To{
			List:    &v1alpha1.LanList{},
			Managed: &v1alpha1.Lan{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ListenerLan")
	}
	mg.Spec.InitProvider.ListenerLan = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ListenerLanRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromFloatPtrValue(mg.Spec.InitProvider.TargetLan),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.TargetLanRef,
		Selector:     mg.Spec.InitProvider.TargetLanSelector,
		To: reference.To{
			List:    &v1alpha1.LanList{},
			Managed: &v1alpha1.Lan{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TargetLan")
	}
	mg.Spec.InitProvider.TargetLan = reference.ToFloatPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TargetLanRef = rsp.ResolvedReference

	return nil
}
