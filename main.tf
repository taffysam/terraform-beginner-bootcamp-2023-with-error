terraform {

}

provider "aws" {
  region = "af-south-1"
}

module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
  user_uuid = var.user_uuid
  bucket_name = var.bucket_name
# validate_file_path.index_html_file_path = index_html_file_path
  index_html_file_path = var.index_html_file_path
  error_html_file_path = var.error_html_file_path
  context_version  = 1
 
  
}
