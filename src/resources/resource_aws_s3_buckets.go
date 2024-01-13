package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsS3Buckets() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsS3BucketsCreate,
		Read:   resourceAwsS3BucketsRead,
		Update: resourceAwsS3BucketsUpdate,
		Delete: resourceAwsS3BucketsDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceAwsS3BucketsCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAwsS3BucketsRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAwsS3BucketsUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAwsS3BucketsDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
