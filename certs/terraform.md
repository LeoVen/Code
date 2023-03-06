# Hashicorp Terraform Associate

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
* `terraform apply`
* `terraform destroy`

## Backend

> Defines where Terraform stores the state data files

* [State](https://developer.hashicorp.com/terraform/language/state/purpose)
* [Remote](https://developer.hashicorp.com/terraform/language/state/remote)
