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
		// Not marked sensitive upstream; without this the credential is
		// published in plain text under status.atProvider.
		r.TerraformResource.Schema["secret_key"].Sensitive = true
		// The resource ID is the S3 access key; publish it with the secret.
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
		// Not marked sensitive upstream; without this the credential is
		// published in plain text under status.atProvider.
		r.TerraformResource.Schema["secretkey"].Sensitive = true
		// Publish the access key alongside the secret key.
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if accesskey, ok := attr["accesskey"].(string); ok && accesskey != "" {
				conn["accesskey"] = []byte(accesskey)
			}
			return conn, nil
		}
	})
}
