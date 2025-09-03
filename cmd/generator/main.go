/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/crossplane/upjet/v2/pkg/pipeline"
	ionosterraform "github.com/ionos-cloud/terraform-provider-ionoscloud/v6/xpprovider"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/ionos-cloud/provider-upjet-ionoscloud/config"
)

func main() {
	var (
		app      = kingpin.New("generator", "Run Upjet code generation pipelines for provider-ionos").DefaultEnvars()
		repoRoot = app.Arg("repo-root", "Root directory for the provider repository").Required().String()
	)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	absRootDir, err := filepath.Abs(*repoRoot)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", *repoRoot))
	}

	fwProvider, sdkProvider := ionosterraform.GetProvider()
	providerCluster, err := config.GetProvider(fwProvider, sdkProvider, true)
	kingpin.FatalIfError(err, "Cannot get terraform provider configuration")
	pns, err := config.GetProviderNamespaced(fwProvider, sdkProvider, true)
	kingpin.FatalIfError(err, "Cannot initialize the provider namespaced configuration")

	pipeline.Run(providerCluster, pns, absRootDir)
}
