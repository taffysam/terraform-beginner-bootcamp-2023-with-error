output "error_html_file_path_output" {
  value = var.error_html_file_path
}

output "index_html_file_path_output" {
  value = var.index_html_file_path
}


output "s3_website_endpoint" {
  description = "The S3 Static Website hosting endpoint" 
  value = module.terrahouse_aws.website_endpoint
}

output "cloudfront_url" {
  description = "The CLoudFront Distribution Domain Name" 
  value = module.terrahouse_aws.cloudfront_url
}

output "s3_bucket_name" {
  description = "Value that is unique."
  value       = aws_s3_bucket.example.id
}

output "UUID" {
  description = "Value that is unique."
  value       = var.teacherseat_user_uuid
}

output "s3bucket_website" {
  description = "the value of the website/url that is produced for the s3 bucket."
  value       = aws_s3_bucket_website_configuration.s3website.website_endpoint
}

output "domain_name" {
  value = aws_cloudfront_distribution.s3_distribution.domain_name
}