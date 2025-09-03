/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"io"
	"os"
	"path/filepath"
	"time"

	xpcontroller "github.com/crossplane/crossplane-runtime/v2/pkg/controller"
	"github.com/crossplane/crossplane-runtime/v2/pkg/feature"
	"github.com/crossplane/crossplane-runtime/v2/pkg/logging"
	"github.com/crossplane/crossplane-runtime/v2/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/v2/pkg/statemetrics"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	"github.com/ionos-cloud/terraform-provider-ionoscloud/v6/xpprovider"
	"gopkg.in/alecthomas/kingpin.v2"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/metrics"

	resolverapis "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/apis"

	clusterapis "github.com/ionos-cloud/provider-upjet-ionoscloud/apis/cluster"
	namespacedapis "github.com/ionos-cloud/provider-upjet-ionoscloud/apis/namespaced"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/config"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/internal/clients"
	clustercontroller "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/cluster"
	namespacedcontroller "github.com/ionos-cloud/provider-upjet-ionoscloud/internal/controller/namespaced"
	"github.com/ionos-cloud/provider-upjet-ionoscloud/internal/features"
)

func main() {
	var (
		app                     = kingpin.New(filepath.Base(os.Args[0]), "Terraform based Crossplane provider for Artifactory").DefaultEnvars()
		debug                   = app.Flag("debug", "Run with debug logging.").Short('d').Bool()
		syncPeriod              = app.Flag("sync", "Controller manager sync period such as 300ms, 1.5h, or 2h45m").Short('s').Default("1h").Duration()
		pollInterval            = app.Flag("poll", "Poll interval controls how often an individual resource should be checked for drift.").Default("10m").Duration()
		pollStateMetricInterval = app.Flag("poll-state-metric", "State metric recording interval").Default("5s").Duration()
		leaderElection          = app.Flag("leader-election", "Use leader election for the controller manager.").Short('l').Default("false").OverrideDefaultFromEnvar("LEADER_ELECTION").Bool()
		maxReconcileRate        = app.Flag("max-reconcile-rate", "The global maximum rate per second at which resources may be checked for drift from the desired state.").Default("10").Int()

		terraformVersion = app.Flag("terraform-version", "Terraform version.").Envar("TERRAFORM_VERSION").Default("1.5.7").String()
		providerSource   = app.Flag("terraform-provider-source", "Terraform provider source.").Envar("TERRAFORM_PROVIDER_SOURCE").Default("ionos-cloud/ionoscloud").String()
		providerVersion  = app.Flag("terraform-provider-version", "Terraform provider version.").Envar("TERRAFORM_PROVIDER_VERSION").Default("6.7.12").String()

		enableManagementPolicies = app.Flag("enable-management-policies", "Enable support for Management Policies.").Default("true").Envar("ENABLE_MANAGEMENT_POLICIES").Bool()
	)
	*debug = true
	kingpin.MustParse(app.Parse(os.Args[1:]))

	zl := zap.New(zap.UseDevMode(*debug))
	log := logging.NewLogrLogger(zl.WithName("provider-upjet-ionoscloud"))
	if *debug {
		// The controller-runtime runs with a no-op logger by default. It is
		// *very* verbose even at info level, so we only provide it a real
		// logger when we're running in debug mode.
		ctrl.SetLogger(zl)
	} else {
		// explicitly provide a no-op logger by default, otherwise controller-runtime gives a warning
		ctrl.SetLogger(zap.New(zap.WriteTo(io.Discard)))
	}

	log.Debug("Starting", "sync-period", syncPeriod.String(), "poll-interval", pollInterval.String(), "max-reconcile-rate", *maxReconcileRate)

	cfg, err := ctrl.GetConfig()
	kingpin.FatalIfError(err, "Cannot get API server rest config")

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		LeaderElection:   *leaderElection,
		LeaderElectionID: "crossplane-leader-election-provider-upjet-ionoscloud",
		Cache: cache.Options{
			SyncPeriod: syncPeriod,
		},
		LeaderElectionResourceLock: resourcelock.LeasesResourceLock,
		LeaseDuration:              func() *time.Duration { d := 60 * time.Second; return &d }(),
		RenewDeadline:              func() *time.Duration { d := 50 * time.Second; return &d }(),
	})
	kingpin.FatalIfError(err, "Cannot create controller manager")
	kingpin.FatalIfError(clusterapis.AddToScheme(mgr.GetScheme()), "Cannot add cluster scoped AWS APIs to scheme")
	kingpin.FatalIfError(resolverapis.BuildScheme(clusterapis.AddToSchemes), "Cannot register cluster scoped AWS APIs with the API resolver's runtime scheme")
	kingpin.FatalIfError(namespacedapis.AddToScheme(mgr.GetScheme()), "Cannot add namespaced AWS APIs to scheme")
	kingpin.FatalIfError(resolverapis.BuildScheme(namespacedapis.AddToSchemes), "Cannot register namespaced AWS APIs with the API resolver's runtime scheme")

	metricRecorder := managed.NewMRMetricRecorder()
	stateMetrics := statemetrics.NewMRStateMetrics()

	metrics.Registry.MustRegister(metricRecorder)
	metrics.Registry.MustRegister(stateMetrics)

	fwprovider, sdkprovider := xpprovider.GetProvider()
	clusterProvider, err := config.GetProvider(fwprovider, sdkprovider, false)
	kingpin.FatalIfError(err, "Cannot get cluster terraform provider configuration")
	//
	// Cluster scoped configuration
	//
	clusterOptions := tjcontroller.Options{
		Options: xpcontroller.Options{
			Logger:                  log,
			GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
			PollInterval:            *pollInterval,
			MaxConcurrentReconciles: *maxReconcileRate,
			Features:                &feature.Flags{},
			MetricOptions: &xpcontroller.MetricOptions{
				PollStateMetricInterval: *pollStateMetricInterval,
				MRMetrics:               metricRecorder,
				MRStateMetrics:          stateMetrics,
			},
		},
		OperationTrackerStore: tjcontroller.NewOperationStore(log),
		Provider:              clusterProvider,
		// use the following WorkspaceStoreOption to enable the shared gRPC mode
		// terraform.WithProviderRunner(terraform.NewSharedProvider(log, os.Getenv("TERRAFORM_NATIVE_PROVIDER_PATH"), terraform.WithNativeProviderArgs("-debuggable")))
		WorkspaceStore: terraform.NewWorkspaceStore(log),
		SetupFn:        clients.TerraformSetupBuilder(*terraformVersion, *providerSource, *providerVersion, clusterProvider.TerraformPluginFrameworkProvider),
	}

	namespacedProvider, err := config.GetProviderNamespaced(fwprovider, sdkprovider, false)
	kingpin.FatalIfError(err, "Cannot get namespaced terraform provider configuration")
	//
	// Namespaced configuration
	//
	namespacedOptions := tjcontroller.Options{
		Options: xpcontroller.Options{
			Logger:                  log,
			GlobalRateLimiter:       ratelimiter.NewGlobal(*maxReconcileRate),
			PollInterval:            *pollInterval,
			MaxConcurrentReconciles: *maxReconcileRate,
			Features:                &feature.Flags{},
			MetricOptions: &xpcontroller.MetricOptions{
				PollStateMetricInterval: *pollStateMetricInterval,
				MRMetrics:               metricRecorder,
				MRStateMetrics:          stateMetrics,
			},
		},
		OperationTrackerStore: tjcontroller.NewOperationStore(log),
		Provider:              namespacedProvider,
		// use the following WorkspaceStoreOption to enable the shared gRPC mode
		// terraform.WithProviderRunner(terraform.NewSharedProvider(log, os.Getenv("TERRAFORM_NATIVE_PROVIDER_PATH"), terraform.WithNativeProviderArgs("-debuggable")))
		WorkspaceStore: terraform.NewWorkspaceStore(log),
		SetupFn:        clients.TerraformSetupBuilder(*terraformVersion, *providerSource, *providerVersion, namespacedProvider.TerraformPluginFrameworkProvider),
	}

	if *enableManagementPolicies {
		clusterOptions.Features.Enable(features.EnableBetaManagementPolicies)
		namespacedOptions.Features.Enable(features.EnableBetaManagementPolicies)
		log.Info("Beta feature enabled", "flag", features.EnableBetaManagementPolicies)
	}

	kingpin.FatalIfError(clustercontroller.Setup(mgr, clusterOptions), "Cannot setup cluster scoped controllers")
	kingpin.FatalIfError(namespacedcontroller.Setup(mgr, namespacedOptions), "Cannot setup namespaced controllers")
	kingpin.FatalIfError(mgr.Start(ctrl.SetupSignalHandler()), "Cannot start controller manager")
}
