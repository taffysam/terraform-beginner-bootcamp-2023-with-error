terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.21.0"
    }
  }
}

# # https://registry.terraform.io/providers/hashicorp/aws/latest/docs#provider-configuration
# provider "aws" {
#   # Configuration options
#   #   region     = "us-west-2"
#   #   access_key = "my-access-key"
#   #   secret_key = "my-secret-key"
# }

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity
data "aws_caller_identity" "current" {}

  

