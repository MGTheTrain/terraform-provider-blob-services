package mgtt

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMgttAzurermRg() *schema.Resource {
	return &schema.Resource{
		Create: resourceMgttAzurermRgCreate,
		Read:   resourceMgttAzurermRgRead,
		Update: resourceMgttAzurermRgUpdate,
		Delete: resourceMgttAzurermRgDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceMgttAzurermRgCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	location := d.Get("location").(string)
	d.Set("name", name)
	d.Set("location", location)
	d.SetId(fmt.Sprintf("%s-%s", name, location))
	return nil
}

func resourceMgttAzurermRgRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMgttAzurermRgUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMgttAzurermRgDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
