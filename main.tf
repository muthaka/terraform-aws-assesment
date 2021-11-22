
provider "aws" {
  profile = var.aws_profile
  region  = var.aws_region
}

resource "aws_instance" "app_server" {
  ami           = var.ami_prefix
  instance_type = var.instance_type

  tags = {
    Name = var.tag_name
    Owner = var.tag_owner
  }

}

resource "aws_s3_bucket" "flugel-s3-bucket" {
  bucket_prefix = var.bucket_prefix
  acl = var.acl
  
   versioning {
    enabled = var.versioning
  }
  
  tags = {
    Name = var.tag_name
    Owner = var.tag_owner
  }
}

