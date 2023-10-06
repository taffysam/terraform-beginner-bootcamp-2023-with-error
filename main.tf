terraform {
  
}
module "terrahouse_aws" {
  source = "./modules/terrahouse_aws"
  user_uuid = var.user_uuid
#  bucket_name = var.bucket_name
# validate_file_path.index_html_file_path = index_html_file_path
index_html_file_path = "c:/Sofwares/Terraform Projects/terraform-beginner-bootcamp-2023-1/modules/terrahouse_aws/public/index.html"
error_file_path = var.error_html_file_path

}
