// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package providerconfig

import (
	"github.com/crossplane/crossplane-runtime/v2/pkg/event"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/providerconfig"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/controller"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/ionos-cloud/provider-upjet-ionoscloud/apis/namespaced/v1beta1"
)

// Setup adds controllers that reconcile ProviderConfigs and
// ClusterProviderConfigs by accounting for their current usage.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	if err := setupNamespaced(mgr, o); err != nil {
		return err
	}
	return setupCluster(mgr, o)
}

// setupNamespaced adds a controller that reconciles namespaced ProviderConfigs
// by accounting for their current usage.
func setupNamespaced(mgr ctrl.Manager, o controller.Options) error {
	name := providerconfig.ControllerName(v1beta1.ProviderConfigGroupKind)

	of := resource.ProviderConfigKinds{
		Config:    v1beta1.ProviderConfigGroupVersionKind,
		Usage:     v1beta1.ProviderConfigUsageGroupVersionKind,
		UsageList: v1beta1.ProviderConfigUsageListGroupVersionKind,
	}

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1beta1.ProviderConfig{}).
		// Only enqueue for usages that reference a namespaced ProviderConfig.
		Watches(&v1beta1.ProviderConfigUsage{}, &resource.EnqueueRequestForProviderConfig{Kind: v1beta1.ProviderConfigKind}).
		Complete(providerconfig.NewReconciler(mgr, of,
			providerconfig.WithLogger(o.Logger.WithValues("controller", name)),
			providerconfig.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))) //nolint:staticcheck // suppress until crossplane-runtime offers the new recorder api
}

// setupCluster adds a controller that reconciles ClusterProviderConfigs by
// accounting for their current usage. Usages are recorded as namespaced
// ProviderConfigUsage objects whose providerConfigRef.kind is
// ClusterProviderConfig; the runtime reconciler matches them across all
// namespaces via the provider.crossplane.io/name and /kind labels.
func setupCluster(mgr ctrl.Manager, o controller.Options) error {
	name := providerconfig.ControllerName(v1beta1.ClusterProviderConfigGroupKind)

	of := resource.ProviderConfigKinds{
		Config:    v1beta1.ClusterProviderConfigGroupVersionKind,
		Usage:     v1beta1.ProviderConfigUsageGroupVersionKind,
		UsageList: v1beta1.ProviderConfigUsageListGroupVersionKind,
	}

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1beta1.ClusterProviderConfig{}).
		// Only enqueue for usages that reference a ClusterProviderConfig.
		Watches(&v1beta1.ProviderConfigUsage{}, &resource.EnqueueRequestForProviderConfig{Kind: v1beta1.ClusterProviderConfigKind}).
		Complete(providerconfig.NewReconciler(mgr, of,
			providerconfig.WithLogger(o.Logger.WithValues("controller", name)),
			providerconfig.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))))) //nolint:staticcheck // suppress until crossplane-runtime offers the new recorder api
}

// SetupGated adds controllers that reconcile ProviderConfigs and
// ClusterProviderConfigs by accounting for their current usage, once all the
// required CRDs are established.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	o.Gate.Register(func() {
		if err := Setup(mgr, o); err != nil {
			mgr.GetLogger().Error(err, "unable to setup reconcilers",
				"gvk", v1beta1.ProviderConfigGroupVersionKind.String(),
				"gvk", v1beta1.ClusterProviderConfigGroupVersionKind.String())
		}
	}, v1beta1.ProviderConfigGroupVersionKind, v1beta1.ClusterProviderConfigGroupVersionKind, v1beta1.ProviderConfigUsageGroupVersionKind)
	return nil
}
