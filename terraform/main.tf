# resource "mgtt_aws_s3_bucket" "this" {
#     name = "great"
# }

resource "mgtt_azurerm_rg" "this" {
    name = "rg-test-5000"
    location = "West Europe"
}

resource "mgtt_azurerm_storage_account" "this" {
    name                 = "uniqacc2400"
    location             = mgtt_azurerm_rg.this.location
    resource_group_name  = mgtt_azurerm_rg.this.name
    kind = "StorageV2"
    

    sku_name = "Standard_LRS"
    sku_tier = "Standard"

    # sku {
    #     name = "Standard_LRS"
    #     tier = "Standard"
    # }
}


# resource "mgtt_azurerm_storage_account_container" "this" {
#     name = "great"
#     location = "West Europe"
#     resource_group_name = "West Europe"
# }