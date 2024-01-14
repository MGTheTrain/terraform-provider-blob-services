package mgtt

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"mgtt_aws_s3_bucket":                     resourceMgttAwsS3Bucket(),
			"mgtt_azurerm_storage_account":           resourceMgttAzurermStorageAccount(),
			"mgtt_azurerm_storage_account_container": resourceMgttAzurermStorageAccountContainer(),
			"mgtt_azurerm_rg":                        resourceMgttAzurermRg(),
		},
	}
}
