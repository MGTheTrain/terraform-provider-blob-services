package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"
	"time"
)

func calculate_hmac_sha1_signature(secretKey, request string) string {
	dateTime := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 -0700")
	stringToSign := fmt.Sprintf(request, dateTime)
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write([]byte(stringToSign))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature
}

func main() {
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")

	request := "GET\n\n\n%s\n/awsexamplebucket1/photos/puppy.jpg" // HTTP request to be signed
	signature := calculate_hmac_sha1_signature(secretKey, request)

	fmt.Printf("AWS %s:%s\n", accessKey, signature)
}
