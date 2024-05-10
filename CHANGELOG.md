# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 28-01-2024

### Added

- Initial setup considering [cli applications](./api-testing/) utilizing REST APIs or SDKs deriving handler structs used in [mgtt terraform provider resources](./mgtt/) 
- Add [tests](./mgtt/test/) and [workflows](./.github/workflows/)

## NOTE

- The `mgtt_azurerm_storage_account_container` is not yet working. Refer to comment in [resourceMgttAzurermStorageAccountContainerCreate function](./mgtt/resource_mgtt_azurerm_storage_account_container.go)