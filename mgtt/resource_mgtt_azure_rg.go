package mgtt

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
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
			"resource_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sku": &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Resource{
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
				},
				Required: true,
			},
			"kind": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"properties": &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keyPolicy": &schema.Schema{
							Type: schema.TypeMap,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"keyExpirationPeriodInDays": &schema.Schema{
										Type:     schema.TypeInt,
										Required: true,
									},
								},
							},
							Required: true,
						},
					},
				},
				Required: true,
			},
		},
	}
}

func convertMapToJSON(input map[string]interface{}) (string, error) {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func resourceMgttAzurermRgCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	location := d.Get("location").(string)
	sku := d.Get("sku").(map[string]interface{})
	skuName := sku["name"].(string)
	skuTier := sku["tier"].(string)
	kind := d.Get("kind").(string)

	properties := d.Get("properties").(map[string]interface{})
	keyPolicy := properties["keyPolicy"].(map[string]interface{})
	keyExpirationPeriodInDays := keyPolicy["keyExpirationPeriodInDays"].(int)

	if err := d.Set("name", name); err != nil {
		return err
	}
	if err := d.Set("resource_group_name", resourceGroupName); err != nil {
		return err
	}
	if err := d.Set("location", location); err != nil {
		return err
	}
	if err := d.Set("sku", map[string]interface{}{"name": skuName, "tier": skuTier}); err != nil {
		return err
	}
	if err := d.Set("kind", kind); err != nil {
		return err
	}
	if err := d.Set("properties", map[string]interface{}{"keyPolicy": map[string]interface{}{"keyExpirationPeriodInDays": keyExpirationPeriodInDays}}); err != nil {
		return err
	}

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")

	createRequestBody := map[string]interface{}{
		"sku": map[string]interface{}{
			"name": skuName,
			"tier": skuTier,
		},
		"kind":     kind,
		"location": location,
		"properties": map[string]interface{}{
			"keyPolicy": map[string]interface{}{
				"keyExpirationPeriodInDays": keyExpirationPeriodInDays,
			},
		},
	}

	azureStorageAccountHandler := NewAzureStorageAccountHandler(subscriptionID, accessToken)
	jsonString, err := convertMapToJSON(createRequestBody)
	if err != nil {
		return fmt.Errorf("Error converting map to JSON: %s", err)
	}

	err = azureStorageAccountHandler.CreateAzureStorageAccount(resourceGroupName, name, jsonString)

	if err != nil {
		return err
	}

	// set
	id := uuid.New()
	d.SetId(id.String())
	return nil
}

func resourceMgttAzurermRgRead(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")

	azureStorageAccountHandler := NewAzureStorageAccountHandler(subscriptionID, accessToken)
	err := azureStorageAccountHandler.GetAzureStorageAccount(resourceGroupName, name)

	if err != nil {
		return err
	}

	return nil
}

func resourceMgttAzurermRgUpdate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	location := d.Get("location").(string)
	sku := d.Get("sku").(map[string]interface{})
	skuName := sku["name"].(string)
	skuTier := sku["tier"].(string)
	kind := d.Get("kind").(string)

	properties := d.Get("properties").(map[string]interface{})
	keyPolicy := properties["keyPolicy"].(map[string]interface{})
	keyExpirationPeriodInDays := keyPolicy["keyExpirationPeriodInDays"].(int)

	if err := d.Set("name", name); err != nil {
		return err
	}
	if err := d.Set("resource_group_name", resourceGroupName); err != nil {
		return err
	}
	if err := d.Set("location", location); err != nil {
		return err
	}
	if err := d.Set("sku", map[string]interface{}{"name": skuName, "tier": skuTier}); err != nil {
		return err
	}
	if err := d.Set("kind", kind); err != nil {
		return err
	}
	if err := d.Set("properties", map[string]interface{}{"keyPolicy": map[string]interface{}{"keyExpirationPeriodInDays": keyExpirationPeriodInDays}}); err != nil {
		return err
	}

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")

	createRequestBody := map[string]interface{}{
		"properties": map[string]interface{}{
			"keyPolicy": map[string]interface{}{
				"keyExpirationPeriodInDays": keyExpirationPeriodInDays,
			},
		},
	}

	azureStorageAccountHandler := NewAzureStorageAccountHandler(subscriptionID, accessToken)
	jsonString, err := convertMapToJSON(createRequestBody)
	if err != nil {
		return fmt.Errorf("Error converting map to JSON: %s", err)
	}

	err = azureStorageAccountHandler.UpdateAzureStorageAccount(resourceGroupName, name, jsonString)

	if err != nil {
		return err
	}
	return nil
}

func resourceMgttAzurermRgDelete(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)

	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")

	azureStorageAccountHandler := NewAzureStorageAccountHandler(subscriptionID, accessToken)
	err := azureStorageAccountHandler.DeleteAzureStorageAccount(resourceGroupName, name)

	if err != nil {
		return err
	}

	return nil
}
