# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudfront_origin_access_control

resource "aws_cloudfront_origin_access_control" "default" {
  name                              = "OAC ${var.bucket_name}"
  description                       = "Origin Access Controls for Static Website Hosting ${var.bucket_name}"
  origin_access_control_origin_type = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
}

locals {
  s3_origin_id = "MyS3Origin"
}

# Define your AWS S3 bucket resource here

resource "aws_cloudfront_distribution" "s3_distribution" {
  origin {
    domain_name = aws_s3_bucket.website_bucket.bucket_regional_domain_name
    origin_id   = local.s3_origin_id
  }

  enabled             = true
  is_ipv6_enabled     = true
  comment             = "Static Website hosting for: ${var.bucket_name}"
  default_root_object = "index.html"

  #logging_config {
  #  include_cookies = false
  #  bucket          = "tafadzwacloudfront-log-terraform2023"
    #bucket =  "${aws_s3_bucket.tafadzwacloudfront-log-terraform2023.id}.s3.amazonaws.com"
   #      }

  # Define any aliases (custom domain names) if needed

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = local.s3_origin_id

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "allow-all"
    min_ttl                = 0
    default_ttl            = 3600
    max_ttl                = 86400
  }

  # Define additional cache behaviors if needed

  price_class = "PriceClass_200"

  restrictions {
    geo_restriction {
      restriction_type = "none"
      locations        = []
    }
  }

  tags = {
    Useruuid = var.user_uuid
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }
}

resource "aws_s3_bucket_policy" "bucket_policy" {
  bucket = aws_s3_bucket.website_bucket.id  # Use 'id' instead of 'bucket' to reference the bucket resource

  policy = jsonencode({
    "Version" = "2012-10-17",
    "Statement" = [
        {
            "Sid" = "AllowCloudFrontServicePrincipalReadOnly",
            "Effect" = "Allow",
            "Principal" = {
                "Service" = "cloudfront.amazonaws.com"
            },
            "Action" = [
                "s3:GetObject",
                "s3:GetBucketAcl"
            ],
            "Resource" = [
                "arn:aws:s3:::${aws_s3_bucket.website_bucket.id}/*",
                "arn:aws:s3:::${aws_s3_bucket.website_bucket.id}"
            ],
            "Condition" = {
                "StringEquals" = {
                    "aws:SourceArn" = "arn:aws:cloudfront::${data.aws_caller_identity.current.account_id}:distribution/${aws_cloudfront_distribution.s3_distribution.id}"
                }
            }
        }
    ]
  })
}

resource "terraform_data" "invalidate_cache"{
  triggers_replace = terraform_data= content_version.output

  provisioner "local-exec" {
    # https://developer.hashicorp.com/terraform/language/expressions/strings
    command << ""

  
  } 

  
}


}