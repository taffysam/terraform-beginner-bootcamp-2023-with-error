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

  cloud {
    organization = "Tafadzwa"
    workspaces {
      name = "terra-house-1"
    }
  }

}