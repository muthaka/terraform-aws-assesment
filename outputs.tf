output "instance_name" {
  value = aws_instance.app_server.id
}

output "s3_bucket_name" {
  value = aws_s3_bucket.flugel-s3-bucket.id
}

output "s3_bucket_region" {
  value = aws_s3_bucket.flugel-s3-bucket.region
}
