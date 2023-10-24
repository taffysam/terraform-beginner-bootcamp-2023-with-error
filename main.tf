terraform {
  required_providers {
    terratowns = {
      source = "local-providers-local-terratowns"
      version = "1.0.0"
    }
  }

#backend "remote" {
# hostname = "app.terraform.io"
# organization = "Tafadzwa"
#  workspaces {
#   name = "terrahouse-2"
#}
#} 
 
  cloud {
    organization = "Tafadzwa"
    workspaces {
      name = "terrahouse-2"
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

resource "terratowns_home" "compass" {
  name = "Ho to play Compass Game"
  description = <<DESCRIPTION
- Type commands to interact with the game. For example, you can type "go north," "examine the chest," or "talk to the merchant."
- Your choices matter! Think carefully before making decisions.
- Save your progress by typing "save" at any time.
- To quit the game, type "quit."

Enjoy your adventure, and may you find the treasures that await you in AdventureQuest!

DESCRIPTION

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
  name = "How to play Hide and Seek"
  description = <<DESCRIPTION
Hide and seek is a classic outdoor game that is popular among children and is also enjoyed by people of all ages. Here are the basic rules and steps to play hide and seek:
Number of Players: Hide and seek can be played with a group of players, but it typically works best with three or more participants.
Objective: The main goal of the game is for one player (the "seeker") to find and tag the other players (the "hiders") who are hiding.
Setup:
Choose one player to be the "seeker." The other players will be the "hiders."
Decide on the boundaries or area where the game will be played. This area could be a backyard, a park, a house, or any other suitable location.
How to Play:
The seeker usually counts to a specific number (often 10, 20, or 30) with their eyes closed and facing a wall or tree, while the other players find places to hide.
While counting, the seeker should call out, "Ready or not, here I come!" once they have finished counting.
The seeker then begins searching for the hiders. The seeker can explore the designated area, trying to find and tag the hiders. The hiders should try to stay hidden.
If a hider is found and tagged by the seeker, they are "out" for that round and should return to a designated "base" or waiting area.
The game continues until all hiders have been found or until you decide to end the game.
Winning the Game:
The game doesn't typically have a winner or loser in the traditional sense. The fun comes from the thrill of hiding and the challenge of finding the other players. However, some players like to keep track of who was found first or who hid the longest in each round.
Variations and Tips:
You can play with different rules or variations. For example, in some versions of the game, hiders can run back to the base without getting tagged by the seeker to avoid being "out."
Hide and seek is often played outdoors, but you can also play it indoors if you have a suitable space.
Make sure to set boundaries and establish safety rules to prevent players from going to unsafe areas or getting lost.
The game is not just for children; adults can also enjoy playing hide and seek for a fun and active social activity.
Be respectful and considerate of others while hiding. Don't hide in places that could cause harm or damage property.

DESCRIPTION


  domain_name = module.terrahouse_aws.cloudfront_url
  town = "missingo"
  content_version = 1
  
}