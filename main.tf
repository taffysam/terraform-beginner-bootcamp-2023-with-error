resource "random_string" "bucket_name" {
  length  = 25
  special = false
}

resource "aws_s3_bucket" "example" {
  bucket = random_string.bucket_name.result
}



