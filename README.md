# terraform-provider-blob-services

## Table of Contents

- [Description](#description)
- [References](#references)
- [How to use](#how-to-use)

## Description

Sample repository implementing a terraform provider designed for managing blob storage services in Azure and AWS

## References

- [Writing Custom Terraform Providers](https://www.hashicorp.com/blog/writing-custom-terraform-providers). This link is deprecated but helpful to understand the basic concepts from the hashicorp founder.
- [How to Develop a Custom Provider in Terraform](https://www.infracloud.io/blogs/developing-terraform-custom-provider/). Deprecated link but helps understanding local testing of custom providers. Therefore consider section `Steps to copy provider executable to plugins directory`.
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

Refer to section `Steps to copy provider executable to plugins directory` of [How to Develop a Custom Provider in Terraform](https://www.infracloud.io/blogs/developing-terraform-custom-provider/).

On modern Windows OS with amd64 CPU architecture run:

```sh
mkdir $env:APPDATA\terraform.d\plugins\local.com\mgthetrain\mgtt\1.0.0\windows_amd64\
cp terraform-provider-mgtt.exe $env:APPDATA\terraform.d\plugins\local.com\mgthetrain\mgtt\1.0.0\windows_amd64\
```

On Unix systems run:

```sh
# Linux Ubuntu or debian with amd64 CPU architecture
mkdir -vp ~/.terraform.d/plugins/local.com/mgthetrain/mgtt/1.0.0/linux_amd64/
cp terraform-provider-mgtt ~/.terraform.d/plugins/local.com/mgthetrain/mgtt/1.0.0/linux_amd64/

# MacOS with amd64 CPU architecture 
mkdir -vp ~/.terraform.d/plugins/local.com/mgthetrain/mgtt/1.0.0/darwin_amd64/
cp terraform-provider-mgtt ~/.terraform.d/plugins/local.com/mgthetrain/mgtt/1.0.0/darwin_amd64/
```

