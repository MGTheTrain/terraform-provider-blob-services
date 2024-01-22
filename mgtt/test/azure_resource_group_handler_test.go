package mgtt

import (
	"os"
	"testing"

	"github.com/MGTheTrain/terraform-provider-mgtt/mgtt"
	"github.com/stretchr/testify/assert"
)

func TestAzureResourceGroupHandler(t *testing.T) {
	// Read parameters from environment variables
	subscriptionID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	accessToken := os.Getenv("AZURE_ACCESS_TOKEN")
	resourceGroupName := "rg-test-100"

	if subscriptionID == "" || resourceGroupName == "" || accessToken == "" {
		t.Fatal("Missing required environment variables")
	}

	handler := mgtt.NewAzureResourceGroupHandler(subscriptionID, accessToken)

	createRequestBody := `{
		"sku": "West Europe"
	}`
	updateRequestBody := `{
		"name": "rg-test-1000"
	}`

	// Test PUT operation
	err := handler.CreateResourceGroup(resourceGroupName, createRequestBody)
	assert.NoError(t, err, "CreateResourceGroup should not return an error")

	// Test PATCH operation
	err = handler.UpdateResourceGroup(resourceGroupName, updateRequestBody)
	assert.NoError(t, err, "UpdateResourceGroup should not return an error")

	// Test GET operation
	err = handler.GetResourceGroup(resourceGroupName)
	assert.NoError(t, err, "GetResourceGroup should not return an error")

	// Test DELETE operation
	err = handler.DeleteResourceGroup(resourceGroupName)
	assert.NoError(t, err, "DeleteResourceGroup should not return an error")
}
