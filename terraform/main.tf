resource "mgtt_aws_s3_bucket" "this" {
    name = "great"
}

resource "mgtt_azurerm_rg" "this" {
    name = "great"
    location = "West Europe"
}

resource "mgtt_azurerm_storage_account" "this" {
    name                 = "uniqacc2400"
    location             = "West Europe"
    resource_group_name  = "West Europe"
    
    sku {
        name = "Standard_LRS"
        tier = "Standard"
    }

    kind = "StorageV2"

    properties {
        key_policy {
            key_expiration_period_in_days = 20
        }
    }
}


resource "mgtt_azurerm_storage_account_container" "this" {
    name = "great"
    location = "West Europe"
    resource_group_name = "West Europe"
}