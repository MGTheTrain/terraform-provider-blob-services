# Azure storage account handler

## Table of Contents

- [Description](#description)
- [References](#references)
- [How to use](#how-to-use)

## Description

Cli applcation for managing Azure storage account services and Azure storage account service containers.  

## References

- [Create an Azure Storage account with the REST API](https://learn.microsoft.com/en-us/rest/api/storagerp/storage-sample-create-account)

## How to use

### Authenticate with az cli

If multi-factor authentication is enabled run following:

```sh
az login --tenant <your tenant id>
```

A browser window will appear, prompting you to log in with the Microsoft authentication provider. Optionally utilize a service principal.

### Display your bearer token after successfull authentication


```sh 
az account get-access-token --resource=https://management.azure.com
```

### Run the cli application

TBD
