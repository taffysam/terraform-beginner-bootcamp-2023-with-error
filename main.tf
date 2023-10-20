terraform {

  required_providers {
    terratowns = {
      source  = "local.providers/local/terratowns"
      version = "=1.0.0"
    }
  }
}

provider "terratowns" {
 endpoint = "https:terratowns.cloud/api"
  user_uuid = "87e8fc20-5f21-4b38-872b-ab8adfb49ed5" 
  token = "09c61961-c2b3-4459-a89b-0fbe2b3c48f0"
}


provider "aws" {
  region = "af-south-1"
}

#module "terrahouse_aws" {
#  source = "./modules/terrahouse_aws"
#  user_uuid = var.user_uuid
#  bucket_name = var.bucket_name
# validate_file_path.index_html_file_path = index_html_file_path
#  index_html_file_path = var.index_html_file_path
#  error_html_file_path = var.error_html_file_path
#  context_version  = 1
   
#}

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