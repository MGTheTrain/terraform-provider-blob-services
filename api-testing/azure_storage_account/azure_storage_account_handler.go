package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var (
	subscriptionID    string
	resourceGroupName string
	accountName       string
	accessToken       string
)

func main() {
	var rootCmd = &cobra.Command{Use: "azure_storage_account_handler"}

	rootCmd.PersistentFlags().StringVarP(&subscriptionID, "subscriptionID", "s", "", "Your Azure subscription ID")
	rootCmd.PersistentFlags().StringVarP(&resourceGroupName, "resourceGroupName", "g", "", "Your Azure resource group name")
	rootCmd.PersistentFlags().StringVarP(&accountName, "accountName", "a", "", "Your Azure Storage account name")
	rootCmd.PersistentFlags().StringVarP(&accessToken, "accessToken", "t", "", "Your Bearer access token")

	rootCmd.AddCommand(&cobra.Command{
		Use:   "manage",
		Short: "Manage Azure Storage Account",
		Run: func(cmd *cobra.Command, args []string) {
			handleAzureStorageAccount(subscriptionID, resourceGroupName, accountName, accessToken)
		},
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// handleAzureStorageAccount handles the Azure Storage Account management.
func handleAzureStorageAccount(subscriptionID, resourceGroupName, accountName, accessToken string) {
	fmt.Printf("Subscription ID: %s\n", subscriptionID)
	fmt.Printf("Resource Group Name: %s\n", resourceGroupName)
	fmt.Printf("Storage Account Name: %s\n", accountName)
	fmt.Printf("Access Token: %s\n", accessToken)

	if subscriptionID == "" || resourceGroupName == "" || accountName == "" || accessToken == "" {
		fmt.Println("Usage: azure_storage_account_handler manage -s <subscription_id> -g <resource_group_name> -a <storage_account_name> -t <bearer_access_token>")
		return
	}

	err := submitHTTPRequest(subscriptionID, resourceGroupName, accountName, accessToken)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// submitHTTPRequest submits an HTTP request to Azure Storage Account.
func submitHTTPRequest(subscriptionID, resourceGroupName, accountName, accessToken string) error {
	url := fmt.Sprintf("https://management.azure.com/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s?api-version=2018-02-01",
		subscriptionID, resourceGroupName, accountName)

	requestBody := map[string]interface{}{
		"sku": map[string]string{
			"name": "Standard_GRS",
		},
		"kind":     "StorageV2",
		"location": "eastus2",
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("Error encoding JSON: %v", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:")
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	fmt.Println(buf.String())

	return nil
}
