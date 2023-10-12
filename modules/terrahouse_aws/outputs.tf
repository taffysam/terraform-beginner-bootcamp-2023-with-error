#value = random_string.bucket_name.id
  
#}

#output "random_bucket_name" {
 # value = random_string.bucket_name.result
  
#}
#output "s3_website_endpoint" {
#value = module.terrahouse_aws.s3_website_endpoint  
#}



#output "cloudfront_url" {
# value = aws_cloudfront_distribution.website_bucket.domain_name
#}