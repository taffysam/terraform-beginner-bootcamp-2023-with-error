#resource "random_string" "bucket_name" {
#  length  = 32
#  special = false
#}

#resource "aws_s3_bucket" "bootcamp" {
#  bucket = random_string.bucket_name.result

#  tags = {
#  Useruuid = var.user_uuid
#  }
#}

resource "aws_s3_bucket" "bootcamp" {
  bucket = "87e8fc20-5f21-4b38-872b-ab8adfb49ed5"

  tags = {
  Useruuid = var.user_uuid
  }
}
