# terraform-provider-mgtt

## Table of Contents

- [Description](#description)
- [References](#references)
- [How to use](#how-to-use)

## Description

Sample repository implementing a terraform provider designed for managing blob storage services in Azure and AWS

## References

- [Writing Custom Terraform Providers](https://www.hashicorp.com/blog/writing-custom-terraform-providers). This link is deprecated but helpful to understand the basic concepts from the hashicorp founder.
- [How to develop/ test existing provider locally?](https://github.com/hashicorp/terraform-provider-aws/issues/5396)
- [terraform-provider-klayer Github repository](https://github.com/ldcorentin/terraform-provider-klayer). **13-01-2023** - Latest practical example
- [https://developer.hashicorp.com/terraform/plugin/sdkv2?collectionSlug=providers&productSlug=terraform](https://developer.hashicorp.com/terraform/plugin/sdkv2). Latest hashicorp documentation on Terraform Plugin SDKv2 for writing custom providers.

## How to use

### Compile custom provider

On modern Windows OS (version 10 or 11) run: 

```sh
go build -o terraform-provider-mgtt.exe
```

On Windows Unix systems run: 

```sh
go build -o terraform-provider-mgtt
```

### Copy provider executable to plugins directory 

Refer to [How to develop/ test existing provider locally?](https://github.com/hashicorp/terraform-provider-aws/issues/5396).

Navigate to:

```sh
cd terraform
```

On modern Windows OS with amd64 CPU architecture run:

```sh
mkdir terraform.d\plugins\windows_amd64\
cp terraform-provider-mgtt.exe terraform.d\plugins\windows_amd64\
```

On Unix systems run:

```sh
# Linux Ubuntu >=18.04 or debian >=11 with amd64 CPU architecture
mkdir -vp terraform.d/plugins/linux_amd64/
cp terraform-provider-mgtt terraform.d/plugins/linux_amd64/

# MacOS with amd64 CPU architecture 
mkdir -vp terraform.d/plugins/darwin_amd64/
cp terraform-provider-mgtt terraform.d/plugins/darwin_amd64/
```

### Test provider executable with hcl files and terraform commands

```sh
cd terraform
terraform init
terraform plan
terraform apply --auto-approve
terraform destroy --auto-approve
```