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
#  value = aws_s3_bucket.my_bucket.website_endpoint
#}
#
#output "s3_website_endpoint" {
#  value = "http://${aws_s3_bucket.bucket.website_endpoint}"
#}

#output "cloudfront_url" {
#  value = aws_cloudfront_distribution.website_bucket.domain_name
#}

