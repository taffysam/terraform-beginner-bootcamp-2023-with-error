terraform {

required_providers {
  terratowns = {
    source = "local.providers/local/terratowns"
    version = "1.0.0"
  }

    aws = {
     source = "hashicorp/aws"
     version = ">=5.18.1"
     }

    random = {
      source = "hashicorp/random"
      version = ">=3.5.1"
    }
  }
}


#cloud {
#   organization = "Tafadzwa"

#    workspaces {
#      name = "terra-house-2"
#    }
#  }



 
 

provider "terratowns" {
  endpoint = "http://localhost:4567"
  user_uuid="e328f4ab-b99f-421c-84c9-4ccea042c7d1" 
  token="9b49b3fb-b8e9-483c-b703-97ba88eef8e0"
  
}