resource "random_string" "bucket_name" {
  length  = 32
  special = false
}

resource "aws_s3_bucket" "bootcamp" {
  bucket = random_string.bucket_name.result

  tags = {
  Useruuid = var.user_uuid
  }
}

