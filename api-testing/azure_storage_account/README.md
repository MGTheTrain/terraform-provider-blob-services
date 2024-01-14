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

### Preconditions

You need an Azure subscription, and you must have Owner permissions for that subscription. Optionally you can utilize a service principal.

### Authenticate with az cli

If multi-factor authentication is enabled run following:

```sh
az login --tenant <your tenant id>
```

A browser window will appear, prompting you to log in with the Microsoft authentication provider. Optionally utilize a service principal.

### Display your bearer token after successfull authentication

```sh 
az account get-access-token --resource https://graph.microsoft.com
```

### Build the cli application

On Windows OS:

```sh
 go build -o azure_storage_account_handler.exe
```

On Unix systems (MacOS or Ubuntu >=18.04/debian >=11):

```sh
 go build -o azure_storage_account_handler
```

### Run the cli application

TBD

