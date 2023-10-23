terraform {
  required_providers {
    terratowns = {
      source = "local.providers/local/terratowns"
      version = "1.0.0"
    }
  }
}


provider "terratowns" {
  endpoint  = var.terratowns_endpoint
  user_uuid = var.teacherseat_user_uuid
  token     = var.terratowns_access_token
}

# https://developer.hashicorp.com/terraform/language/modules/sources


provider "aws" {
  region = "af-south-1"
}

module "terrahouse_aws" {
source = "./modules/terrahouse_aws"
user_uuid = var.teacherseat_user_uuid
bucket_name = var.bucket_name
#validate_file_path.index_html_file_path = index_html_file_path
index_html_file_path = var.index_html_file_path
error_html_file_path = var.error_html_file_path
context_version  = var.context_version
   
}

resource "terratowns_home" "home" {
  name = "Tips to playing a Game"
  description = <<DESCRIPTION
- Type commands to interact with the game. For example, you can type "go north," "examine the chest," or "talk to the merchant."
- Your choices matter! Think carefully before making decisions.
- Save your progress by typing "save" at any time.
- To quit the game, type "quit."

Enjoy your adventure, and may you find the treasures that await you in AdventureQuest!

DESCRIPTION

  domain_name = module.terrahouse_aws.cloudfront_url
  town = "missingo"
  content_version = 1
  
}