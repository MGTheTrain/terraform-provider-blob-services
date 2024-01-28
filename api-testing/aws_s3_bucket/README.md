# AWS S3 bucket handler

## Table of Contents

- [Description](#description)
- [References](#references)
- [How to use](#how-to-use)

## Description

Cli applcation for managing AWS S3 buckets.  

## References

- [Sample HTTP Request for general purpose buckets](https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html)
- [Signing and authenticating REST requests](https://docs.aws.amazon.com/AmazonS3/latest/userguide/RESTAuthentication.html)
- [Authenticating Requests: Using Query Parameters (AWS Signature Version 4)](https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-query-string-auth.html#query-string-auth-v4-signing-example)
. [Amazon S3 REST API with curl](https://czak.pl/2015/09/15/s3-rest-api-with-curl.html)

## How to use

### Preconditions

You need 
- an AWS subscription
- the `AWS_ACCESS_KEY_ID` environment variable
- the `AWS_SECRET_ACCESS_KEY` environment variable

### Build the cli application

On Windows OS:

```sh
go mod tidy
go build -o aws_s3_bucket_handler.exe
```

On Unix systems (MacOS or Ubuntu >=18.04/debian >=11):

```sh
go mod tidy
go build -o aws_s3_bucket_handler
```

### Run the cli application

```sh
# [C]reate an aws s3 bucket
aws_s3_bucket_handler bucket create -b <bucket name> -r <region>  -a <aws access key id> -s <aws secret access key>
aws_s3_bucket_handler bucket create --bucket_name <bucket name> --region <region>  --access_key_id <aws access key id> --secret_access_key <aws secret access key>

# [R]ead - Get an aws s3 bucket
aws_s3_bucket_handler bucket get -b <bucket name> -r <region>  -a <aws access key id> -s <aws secret access key>
aws_s3_bucket_handler bucket get --bucket_name <bucket name> --region <region>  --access_key_id <aws access key id> --secret_access_key <aws secret access key>

# [D]elete an aws s3 bucket 
aws_s3_bucket_handler bucket delete -b <bucket name> -r <region>  -a <aws access key id> -s <aws secret access key>
aws_s3_bucket_handler bucket delete --bucket_name <bucket name> --region <region>  --access_key_id <aws access key id> --secret_access_key <aws secret access key>
```