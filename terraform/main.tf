resource "mgtt_aws_s3_bucket" "this" {
    name = "great"
}

resource "mgtt_azurerm_rg" "this" {
    name = "great"
    location = "West Europe"
}

resource "mgtt_azurerm_storage_account" "this" {
    name = "great"
}

resource "mgtt_azurerm_storage_account_container" "this" {
    name = "great"
}