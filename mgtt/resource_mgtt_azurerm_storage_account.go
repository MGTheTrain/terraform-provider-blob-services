package mgtt

import (
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func skuSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tier": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

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
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"resource_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			// "sku": &schema.Schema{
			// 	Type:     schema.TypeSet,
			// 	Elem:     skuSchema(),
			// 	Required: true,
			// },
			"sku_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sku_tier": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"kind": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceMgttAzurermStorageAccountCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	location := d.Get("location").(string)
	kind := d.Get("kind").(string)
	skuName := d.Get("sku_name").(string)
	skuTier := d.Get("sku_tier").(string)

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	azureStorageAccountHandler := NewAzureStorageAccountHandler(subscriptionID, accessToken)

	createRequestBody := map[string]interface{}{
		"sku": map[string]interface{}{
			"name": skuName,
			"tier": skuTier,
		},
		"kind":     kind,
		"location": location,
	}

	jsonString, err := ConvertMapToJSON(createRequestBody)
	if err != nil {
		return fmt.Errorf("Error converting map to JSON: %s", err)
	}

	err = azureStorageAccountHandler.CreateStorageAccount(resourceGroupName, name, jsonString)

	if err != nil {
		return err
	}

	id := uuid.New()
	d.SetId(id.String())

	if err := d.Set("name", name); err != nil {
		return err
	}
	if err := d.Set("resource_group_name", resourceGroupName); err != nil {
		return err
	}
	if err := d.Set("location", location); err != nil {
		return err
	}
	if err := d.Set("sku_name", skuName); err != nil {
		return err
	}
	if err := d.Set("sku_tier", skuTier); err != nil {
		return err
	}
	if err := d.Set("kind", kind); err != nil {
		return err
	}

	return nil
}

func resourceMgttAzurermStorageAccountRead(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	azureStorageAccountHandler := NewAzureStorageAccountHandler(subscriptionID, accessToken)

	err := azureStorageAccountHandler.GetStorageAccount(resourceGroupName, name)

	if err != nil {
		return err
	}
	return nil
}

func resourceMgttAzurermStorageAccountUpdate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	location := d.Get("location").(string)
	kind := d.Get("kind").(string)
	skuName := d.Get("sku_name").(string)
	skuTier := d.Get("sku_tier").(string)

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	azureStorageAccountHandler := NewAzureStorageAccountHandler(subscriptionID, accessToken)

	azureStorageAccountHandler.DeleteStorageAccount(resourceGroupName, name)

	createRequestBody := map[string]interface{}{
		"sku": map[string]interface{}{
			"name": skuName,
			"tier": skuTier,
		},
		"kind":     kind,
		"location": location,
	}

	jsonString, err := ConvertMapToJSON(createRequestBody)
	if err != nil {
		return fmt.Errorf("Error converting map to JSON: %s", err)
	}

	err = azureStorageAccountHandler.CreateStorageAccount(resourceGroupName, name, jsonString)

	if err != nil {
		return err
	}

	if err := d.Set("name", name); err != nil {
		return err
	}
	if err := d.Set("resource_group_name", resourceGroupName); err != nil {
		return err
	}
	if err := d.Set("location", location); err != nil {
		return err
	}
	if err := d.Set("sku_name", skuName); err != nil {
		return err
	}
	if err := d.Set("sku_tier", skuTier); err != nil {
		return err
	}
	if err := d.Set("kind", kind); err != nil {
		return err
	}

	return nil
}

func resourceMgttAzurermStorageAccountDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	azureStorageAccountHandler := NewAzureStorageAccountHandler(subscriptionID, accessToken)

	err := azureStorageAccountHandler.DeleteStorageAccount(resourceGroupName, name)

	if err != nil {
		return err
	}
	return nil
}
