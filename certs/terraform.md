- [Hashicorp Terraform Associate](#hashicorp-terraform-associate)
  - [Exam](#exam)
  - [Workflow](#workflow)
  - [Backend](#backend)

# Hashicorp Terraform Associate

Sources

* [Terraform Language](https://developer.hashicorp.com/terraform/language)
* [CLI](https://developer.hashicorp.com/terraform/cli)

## Exam

1. Understand Infrastructure as Code concepts
2. Understand Terraform`s purpose
3. Understand Terraform basics
4. Use Terraform CLI
5. Interact with Terraform Modules
6. Navigate Terraform workflow
7. Implement and maintain state
8. Read, generate and modify configuration
9. Understand Terraform Cloud and Enterprise capabilities

## Workflow

0. (Optional) Import
1. Write IaC
2. Plan (review)
3. Apply (deploy)

* `terraform init`
    * Downloads ancillary components (modules, plugins)
    * Setup backend (state storing)
* `terraform plan`
  * Shows a plan of execution
  * Review changes
* `terraform apply`
  * Deploys changes to the infrastructure
  * Update the state file
* `terraform destroy`
  * Destroys resources tracked by the state file
  * Should be used with caution

## Backend

> Defines where Terraform stores the state data files

* [State](https://developer.hashicorp.com/terraform/language/state/purpose)
* [Remote](https://developer.hashicorp.com/terraform/language/state/remote)
