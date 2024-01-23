package mgtt

import (
	"os"

	"github.com/google/uuid"
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
			"account_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceMgttAzurermStorageAccountContainerCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	account_name := d.Get("account_name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	azureStorageAccountContainerHandler := NewAzureStorageAccountContainerHandler(subscriptionID, accessToken)

	err := azureStorageAccountContainerHandler.CreateStorageAccountContainer(resourceGroupName, account_name, name, `{}`)

	if err != nil {
		return err
	}

	id := uuid.New()
	d.SetId(id.String())

	if err := d.Set("name", name); err != nil {
		return err
	}
	if err := d.Set("account_name", account_name); err != nil {
		return err
	}
	if err := d.Set("resource_group_name", resourceGroupName); err != nil {
		return err
	}

	return nil
}

func resourceMgttAzurermStorageAccountContainerRead(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	account_name := d.Get("account_name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	azureStorageAccountContainerHandler := NewAzureStorageAccountContainerHandler(subscriptionID, accessToken)

	err := azureStorageAccountContainerHandler.GetStorageAccountContainer(resourceGroupName, account_name, name)

	if err != nil {
		return err
	}
	return nil
}

func resourceMgttAzurermStorageAccountContainerUpdate(d *schema.ResourceData, m interface{}) error {
	oldName, _ := d.GetChange("name")
	oldAccountName, _ := d.GetChange("account_name")
	oldResourceGroupName, _ := d.GetChange("resource_group_name")

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	azureStorageAccountContainerHandler := NewAzureStorageAccountContainerHandler(subscriptionID, accessToken)

	err := azureStorageAccountContainerHandler.DeleteStorageAccountContainer(oldResourceGroupName.(string), oldAccountName.(string), oldName.(string))

	if err != nil {
		return err
	}

	name := d.Get("name").(string)
	account_name := d.Get("account_name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)

	err = azureStorageAccountContainerHandler.CreateStorageAccountContainer(resourceGroupName, account_name, name, `{}`)

	if err != nil {
		return err
	}

	id := uuid.New()
	d.SetId(id.String())

	if err := d.Set("name", name); err != nil {
		return err
	}
	if err := d.Set("account_name", account_name); err != nil {
		return err
	}
	if err := d.Set("resource_group_name", resourceGroupName); err != nil {
		return err
	}

	return nil
}

func resourceMgttAzurermStorageAccountContainerDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	account_name := d.Get("account_name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	azureStorageAccountContainerHandler := NewAzureStorageAccountContainerHandler(subscriptionID, accessToken)

	err := azureStorageAccountContainerHandler.DeleteStorageAccountContainer(resourceGroupName, account_name, name)

	if err != nil {
		return err
	}
	return nil
}
