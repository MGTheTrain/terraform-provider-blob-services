package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	accessKeyId     string
	secretAccessKey string
	bucketName      string
	region          string
)

func main() {
	var rootCmd = &cobra.Command{Use: "aws_s3_bucket_handler"}

	rootCmd.PersistentFlags().StringVarP(&accessKeyId, "accessKeyId", "a", "", "Your AWS access key id")
	rootCmd.PersistentFlags().StringVarP(&secretAccessKey, "secretAccessKey", "s", "", "Your AWS secret access key")
	rootCmd.PersistentFlags().StringVarP(&bucketName, "bucketName", "b", "", "Your AWS S3 bucket name")
	rootCmd.PersistentFlags().StringVarP(&region, "region", "r", "", "The AWS region used for deploying the AWS S3 bucket")
	// Create a new aws_s3_bucket command
	storageAccountCmd := &cobra.Command{
		Use:   "aws_s3_bucket",
		Short: "Manage AWS S3 bucket",
	}

	// Add subcommands to the aws_s3_bucket command
	storageAccountCmd.AddCommand(&cobra.Command{
		Use:   "create",
		Short: "Create AWS S3 bucket by bucket name",
		Run: func(cmd *cobra.Command, args []string) {
			handleCreateAwsS3Bucket(accessKeyId, secretAccessKey, bucketName, region)
		},
	})

	storageAccountCmd.AddCommand(&cobra.Command{
		Use:   "get",
		Short: "Get an AWS S3 bucket by bucket name",
		Run: func(cmd *cobra.Command, args []string) {
			handleReadAwsS3Bucket(accessKeyId, secretAccessKey, bucketName, region)
		},
	})

	storageAccountCmd.AddCommand(&cobra.Command{
		Use:   "delete",
		Short: "Delete an AWS S3 bucket by bucket name",
		Run: func(cmd *cobra.Command, args []string) {
			handleDeleteAwsS3Bucket(accessKeyId, secretAccessKey, bucketName, region)
		},
	})

	rootCmd.AddCommand(storageAccountCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handleCreateAwsS3Bucket(accessKeyId, secretAccessKey, bucketName, region string) {
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com", bucketName, region)

	authorizationString := generateAuthorizationString("PUT", url, accessKeyId, secretAccessKey)

	sendHTTPRequest("PUT", url, nil, authorizationString)
}

func handleReadAwsS3Bucket(accessKeyId, secretAccessKey, bucketName, region string) {
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com", bucketName, region)

	authorizationString := generateAuthorizationString("GET", url, accessKeyId, secretAccessKey)

	sendHTTPRequest("GET", url, nil, authorizationString)
}

func handleDeleteAwsS3Bucket(accessKeyId, secretAccessKey, bucketName, region string) {
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com", bucketName, region)

	authorizationString := generateAuthorizationString("DELETE", url, accessKeyId, secretAccessKey)

	sendHTTPRequest("DELETE", url, nil, authorizationString)
}

func generateAuthorizationString(method, url, accessKeyId, secretAccessKey string) string {
	dateTime := time.Now().UTC().Format(time.RFC1123)
	stringToSign := fmt.Sprintf("%s\n\n\n%s\n%s", method, dateTime, url)

	h := hmac.New(sha1.New, []byte(secretAccessKey))
	h.Write([]byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return fmt.Sprintf("AWS %s:%s", accessKeyId, signature)
}

func sendHTTPRequest(method, url string, requestBody map[string]interface{}, accessToken string) {
	var req *http.Request
	var err error

	if requestBody != nil {
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
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
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		fmt.Println("Error reading from buffer:", err)
		return
	}
	fmt.Println(buf.String())
}
