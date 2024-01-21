package mgtt

import (
	"os"
	"testing"

	"github.com/MGTheTrain/terraform-provider-mgtt/mgtt"
	"github.com/stretchr/testify/assert"
)

func TestAzureStorageAccountHandler(t *testing.T) {
	// Read parameters from environment variables
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	resourceGroupName := os.Getenv("AZURE_RESOURCE_GROUP_NAME")
	accountName := os.Getenv("AZURE_STORAGE_ACCOUNT_NAME")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")

	if subscriptionID == "" || resourceGroupName == "" || accountName == "" || accessToken == "" {
		t.Fatal("Missing required environment variables")
	}

	// Create a new instance of AzureStorageAccountHandler
	handler := mgtt.NewAzureStorageAccountHandler(subscriptionID, accessToken)

	// Example request body for PUT operation
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

	// Example request body for PATCH operation
	updateRequestBody := `{
		"properties": {
			"keyPolicy": {
				"keyExpirationPeriodInDays": 10
			}
		}
	}`

	// Test PUT operation
	err := handler.CreateAzureStorageAccount(resourceGroupName, accountName, createRequestBody)
	assert.NoError(t, err, "CreateAzureStorageAccount should not return an error")

	// Test PATCH operation
	err = handler.UpdateAzureStorageAccount(resourceGroupName, accountName, updateRequestBody)
	assert.NoError(t, err, "UpdateAzureStorageAccount should not return an error")

	// Test GET operation
	err = handler.GetAzureStorageAccount(resourceGroupName, accountName)
	assert.NoError(t, err, "GetAzureStorageAccount should not return an error")

	// Test DELETE operation
	err = handler.DeleteAzureStorageAccount(resourceGroupName, accountName)
	assert.NoError(t, err, "DeleteAzureStorageAccount should not return an error")
}
