terraform {
  
}

resource "random_string" "bucket_name" {
  length  = 32
  special = false
}

resource "aws_s3_bucket" "bootcamp3" {
 bucket = random_string.bucket_name.result
}
  #tags = {
  #Useruuid = var.user_uuid
  #}
#}

resource "aws_s3_bucket" "bootcamp" {
  bucket = "87e8fc20-5f21-4b38-872b-ab8adfb49ed5"
}

#resource "aws_s3_bucket" "bootcamp2" {
#  bucket = "22e8fc20-5f21-4b38-872b-ab8adfb49ed5"
#}
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_website_configuration
resource "aws_s3_bucket_website_configuration" "website_configuration" {
bucket = "web8fc20-5f21-4b38-872b-ab8adfb49ed5"

index_document {
    suffix = "index.html"
 }

error_document {
    key = "error.html"
  }

}

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_object
resource "aws_s3_object" "object" {
  bucket = "87e8fc20-5f21-4b38-872b-ab8adfb49ed5"
  key    = "index.html"
  
  source = var.index_html_file_path
    
  #source = "${path.root}/public.index.html"
  
  #source = "c:/Sofwares/Terraform Projects/terraform-beginner-bootcamp-2023-1/modules/terrahouse_aws/public/index.html"

    # The filemd5() function is available in Terraform 0.11.12 and later
  # For Terraform 0.11.11 and earlier, use the md5() function and the file() function:
  # etag = "${md5(file("path/to/file"))}"
  # etag = filemd5("path/to/file")


}



  #tags = {
  #Useruuid = var.user_uuid
  #}


  

