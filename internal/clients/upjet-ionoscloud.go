/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/ionos-cloud/sdk-go-bundle/shared"
	"github.com/ionos-cloud/terraform-provider-ionoscloud/v6/services/bundleclient"
	"github.com/ionos-cloud/terraform-provider-ionoscloud/v6/services/clientoptions"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/v2/pkg/terraform"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal upjet-ionoscloud credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string, fwProvider provider.Provider) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		providerConfig, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return terraform.Setup{}, err
		}
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
			FrameworkProvider: fwProvider,
		}

		data, err := resource.CommonCredentialExtractor(ctx, providerConfig.Spec.Credentials.Source, client, providerConfig.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}
		boolInsecure := false
		if insecure, exists := creds["insecure"]; exists {
			if boolInsecure, err = strconv.ParseBool(insecure); err != nil {
				return ps, errors.Wrapf(err, "cannot parse insecure value %q as boolean", insecure)
			}
		}

		ionosSDKBundleClient := bundleclient.New(clientoptions.TerraformClientOptions{
			ClientOptions: shared.ClientOptions{
				Endpoint:      creds["endpoint"],
				SkipTLSVerify: boolInsecure,
				Credentials: shared.Credentials{
					Username:    creds["user"],
					Password:    creds["password"],
					Token:       creds["token"],
					S3AccessKey: creds["s3_access_key"],
					S3SecretKey: creds["s3_secret_key"],
				},
			},
			StorageOptions: clientoptions.StorageOptions{
				AccessKey: creds["s3_access_key"],
				SecretKey: creds["s3_secret_key"],
				Region:    creds["s3_region"],
			},
			TerraformVersion: version,
			Version:          providerVersion,
		}, nil)

		ps.Meta = *ionosSDKBundleClient

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]any{
			"username":        creds["user"],
			"password":        creds["password"],
			"token":           creds["token"],
			"s3_access_key":   creds["s3_access_key"],
			"s3_secret_key":   creds["s3_secret_key"],
			"s3_region":       creds["s3_region"],
			"insecure":        boolInsecure,
			"endpoint":        creds["endpoint"],
			"contract_number": creds["contract_number"],
		}

		return ps, nil
	}
}
