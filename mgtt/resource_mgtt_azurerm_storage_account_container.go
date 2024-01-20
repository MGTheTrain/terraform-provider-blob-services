package mgtt

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMgttAzurermStorageAccountContainer() *schema.Resource {
	return &schema.Resource{
		Create: resourceMgttAzurermStorageAccountContainerCreate,
		Read:   resourceMgttAzurermStorageAccountContainerRead,
		Update: resourceMgttAzurermStorageAccountContainerUpdate,
		Delete: resourceMgttAzurermStorageAccountContainerDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceMgttAzurermStorageAccountContainerCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	if err := d.Set("name", name); err != nil {
		return err
	}
	d.SetId(name)
	return nil
}

func resourceMgttAzurermStorageAccountContainerRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMgttAzurermStorageAccountContainerUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMgttAzurermStorageAccountContainerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
