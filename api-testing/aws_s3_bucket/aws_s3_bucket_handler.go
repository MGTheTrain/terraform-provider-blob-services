package main

// Based on:
// - https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/go/example_code/s3/s3_delete_bucket.go

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

var (
	accessKeyId     string
	secretAccessKey string
	bucketName      string
	region          string
)

// AwsS3BucketHandler is a struct for handling AWS S3 bucket operations
type AwsS3BucketHandler struct {
	svc    *s3.S3
	bucket string
	region string
}

// NewAwsS3BucketHandler creates a new AwsS3BucketHandler instance
func NewAwsS3BucketHandler(accessKeyID, secretAccessKey, bucketName, region string) (*AwsS3BucketHandler, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	return &AwsS3BucketHandler{
		svc:    svc,
		bucket: bucketName,
		region: region,
	}, nil
}

// handleCreateAwsS3Bucket creates an AWS S3 bucket
func (handler *AwsS3BucketHandler) handleCreateAwsS3Bucket() {
	_, err := handler.svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(handler.bucket),
	})
	if err != nil {
		exitErrorf("Unable to create bucket %q, %v", handler.bucket, err)
	}

	fmt.Printf("Waiting for bucket %q to be created...\n", handler.bucket)

	err = handler.svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(handler.bucket),
	})
	if err != nil {
		exitErrorf("Error occurred while waiting for bucket to be created, %v", handler.bucket)
	}

	fmt.Printf("Bucket %q successfully created\n", handler.bucket)
}

// handleDeleteAwsS3Bucket deletes an AWS S3 bucket
func (handler *AwsS3BucketHandler) handleDeleteAwsS3Bucket() {
	_, err := handler.svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(handler.bucket),
	})
	if err != nil {
		exitErrorf("Unable to delete bucket %q, %v", handler.bucket, err)
	}

	fmt.Printf("Waiting for bucket %q to be deleted...\n", handler.bucket)

	err = handler.svc.WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: aws.String(handler.bucket),
	})
	if err != nil {
		exitErrorf("Error occurred while waiting for bucket to be deleted, %v", handler.bucket)
	}

	fmt.Printf("Bucket %q successfully deleted\n", handler.bucket)
}

func main() {
	var rootCmd = &cobra.Command{Use: "aws_s3_bucket_handler"}

	rootCmd.PersistentFlags().StringVarP(&accessKeyId, "access_key_id", "a", "", "Your AWS access key id")
	rootCmd.PersistentFlags().StringVarP(&secretAccessKey, "secret_access_key", "s", "", "Your AWS secret access key")
	rootCmd.PersistentFlags().StringVarP(&bucketName, "bucket_name", "b", "", "Your AWS S3 bucket name")
	rootCmd.PersistentFlags().StringVarP(&region, "region", "r", "", "The AWS region used for deploying the AWS S3 bucket")

	s3bucketCmd := &cobra.Command{
		Use:   "bucket",
		Short: "Manage AWS S3 bucket",
		Run: func(cmd *cobra.Command, args []string) {
			handler, err := NewAwsS3BucketHandler(accessKeyId, secretAccessKey, bucketName, region)
			if err != nil {
				exitErrorf("Error creating AWS S3 bucket handler: %v", err)
			}

			handler.handleCreateAwsS3Bucket()
		},
	}

	s3bucketCmd.AddCommand(&cobra.Command{
		Use:   "delete",
		Short: "Delete AWS S3 bucket by bucket name",
		Run: func(cmd *cobra.Command, args []string) {
			handler, err := NewAwsS3BucketHandler(accessKeyId, secretAccessKey, bucketName, region)
			if err != nil {
				exitErrorf("Error creating AWS S3 bucket handler: %v", err)
			}

			handler.handleDeleteAwsS3Bucket()
		},
	})

	rootCmd.AddCommand(s3bucketCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
