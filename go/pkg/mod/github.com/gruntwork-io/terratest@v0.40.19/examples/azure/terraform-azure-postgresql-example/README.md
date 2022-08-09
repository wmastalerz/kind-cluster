# Terraform Azure PostgreSQL DB Example

This folder contains a Terraform module that deploys resources in [Azure](https://azure.microsoft.com/) to demonstrate how you can use Terratest to write automated tests for your Azure Terraform code.
This module deploys a database for PostgreSQL.

- A [Azure PostgreSQL Database](https://azure.microsoft.com/services/postgresql/).

Check out [test/azure/terraform_azure_postgresqldb_example_test.go](./../../../test/azure/terraform_azure_postgresqldb_example_test.go) to see how you can write automated tests for this module and validate the configuration of the parameters and options. 

**WARNING**: This module and the automated tests for it deploy real resources into your Azure account which can cost you money. 

## Running this module manually
1. Sign up for [Azure](https://azure.microsoft.com/).
1. Configure your Azure credentials using one of the [supported methods for Azure CLI
   tools](https://docs.microsoft.com/en-us/cli/azure/azure-cli-configuration?view=azure-cli-latest)
1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`.
1. Ensure [environment variables](../README.md#review-environment-variables) are available
1. Run `terraform init`
1. Run `terraform apply`
1. When you're done, run `terraform destroy`.


## Running automated tests against this module
1. Sign up for [Azure](https://azure.microsoft.com/)
1. Configure your Azure credentials using one of the [supported methods for Azure CLI
   tools](https://docs.microsoft.com/en-us/cli/azure/azure-cli-configuration?view=azure-cli-latest)
1. Install [Terraform](https://www.terraform.io/) and make sure it's on your `PATH`
1. Configure your Terratest [Go test environment](../README.md) 
1. `cd test/azure`
1. `go build terraform_azure_postgresql_example_test.go`
1. `go test -v -timeout 60m -tags azure -run TestPostgreSQLDatabase`
