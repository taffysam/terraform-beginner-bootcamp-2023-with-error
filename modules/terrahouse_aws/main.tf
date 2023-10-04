terraform {
  
}

#resource "random_string" "bucket_name" {
#  length  = 32
#  special = false
#}

#resource "aws_s3_bucket" "bootcamp3" {
#  bucket = random_string.bucket_name.result

#  tags = {
#  Useruuid = var.user_uuid
#  }
#}

#resource "aws_s3_bucket" "bootcamp1" {
#  bucket = "87e8fc20-5f21-4b38-872b-ab8adfb49ed5"
#}

#resource "aws_s3_bucket" "bootcamp2" {
#  bucket = "22e8fc20-5f21-4b38-872b-ab8adfb49ed5"
#}

resource "aws_s3_bucket_website_configuration" "website_configuration" {
bucket = "87e8fc20-5f21-4b38-872b-ab8adfb49ed5"

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "error.html"
  }



  #tags = {
  #Useruuid = var.user_uuid
  #}
}

