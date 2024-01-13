package terraform_provider_blob_services

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAwsS3Bucket() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsS3BucketCreate,
		Read:   resourceAwsS3BucketRead,
		Update: resourceAwsS3BucketUpdate,
		Delete: resourceAwsS3BucketDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAwsS3BucketCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAwsS3BucketRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAwsS3BucketUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAwsS3BucketDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
