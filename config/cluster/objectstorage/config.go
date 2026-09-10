package objectstorage

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

const storage = "objectstorage"

// Configure configures the object storage resources
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_s3_bucket", func(r *config.Resource) {
		r.ShortGroup = storage
	})

	p.AddResourceConfigurator("ionoscloud_s3_bucket_policy", func(r *config.Resource) {
		r.ShortGroup = storage
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_bucket_versioning", func(r *config.Resource) {
		r.ShortGroup = storage
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_bucket_cors_configuration", func(r *config.Resource) {
		r.ShortGroup = storage
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_bucket_lifecycle_configuration", func(r *config.Resource) {
		r.ShortGroup = storage
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_bucket_public_access_block", func(r *config.Resource) {
		r.ShortGroup = storage
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_bucket_website_configuration", func(r *config.Resource) {
		r.ShortGroup = storage
	})

	p.AddResourceConfigurator("ionoscloud_s3_bucket_object_lock_configuration", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_bucket_server_side_encryption_configuration", func(r *config.Resource) {
		r.ShortGroup = storage
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_object", func(r *config.Resource) {
		r.ShortGroup = storage
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_object_copy", func(r *config.Resource) {
		r.ShortGroup = storage
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_key", func(r *config.Resource) {
		r.ShortGroup = storage
		r.References["user_id"] = config.Reference{
			TerraformName: "ionoscloud_user",
		}
		// The Key's own ID is its access key ID; publish it alongside the
		// automatically-published secret_key so both halves of the credential
		// pair are available from a single connection secret.
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if id, ok := attr["id"].(string); ok && id != "" {
				conn["accesskey"] = []byte(id)
			}
			return conn, nil
		}
	})
	p.AddResourceConfigurator("ionoscloud_object_storage_accesskey", func(r *config.Resource) {
		r.ShortGroup = storage
		// "accesskey" (the usable access key string) is a separate schema
		// attribute from "id" (the resource's internal UUID) on this
		// resource; publish it alongside the automatically-published
		// secretkey so both halves of the credential pair are available
		// from a single connection secret.
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if accesskey, ok := attr["accesskey"].(string); ok && accesskey != "" {
				conn["accesskey"] = []byte(accesskey)
			}
			return conn, nil
		}
	})
}
