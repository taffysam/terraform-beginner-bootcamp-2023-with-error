terraform {

  required_providers {
    terratowns = {
      source  = "local.providers/local/terratowns"
      version = "=1.0.0"
    }
  }
}

provider "terratowns" {
 endpoint = "http://localhost:4567"
  user_uuid = "e328f4ab-b99f-421c-84c9-4ccea042c7d1" 
  token = "9b49b3fb-b8e9-483c-b703-97ba88eef8e0"
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
  town = "gamers-grotto"
  content_version = 1



  
}