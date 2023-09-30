variable "user_uid" {
  type        = number
  description = "User UID"
  
  validation {
    condition     = can(string(regex("^[0-9]+$", var.user_uid)))
    error_message = "User UID must be a valid integer."
  }
}
