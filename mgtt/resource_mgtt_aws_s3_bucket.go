package mgtt

import (
	"os"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMgttAwsS3Bucket() *schema.Resource {
	return &schema.Resource{
		Create: resourceMgttAwsS3BucketCreate,
		Read:   resourceMgttAwsS3BucketRead,
		Update: resourceMgttAwsS3BucketUpdate,
		Delete: resourceMgttAwsS3BucketDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func getAwsS3BucketHandler() (*AwsS3BucketHandler, error) {
	accessKeyId := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	return NewAwsS3BucketHandler(accessKeyId, secretAccessKey, region)
}

func createS3Bucket(name string, handler *AwsS3BucketHandler) error {
	return handler.CreateAwsS3Bucket(name)
}

func deleteS3Bucket(name string, handler *AwsS3BucketHandler) error {
	return handler.DeleteAwsS3Bucket(name)
}

// Helper functions

func setS3BucketData(d *schema.ResourceData, name string) error {
	if err := d.Set("name", name); err != nil {
		return err
	}
	return nil
}

func extractOldS3BucketData(d *schema.ResourceData) string {
	oldName, _ := d.GetChange("name")
	return oldName.(string)
}

func resourceMgttAwsS3BucketCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	handler, err := getAwsS3BucketHandler()
	if err != nil {
		return err
	}

	err = createS3Bucket(name, handler)
	if err != nil {
		return err
	}

	id := uuid.New()
	d.SetId(id.String())

	if err := setS3BucketData(d, name); err != nil {
		return err
	}
	return nil
}

func resourceMgttAwsS3BucketRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMgttAwsS3BucketUpdate(d *schema.ResourceData, m interface{}) error {
	oldName := extractOldS3BucketData(d)
	handler, err := getAwsS3BucketHandler()
	if err != nil {
		return err
	}

	err = deleteS3Bucket(oldName, handler)
	if err != nil {
		return err
	}

	name := d.Get("name").(string)
	err = createS3Bucket(name, handler)
	if err != nil {
		return err
	}

	return setS3BucketData(d, name)
}

func resourceMgttAwsS3BucketDelete(d *schema.ResourceData, m interface{}) error {
	name := extractOldS3BucketData(d)
	handler, err := getAwsS3BucketHandler()
	if err != nil {
		return err
	}

	err = deleteS3Bucket(name, handler)
	if err != nil {
		return err
	}

	return nil
}
