package mgtt

import (
	"os"
	"testing"

	"github.com/MGTheTrain/terraform-provider-mgtt/mgtt"
	"github.com/stretchr/testify/assert"
)

func TestStorageAccountHandler(t *testing.T) {
	// Read parameters from environment variables
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	accountName := "testaccount54321"
	resourceGroupName := "rg-test-100"

	if subscriptionID == "" || resourceGroupName == "" || accountName == "" || accessToken == "" {
		t.Fatal("Missing required environment variables")
	}

	resource_group_handler := mgtt.NewAzureResourceGroupHandler(subscriptionID, accessToken)
	handler := mgtt.NewAzureStorageAccountHandler(subscriptionID, accessToken)

	createRequestBody := `{
		"sku": {
			"name": "Standard_LRS",
			"tier": "Standard"
		},
		"kind": "StorageV2",
		"location": "West Europe",
		"properties": {
			"keyPolicy": {
				"keyExpirationPeriodInDays": 20
			}
		}
	}`
	updateRequestBody := `{
		"properties": {
			"keyPolicy": {
				"keyExpirationPeriodInDays": 10
			}
		}
	}`

	err := resource_group_handler.CreateResourceGroup(resourceGroupName, createRequestBody)
	assert.NoError(t, err, "CreateResourceGroup should not return an error")

	// Test PUT operation
	err = handler.CreateStorageAccount(resourceGroupName, accountName, createRequestBody)
	assert.NoError(t, err, "CreateStorageAccount should not return an error")

	// Test PATCH operation
	err = handler.UpdateStorageAccount(resourceGroupName, accountName, updateRequestBody)
	assert.NoError(t, err, "UpdateStorageAccount should not return an error")

	// Test GET operation
	err = handler.GetStorageAccount(resourceGroupName, accountName)
	assert.NoError(t, err, "GetStorageAccount should not return an error")

	// Test DELETE operation
	err = handler.DeleteStorageAccount(resourceGroupName, accountName)
	assert.NoError(t, err, "DeleteStorageAccount should not return an error")

	err = resource_group_handler.DeleteResourceGroup(resourceGroupName)
	assert.NoError(t, err, "DeleteResourceGroup should not return an error")
}
