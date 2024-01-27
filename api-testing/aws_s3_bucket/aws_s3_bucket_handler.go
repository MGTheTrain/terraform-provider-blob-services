package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
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

const (
	defaultEncoding = "UTF-8"
)

var encodedCharactersPattern = regexp.MustCompile(`\+|\*|%7E|%2F`)

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
	requestURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com", bucketName, region)

	signature, dateTime := generateSignature("PUT", requestURL, accessKeyId, secretAccessKey)
	queryParams := map[string]string{
		"X-Amz-Signature":     signature,
		"X-Amz-Algorithm":     "AWS4-HMAC-SHA256",
		"X-Amz-Credential":    urlEncode(fmt.Sprintf("%s/%s/%s/s3/aws4_request", accessKeyId, dateTime.Format("20060102"), region)),
		"X-Amz-Date":          dateTime.Format("20060102T150405Z"),
		"X-Amz-Expires":       "86400",
		"X-Amz-SignedHeaders": "host",
	}

	sendHTTPRequest("PUT", requestURL, nil, queryParams)
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
	stringToSign := fmt.Sprintf("AWS4-HMAC-SHA256\n%s\n%s/%s/s3/aws4_request\n%s",
		dateTime.Format("20060102T150405Z"),
		date,
		region,
		sha256Hex(canonicalRequest),
	)

	// Signing Key
	signingKey := getSigningKey(secretAccessKey, dateTime, region)

	// Signing the String to Sign
	h := hmac.New(sha256.New, signingKey)
	h.Write([]byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return signature, dateTime
}

// Function to compute SHA-256 hash and return it as a hexadecimal string
func sha256Hex(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func getSigningKey(secretAccessKey string, dateTime time.Time, region string) []byte {
	date := dateTime.Format("20060102")
	kDate := hmac.New(sha256.New, []byte("AWS4"+secretAccessKey))
	kDate.Write([]byte(date))
	kRegion := hmac.New(sha256.New, kDate.Sum(nil))
	kRegion.Write([]byte(region))
	kService := hmac.New(sha256.New, kRegion.Sum(nil))
	kService.Write([]byte("s3"))
	signingKey := hmac.New(sha256.New, kService.Sum(nil))
	signingKey.Write([]byte("aws4_request"))
	return signingKey.Sum(nil)
}

func sendHTTPRequest(method, url string, requestBody map[string]interface{}, queryParams map[string]string) {
	urlWithParams := addQueryParamsToURL(url, queryParams)

	var req *http.Request
	var err error

	if requestBody != nil {
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
		req, err = http.NewRequest(method, urlWithParams, bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json") // Add this line for content type
	} else {
		req, err = http.NewRequest(method, urlWithParams, nil)
	}

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("URL with params:", urlWithParams)
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

func addQueryParamsToURL(url string, queryParams map[string]string) string {
	if len(queryParams) == 0 {
		return url
	}

	var params []string
	for key, value := range queryParams {
		// Do not encode the values for certain parameters
		if key == "X-Amz-Credential" {
			params = append(params, fmt.Sprintf("%s=%s", key, value))
		} else {
			params = append(params, fmt.Sprintf("%s=%s", key, urlEncode(value)))
		}
	}

	// Construct the URL with encoded parameters
	urlWithParams := fmt.Sprintf("%s?%s", url, strings.Join(params, "&"))

	// Replace "&amp;" with "&" in the URL
	urlWithParams = strings.ReplaceAll(urlWithParams, "&amp;", "&")

	return urlWithParams
}

func urlEncode(value string) string {
	if value == "" {
		return ""
	}

	encoded := url.QueryEscape(value)

	replacementFunc := func(match string) string {
		switch match {
		case "+":
			return "%20"
		case "*":
			return "%2A"
		case "%7E":
			return "~"
		case "%2F":
			return "/"
		}
		return match
	}

	// Replace encoded slashes
	encoded = strings.ReplaceAll(encoded, "%2F", "/")

	encoded = encodedCharactersPattern.ReplaceAllStringFunc(encoded, replacementFunc)
	return encoded
}
