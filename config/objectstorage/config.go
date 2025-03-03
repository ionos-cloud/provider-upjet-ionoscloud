package objectstorage

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures the object storage resources
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ionoscloud_s3_bucket", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
	})

	p.AddResourceConfigurator("ionoscloud_s3_bucket_policy", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_bucket_versioning", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_bucket_cors_configuration", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_bucket_lifecycle_configuration", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_bucket_public_access_block", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_bucket_website_configuration", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
	})

	p.AddResourceConfigurator("ionoscloud_s3_bucket_object_lock_configuration", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_bucket_server_side_encryption_configuration", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_object", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_object_copy", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
		r.References["bucket"] = config.Reference{
			TerraformName: "ionoscloud_s3_bucket",
		}
	})

	p.AddResourceConfigurator("ionoscloud_s3_key", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
		r.References["user_id"] = config.Reference{
			TerraformName: "ionoscloud_user",
		}
	})
	p.AddResourceConfigurator("ionoscloud_object_storage_accesskey", func(r *config.Resource) {
		r.ShortGroup = "objectstorage"
	})
}
