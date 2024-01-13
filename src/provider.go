package terraform_provider_blob_services

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"mgtt_aws_s3_bucket":                resourceAwsS3Bucket(),
			"mgtt_az_storage_account":           resourceAzStorageAccount(),
			"mgtt_az_storage_account_container": resourceAzStorageAccountContainer(),
		},
	}
}
