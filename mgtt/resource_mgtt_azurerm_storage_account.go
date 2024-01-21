package mgtt

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMgttAzurermStorageAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceMgttAzurermStorageAccountCreate,
		Read:   resourceMgttAzurermStorageAccountRead,
		Update: resourceMgttAzurermStorageAccountUpdate,
		Delete: resourceMgttAzurermStorageAccountDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceMgttAzurermStorageAccountCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	if err := d.Set("name", name); err != nil {
		return err
	}
	d.SetId(name)
	return nil
}

func resourceMgttAzurermStorageAccountRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMgttAzurermStorageAccountUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMgttAzurermStorageAccountDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
