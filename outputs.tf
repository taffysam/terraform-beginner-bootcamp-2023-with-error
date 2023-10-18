 #value = random_string.bucket_name.id
  
#}

#output "random_bucket_name" {
 # value = random_string.bucket_name.result
  
#}

output "error_html_file_path_output" {
  value = var.error_html_file_path
}

output "index_html_file_path_output" {
  value = var.index_html_file_path
}


#output "s3_website_endpoint" {
#  description = "The S3 Static Website hosting endpoint" 
#  value = module.terrahouse_aws.website_endpoint
#}

#output "cloudfront_url" {
#  description = "The CLoudFront Distribution Domain Name" 
#  value = module.terrahouse_aws.cloudfront_url
#}

#output "s3_website_endpoint" {
#  value = module.terrahouse_aws.website_endpoint
#}