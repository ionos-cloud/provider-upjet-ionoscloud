/*
Copyright 2021 Upbound Inc.
*/

package providerconfig

import (
	"github.com/crossplane/crossplane-runtime/v2/pkg/event"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/controller"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/ionos-cloud/provider-upjet-ionoscloud/apis/v1beta1"
)

// Setup adds a controller that reconciles ProviderConfigs by accounting for
// their current usage.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	name := managed.ControllerName(v1beta1.ProviderConfigGroupKind)

	of := resource.ManagedKind{
		Group:   v1beta1.ProviderConfigGroupVersionKind.Group,
		Version: v1beta1.ProviderConfigGroupVersionKind.Version,
		Kind:    v1beta1.ProviderConfigGroupKind,
	}

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1beta1.ProviderConfig{}).
		Watches(&v1beta1.ProviderConfigUsage{}, &resource.EnqueueRequestForProviderConfig{}).
		Complete(managed.NewReconciler(mgr, of,
			managed.WithLogger(o.Logger.WithValues("controller", name)),
			managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name)))))
}
