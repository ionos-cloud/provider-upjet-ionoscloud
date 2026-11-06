package common

import (
	"context"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func newK8sClusterResource(read schema.ReadContextFunc) *schema.Resource {
	return &schema.Resource{
		ReadContext: read,
		Schema: map[string]*schema.Schema{
			"name":     {Type: schema.TypeString, Required: true},
			"location": {Type: schema.TypeString, Optional: true},
		},
	}
}

func TestConfigureK8sClusterKubeConfigSchema(t *testing.T) {
	tr := newK8sClusterResource(nil)
	ConfigureK8sClusterKubeConfig(&config.Resource{TerraformResource: tr})

	s, ok := tr.Schema[K8sKubeConfigAttr]
	if !ok {
		t.Fatalf("expected %q attribute to be added", K8sKubeConfigAttr)
	}
	if !s.Computed || !s.Sensitive || s.Type != schema.TypeString {
		t.Fatalf("unexpected schema for %q: %+v", K8sKubeConfigAttr, s)
	}
	if tr.ReadContext != nil {
		t.Fatal("ReadContext must stay nil when the resource has no Read (code generation)")
	}
}

func TestConfigureK8sClusterKubeConfigWrapsReadOnce(t *testing.T) {
	calls := 0
	read := func(context.Context, *schema.ResourceData, any) diag.Diagnostics {
		calls++
		return nil
	}
	tr := newK8sClusterResource(read)

	// Cluster-scoped and namespaced configurations share the same *schema.Resource.
	ConfigureK8sClusterKubeConfig(&config.Resource{TerraformResource: tr})
	wrapped := tr.ReadContext
	sch := tr.Schema[K8sKubeConfigAttr]
	ConfigureK8sClusterKubeConfig(&config.Resource{TerraformResource: tr})

	if tr.Schema[K8sKubeConfigAttr] != sch {
		t.Fatal("schema attribute must not be replaced on the second configuration")
	}

	d := tr.TestResourceData()
	d.SetId("cluster-id")
	// meta is not an SdkBundle, so the wrapper must only delegate to the original Read.
	if diags := wrapped(context.Background(), d, nil); diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if diags := tr.ReadContext(context.Background(), d, nil); diags.HasError() {
		t.Fatalf("unexpected diagnostics: %v", diags)
	}
	if calls != 2 {
		t.Fatalf("expected the original Read to be called once per refresh, got %d calls for 2 refreshes", calls)
	}
}

func TestK8sClusterAdditionalConnectionDetails(t *testing.T) {
	r := &config.Resource{}
	ConfigureK8sClusterKubeConfig(r)

	cases := map[string]struct {
		attr map[string]any
		want map[string]string
	}{
		"WithKubeConfig": {
			attr: map[string]any{K8sKubeConfigAttr: "apiVersion: v1"},
			want: map[string]string{K8sKubeConfigConnectionKey: "apiVersion: v1"},
		},
		"EmptyKubeConfig": {
			attr: map[string]any{K8sKubeConfigAttr: ""},
			want: map[string]string{},
		},
		"MissingKubeConfig": {
			attr: map[string]any{},
			want: map[string]string{},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := r.Sensitive.AdditionalConnectionDetailsFn(tc.attr)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(got) != len(tc.want) {
				t.Fatalf("got %d keys, want %d: %v", len(got), len(tc.want), got)
			}
			for k, v := range tc.want {
				if string(got[k]) != v {
					t.Fatalf("key %q: got %q, want %q", k, got[k], v)
				}
			}
		})
	}
}
