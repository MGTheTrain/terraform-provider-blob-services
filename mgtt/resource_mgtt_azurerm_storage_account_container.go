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

func getAzureStorageAccountContainerHandler() *AzureStorageAccountContainerHandler {
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	return NewAzureStorageAccountContainerHandler(subscriptionID, accessToken)
}

func resourceMgttAzurermStorageAccountContainerCreate(d *schema.ResourceData, m interface{}) error {
	name, accountName, resourceGroupName := extractStorageAccountContainerData(d)
	azureStorageAccountContainerHandler := getAzureStorageAccountContainerHandler()

	err := azureStorageAccountContainerHandler.CreateStorageAccountContainer(resourceGroupName, accountName, name, `{}`)
	if err != nil {
		return err
	}

	id := uuid.New()
	d.SetId(id.String())
	setStorageAccountContainerData(d, name, accountName, resourceGroupName)
	return nil
}

func resourceMgttAzurermStorageAccountContainerRead(d *schema.ResourceData, m interface{}) error {
	name, accountName, resourceGroupName := extractStorageAccountContainerData(d)
	azureStorageAccountContainerHandler := getAzureStorageAccountContainerHandler()

	err := azureStorageAccountContainerHandler.GetStorageAccountContainer(resourceGroupName, accountName, name)
	return err
}

func resourceMgttAzurermStorageAccountContainerUpdate(d *schema.ResourceData, m interface{}) error {
	oldName, oldAccountName, oldResourceGroupName := extractOldStorageAccountContainerData(d)
	azureStorageAccountContainerHandler := getAzureStorageAccountContainerHandler()

	err := azureStorageAccountContainerHandler.DeleteStorageAccountContainer(oldResourceGroupName, oldAccountName, oldName)
	if err != nil {
		return err
	}

	name, accountName, resourceGroupName := extractStorageAccountContainerData(d)
	err = azureStorageAccountContainerHandler.CreateStorageAccountContainer(resourceGroupName, accountName, name, `{}`)
	if err != nil {
		return err
	}

	setStorageAccountContainerData(d, name, accountName, resourceGroupName)
	return nil
}

func resourceMgttAzurermStorageAccountContainerDelete(d *schema.ResourceData, m interface{}) error {
	name, accountName, resourceGroupName := extractStorageAccountContainerData(d)
	azureStorageAccountContainerHandler := getAzureStorageAccountContainerHandler()

	err := azureStorageAccountContainerHandler.DeleteStorageAccountContainer(resourceGroupName, accountName, name)
	return err
}

// Helper functions

func extractStorageAccountContainerData(d *schema.ResourceData) (string, string, string) {
	return d.Get("name").(string), d.Get("account_name").(string), d.Get("resource_group_name").(string)
}

func setStorageAccountContainerData(d *schema.ResourceData, name, accountName, resourceGroupName string) error {
	if err := d.Set("name", name); err != nil {
		return err
	}
	if err := d.Set("account_name", accountName); err != nil {
		return err
	}
	if err := d.Set("resource_group_name", resourceGroupName); err != nil {
		return err
	}

	return nil
}

func extractOldStorageAccountContainerData(d *schema.ResourceData) (string, string, string) {
	oldName, _ := d.GetChange("name")
	oldAccountName, _ := d.GetChange("account_name")
	oldResourceGroupName, _ := d.GetChange("resource_group_name")

	return oldName.(string), oldAccountName.(string), oldResourceGroupName.(string)
}
