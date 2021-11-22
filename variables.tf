variable "aws_region" {
  description = "The AWS region to use to create resources."
  default     = "us-west-2"
}

variable "aws_profile" {
  description = "The AWS profile to use to create resources."
  default     = "default"
}

variable "ami_prefix" {
    type        = string
    description = "Creates a unique EC2 instance name beginning with the specified prefix."
    default     = "ami-830c94e3"
}

variable "instance_type" {
    type        = string
    description = "Defines the EC2 instance type of the resource"
    default     = "t2.micro"
}

variable "tag_name" {
    type        = string
    description = "Creates a name of the tag name"
    default     = "Flugel"
}

variable "tag_owner" {
    type        = string
    description = "Creates a name of the tag owner"
    default     = "InfraTeam"
}

variable "bucket_prefix" {
    type        = string
    description = "Creates a unique bucket name beginning with the specified prefix."
    default     = "flugel-s3bucket-"
}

variable "versioning" {
    type        = bool
    description = "A state of versioning."
    default     = true
}

variable "acl" {
    type        = string
    description = "Defaults to private"
    default     = "private"
}

