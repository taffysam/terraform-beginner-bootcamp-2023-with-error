terraform {
  required_providers {
   aws = {
      source = "hashicorp/aws"
     version = "5.18.1"
    }
  
    random = {
      source = "hashicorp/random"
      version = "3.5.1"
    }
  }

resource "aws_s3_bucket" "website_bucket"{
  bucket = var.bucket-name
  
  tags = {
    UserUuid = var.user_uuid
  }
}
}

module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
  user_uuid = var.user_uuid
  bucket_name = var.bucket_name
}
provider "random" {
  # Configuration options for the random provider, if needed.
}

resource "random_string" "bucket_name" {
  length  = 16
  special = false
}
