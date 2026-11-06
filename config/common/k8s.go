package common

import (
	"context"
	"fmt"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ionos-cloud/terraform-provider-ionoscloud/v6/services/bundleclient"
)

const (
	// K8sKubeConfigAttr is the Terraform attribute added to the
	// ionoscloud_k8s_cluster resource that holds the cluster kubeconfig.
	K8sKubeConfigAttr = "kube_config"
	// K8sKubeConfigConnectionKey is the connection secret key the kubeconfig
	// is published under.
	K8sKubeConfigConnectionKey = "kubeconfig"
)

// ConfigureK8sClusterKubeConfig exposes the cluster kubeconfig as a
// connection detail of the ionoscloud_k8s_cluster resource.
func ConfigureK8sClusterKubeConfig(r *config.Resource) {
	if r.TerraformResource != nil {
		withKubeConfig(r.TerraformResource)
	}
	r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
		conn := map[string][]byte{}
		if kc, ok := attr[K8sKubeConfigAttr].(string); ok && kc != "" {
			conn[K8sKubeConfigConnectionKey] = []byte(kc)
		}
		return conn, nil
	}
}

// withKubeConfig adds a computed, sensitive kube_config attribute to the
// cluster resource (the Terraform resource does not expose one, only the data
// source does) and wraps Read so it is populated from the K8s kubeconfig API.
func withKubeConfig(r *schema.Resource) {
	// The *schema.Resource is shared between the cluster-scoped and the
	// namespaced provider configurations, so only patch it once.
	if _, done := r.Schema[K8sKubeConfigAttr]; done {
		return
	}
	r.Schema[K8sKubeConfigAttr] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Sensitive:   true,
		Description: "The kubeconfig of the cluster.",
	}
	read := r.ReadContext
	if read == nil {
		// The JSON-schema provider used for code generation has no Read.
		return
	}
	r.ReadContext = func(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
		diags := read(ctx, d, meta)
		if diags.HasError() || d.Id() == "" {
			return diags
		}
		bundle, ok := meta.(bundleclient.SdkBundle)
		if !ok {
			return diags
		}
		client, err := bundle.NewCloudAPIClient(ctx, d.Get("location").(string))
		if err != nil {
			return append(diags, diag.FromErr(err)...)
		}
		kc, _, err := client.KubernetesApi.K8sKubeconfigGet(ctx, d.Id()).Execute()
		if err != nil {
			// Do not fail the refresh: the kubeconfig may be unavailable while
			// the cluster is not ACTIVE (deploying, updating, destroying). The
			// previously known value is kept.
			return append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  fmt.Sprintf("cannot fetch kubeconfig for k8s cluster %s", d.Id()),
				Detail:   err.Error(),
			})
		}
		if err := d.Set(K8sKubeConfigAttr, kc); err != nil {
			return append(diags, diag.FromErr(err)...)
		}
		return diags
	}
}
