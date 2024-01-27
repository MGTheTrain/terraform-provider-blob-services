package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
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

	rootCmd.PersistentFlags().StringVarP(&accessKeyId, "access_key_id", "a", "", "Your AWS access key id")
	rootCmd.PersistentFlags().StringVarP(&secretAccessKey, "secret_access_key", "s", "", "Your AWS secret access key")
	rootCmd.PersistentFlags().StringVarP(&bucketName, "bucket_name", "b", "", "Your AWS S3 bucket name")
	rootCmd.PersistentFlags().StringVarP(&region, "region", "r", "", "The AWS region used for deploying the AWS S3 bucket")

	s3bucketCmd := &cobra.Command{
		Use:   "bucket",
		Short: "Manage AWS S3 bucket",
	}

	s3bucketCmd.AddCommand(&cobra.Command{
		Use:   "create",
		Short: "Create AWS S3 bucket by bucket name",
		Run: func(cmd *cobra.Command, args []string) {
			handleCreateAwsS3Bucket(accessKeyId, secretAccessKey, bucketName, region)
		},
	})

	rootCmd.AddCommand(s3bucketCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func handleCreateAwsS3Bucket(accessKeyId, secretAccessKey, bucketName, region string) {
	requestURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/", bucketName, region)

	signature, dateTime := generateSignature("PUT", requestURL, accessKeyId, secretAccessKey)

	headers := map[string]string{
		"Authorization": fmt.Sprintf("AWS4-HMAC-SHA256 Credential=%s/%s/%s/s3/aws4_request, SignedHeaders=host;x-amz-content-sha256;x-amz-date, Signature=%s",
			accessKeyId,
			dateTime.Format("20060102"),
			region,
			signature,
		),
		"X-Amz-Content-Sha256": "UNSIGNED-PAYLOAD",
		"X-Amz-Date":           dateTime.Format("20060102T150405Z"),
	}

	sendHTTPRequest("PUT", requestURL, nil, headers)
}

func generateSignature(method, requestURL, accessKeyId, secretAccessKey string) (string, time.Time) {
	dateTime := time.Now().UTC()
	date := dateTime.Format("20060102")

	// Canonical Request
	canonicalRequest := fmt.Sprintf("%s\n%s\n\nhost:%s\n\nhost\nUNSIGNED-PAYLOAD",
		method,
		"/",
		requestURL,
	)

	// String to Sign
	stringToSign := fmt.Sprintf("AWS4-HMAC-SHA256\n%s\n%s/%s/s3/aws4_request\n%x",
		dateTime.Format("20060102T150405Z"),
		date,
		region,
		sha256Hash(canonicalRequest),
	)

	// Signing Key
	signingKey := getSigningKey(secretAccessKey, dateTime, region)

	// Signing the String to Sign
	signature := hmacSHA256(signingKey, []byte(stringToSign))
	signatureString := base64.StdEncoding.EncodeToString(signature)

	return signatureString, dateTime
}

// Function to compute SHA-256 hash and return it as a byte slice
func sha256Hash(data string) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hasher.Sum(nil)
}

func getSigningKey(secretAccessKey string, dateTime time.Time, region string) []byte {
	date := dateTime.Format("20060102")
	kDate := hmacSHA256([]byte("AWS4"+secretAccessKey), []byte(date))
	kRegion := hmacSHA256(kDate, []byte(region))
	kService := hmacSHA256(kRegion, []byte("s3"))
	signingKey := hmacSHA256(kService, []byte("aws4_request"))
	return signingKey
}

func hmacSHA256(key, data []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

func sendHTTPRequest(method, url string, requestBody map[string]interface{}, headers map[string]string) {
	var req *http.Request
	var err error

	if requestBody != nil {
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json") // Add this line for content type
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Print all headers
	fmt.Println("Request Headers:")
	for key, values := range req.Header {
		fmt.Printf("%s: %s\n", key, strings.Join(values, ", "))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("URL:", url)
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
