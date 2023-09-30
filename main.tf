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
}

resource "random_string" "bucket_name" {
  length  = 16
  special = false
}

output "random_bucket_id" {
    value = random_string.bucket_name.id
  
}

output "random_bucket_name" {
    value = random_string.bucket_name.result
  
}