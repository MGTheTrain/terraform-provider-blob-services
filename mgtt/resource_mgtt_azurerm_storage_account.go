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

func propertiesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     keyPolicySchema(),
			},
		},
	}
}

func keyPolicySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"key_expiration_period_in_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
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
			"sku": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     skuSchema(),
				Required: true,
			},
			"properties": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     propertiesSchema(),
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

	skuSet := d.Get("sku").(*schema.Set)
	skuName := ""
	skuTier := ""
	skuList := skuSet.List()
	if len(skuList) > 0 {
		skuMap := skuList[0].(map[string]interface{})
		skuName = skuMap["name"].(string)
		skuTier = skuMap["tier"].(string)
		fmt.Printf("SKU Name: %s, SKU Tier: %s\n", skuName, skuTier)
	}

	propertiesSet := d.Get("properties").(*schema.Set)
	keyExpirationPeriodInDays := 0
	propertiesList := propertiesSet.List()
	if len(propertiesList) > 0 {
		propertiesMap := propertiesList[0].(map[string]interface{})
		keyPolicySet := propertiesMap["key_policy"].(*schema.Set)
		keyPolicyList := keyPolicySet.List()
		if len(keyPolicyList) > 0 {
			keyPolicyMap := keyPolicyList[0].(map[string]interface{})
			keyExpirationPeriodInDays = keyPolicyMap["key_expiration_period_in_days"].(int)
		}
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
	azureStorageAccountHandler := NewAzureStorageAccountHandler(subscriptionID, accessToken)

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

	jsonString, err := ConvertMapToJSON(createRequestBody)
	if err != nil {
		return fmt.Errorf("Error converting map to JSON: %s", err)
	}

	err = azureStorageAccountHandler.CreateAzureStorageAccount(resourceGroupName, name, jsonString)

	if err != nil {
		return err
	}

	id := uuid.New()
	d.SetId(id.String())
	return nil
}

func resourceMgttAzurermStorageAccountRead(d *schema.ResourceData, m interface{}) error {
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

func resourceMgttAzurermStorageAccountUpdate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	resourceGroupName := d.Get("resource_group_name").(string)
	location := d.Get("location").(string)
	kind := d.Get("kind").(string)

	skuSet := d.Get("sku").(*schema.Set)
	skuName := ""
	skuTier := ""
	skuList := skuSet.List()
	if len(skuList) > 0 {
		skuMap := skuList[0].(map[string]interface{})
		skuName = skuMap["name"].(string)
		skuTier = skuMap["tier"].(string)
		fmt.Printf("SKU Name: %s, SKU Tier: %s\n", skuName, skuTier)
	}

	propertiesSet := d.Get("properties").(*schema.Set)
	keyExpirationPeriodInDays := 0
	propertiesList := propertiesSet.List()
	if len(propertiesList) > 0 {
		propertiesMap := propertiesList[0].(map[string]interface{})
		keyPolicySet := propertiesMap["key_policy"].(*schema.Set)
		keyPolicyList := keyPolicySet.List()
		if len(keyPolicyList) > 0 {
			keyPolicyMap := keyPolicyList[0].(map[string]interface{})
			keyExpirationPeriodInDays = keyPolicyMap["key_expiration_period_in_days"].(int)
		}
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
	azureStorageAccountHandler := NewAzureStorageAccountHandler(subscriptionID, accessToken)

	createRequestBody := map[string]interface{}{
		"properties": map[string]interface{}{
			"keyPolicy": map[string]interface{}{
				"keyExpirationPeriodInDays": keyExpirationPeriodInDays,
			},
		},
	}

	jsonString, err := ConvertMapToJSON(createRequestBody)
	if err != nil {
		return fmt.Errorf("Error converting map to JSON: %s", err)
	}

	err = azureStorageAccountHandler.UpdateAzureStorageAccount(resourceGroupName, name, jsonString)

	if err != nil {
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

	err := azureStorageAccountHandler.DeleteAzureStorageAccount(resourceGroupName, name)

	if err != nil {
		return err
	}
	return nil
}
