package mgtt

import (
	"bytes"
	"fmt"
	"net/http"
)

// AzureStorageManager represents a manager for handling Azure Storage operations.
type AzureStorageManager struct{}

// NewAzureStorageManager creates a new instance of AzureStorageManager.
func NewAzureStorageManager() *AzureStorageManager {
	return &AzureStorageManager{}
}

// HandleCreateAzureStorageAccount creates an Azure Storage account.
func (m *AzureStorageManager) HandleCreateAzureStorageAccount(subscriptionID, resourceGroupName, accountName, accessToken, requestBody string) {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s?api-version=2023-01-01",
		subscriptionID, resourceGroupName, accountName)

	m.sendHTTPRequest("PUT", url, []byte(requestBody), accessToken)
}

// HandleReadAzureStorageAccount reads information about an Azure Storage account.
func (m *AzureStorageManager) HandleReadAzureStorageAccount(subscriptionID, resourceGroupName, accountName, accessToken string) {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s?api-version=2023-01-01",
		subscriptionID, resourceGroupName, accountName)

	m.sendHTTPRequest("GET", url, nil, accessToken)
}

// HandleUpdateAzureStorageAccount updates an Azure Storage account.
func (m *AzureStorageManager) HandleUpdateAzureStorageAccount(subscriptionID, resourceGroupName, accountName, accessToken, requestBody string) {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s?api-version=2023-01-01",
		subscriptionID, resourceGroupName, accountName)

	m.sendHTTPRequest("PATCH", url, []byte(requestBody), accessToken)
}

// HandleDeleteAzureStorageAccount deletes an Azure Storage account.
func (m *AzureStorageManager) HandleDeleteAzureStorageAccount(subscriptionID, resourceGroupName, accountName, accessToken string) {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s?api-version=2023-01-01",
		subscriptionID, resourceGroupName, accountName)

	m.sendHTTPRequest("DELETE", url, nil, accessToken)
}

// sendHTTPRequest sends an HTTP request.
func (m *AzureStorageManager) sendHTTPRequest(method, url string, requestBody []byte, accessToken string) {
	var req *http.Request
	var err error

	if requestBody != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:")
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())
}
