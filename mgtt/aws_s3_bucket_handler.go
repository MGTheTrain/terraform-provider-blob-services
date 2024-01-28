package mgtt

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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
		return nil, fmt.Errorf("error creating AWS session: %v", err)
	}

	svc := s3.New(sess)

	return &AwsS3BucketHandler{
		svc:    svc,
		bucket: bucketName,
		region: region,
	}, nil
}

// CreateAwsS3Bucket creates an AWS S3 bucket
func (handler *AwsS3BucketHandler) CreateAwsS3Bucket() error {
	_, err := handler.svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(handler.bucket),
	})
	if err != nil {
		return fmt.Errorf("unable to create bucket %q: %v", handler.bucket, err)
	}

	fmt.Printf("Waiting for bucket %q to be created...\n", handler.bucket)

	err = handler.svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(handler.bucket),
	})
	if err != nil {
		return fmt.Errorf("error occurred while waiting for bucket to be created: %v", err)
	}

	fmt.Printf("Bucket %q successfully created\n", handler.bucket)
	return nil
}

// DeleteAwsS3Bucket deletes an AWS S3 bucket
func (handler *AwsS3BucketHandler) DeleteAwsS3Bucket() error {
	_, err := handler.svc.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(handler.bucket),
	})
	if err != nil {
		return fmt.Errorf("unable to delete bucket %q: %v", handler.bucket, err)
	}

	fmt.Printf("Waiting for bucket %q to be deleted...\n", handler.bucket)

	err = handler.svc.WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: aws.String(handler.bucket),
	})
	if err != nil {
		return fmt.Errorf("error occurred while waiting for bucket to be deleted: %v", err)
	}

	fmt.Printf("Bucket %q successfully deleted\n", handler.bucket)
	return nil
}
