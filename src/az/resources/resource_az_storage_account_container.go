package terraform_provider_blob_services

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAzStorageAccountContainer() *schema.Resource {
	return &schema.Resource{
		Create: resourceAzStorageAccountContainerCreate,
		Read:   resourceAzStorageAccountContainerRead,
		Update: resourceAzStorageAccountContainerUpdate,
		Delete: resourceAzStorageAccountContainerDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAzStorageAccountContainerCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAzStorageAccountContainerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAzStorageAccountContainerUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAzStorageAccountContainerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
