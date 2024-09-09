/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/crossplane/upjet/pkg/pipeline"

	"github.com/ionos-cloud/provider-upjet-ionoscloud/config"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("root directory is required to be given as argument")
	}
	rootDir := os.Args[1]
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", rootDir))
	}

	ctx := context.Background()
	provider, err := config.GetProvider(ctx)
	if err != nil {
		panic(fmt.Sprintf("cannot get provider configuration: %v", err))
	}

	pipeline.Run(provider, absRootDir)
}
