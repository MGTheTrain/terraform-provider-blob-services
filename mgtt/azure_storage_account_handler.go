package mgtt

import (
	"bytes"
	"fmt"
	"net/http"
)

// AzureStorageAccountHandler represents a manager for handling Azure Storage operations.
type AzureStorageAccountHandler struct {
	SubscriptionID string
	AccessToken    string
}

// NewAzureStorageAccountHandler creates a new instance of AzureStorageAccountHandler.
func NewAzureStorageAccountHandler(subscriptionID, accessToken string) *AzureStorageAccountHandler {
	return &AzureStorageAccountHandler{
		SubscriptionID: subscriptionID,
		AccessToken:    accessToken,
	}
}

// CreateAzureStorageAccount creates an Azure Storage account.
func (m *AzureStorageAccountHandler) CreateAzureStorageAccount(resourceGroupName, accountName, requestBody string) error {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s?api-version=2023-01-01",
		m.SubscriptionID, resourceGroupName, accountName)

	return m.sendHTTPRequest("PUT", url, []byte(requestBody), m.AccessToken)
}

// GetAzureStorageAccount reads information about an Azure Storage account.
func (m *AzureStorageAccountHandler) GetAzureStorageAccount(resourceGroupName, accountName string) error {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s?api-version=2023-01-01",
		m.SubscriptionID, resourceGroupName, accountName)

	return m.sendHTTPRequest("GET", url, nil, m.AccessToken)
}

// UpdateAzureStorageAccount updates an Azure Storage account.
func (m *AzureStorageAccountHandler) UpdateAzureStorageAccount(resourceGroupName, accountName, requestBody string) error {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s?api-version=2023-01-01",
		m.SubscriptionID, resourceGroupName, accountName)

	return m.sendHTTPRequest("PATCH", url, []byte(requestBody), m.AccessToken)
}

// DeleteAzureStorageAccount deletes an Azure Storage account.
func (m *AzureStorageAccountHandler) DeleteAzureStorageAccount(resourceGroupName, accountName string) error {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s?api-version=2023-01-01",
		m.SubscriptionID, resourceGroupName, accountName)

	return m.sendHTTPRequest("DELETE", url, nil, m.AccessToken)
}

// sendHTTPRequest sends an HTTP request.
func (m *AzureStorageAccountHandler) sendHTTPRequest(method, url string, requestBody []byte, accessToken string) error {
	var req *http.Request
	var err error

	if requestBody != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return fmt.Errorf("Error creating request:  %s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error making request  %s:", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:")
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading from buffer  %s:", err)
	}
	fmt.Println(buf.String())

	return nil
}
