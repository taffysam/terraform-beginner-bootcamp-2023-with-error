variable "user_uuid" {
 type        = string
  description = "User UUID"

    validation {
       condition     = can(regex("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$", var.user_uuid))
       error_message  = "user_uuid must be in the format of a UUID (e.g., 123e4567-e89b-12d3-a456-426614174000)"

}

}

#variable "index_html_file_path" {
#  description = "Path to the index.html file"
#  type        = string

#  validation {
#    condition     = can(fileexists(var.index_html_file_path))
#    error_message = "The specified index.html file path does not exist."
#  }
#}

#variable "error_html_file_path" {
#  description = "Path to the index.html file"
#  type        = string

#  validation {
#    condition     = can(fileexists(var.error_file_path))
#    error_message = "The specified index.html file path does not exist."
#  }
#}


variable "index_html_file_path" {
  description = "Path to the index.html file"
  type        = string

  validation {
    condition     = length(var.index_html_file_path) > 0
    error_message = "The specified index.html file path must not be empty."
  }
}

resource "null_resource" "validate_file_path" {
  triggers = {
    index_html_file_path = var.index_html_file_path
  }


}


variable "error_file_path" {
  description = "Path to the index.html file"
  type        = string

  validation {
    condition     = length(var.error_file_path) > 0
    error_message = "The specified index.html file path must not be empty."
  }
}




