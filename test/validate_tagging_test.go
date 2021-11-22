package test

import (
   "fmt"
   "path/filepath"
   "testing"
   "github.com/gruntwork-io/terratest/modules/aws"
   "github.com/gruntwork-io/terratest/modules/terraform"
   test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
   "github.com/stretchr/testify/assert"
)

// Function to test the tags.
func TestResourcesTags(t *testing.T) {
   t.Parallel()

   // Make a copy of the terraform module to a temporary directory. This allows running multiple tests in parallel
   // against the same terraform module.
   exampleFolder := test_structure.CopyTerraformFolderToTemp(t, "../", "")

   // Assign the expected Tag Name and Owner
   expectedName := fmt.Sprintf("Flugel")
   expectedOwner := fmt.Sprintf("InfraTeam")
   

   // Pick a random AWS region to test in which helps ensure the code to works in all regions.
   awsRegion := aws.GetRandomStableRegion(t, nil, nil)

  
   // Configure the options with default retryable errors to handle the most common retryable errors encountered in
   // terraform testing.
   planFilePath := filepath.Join(exampleFolder, "plan.out")
   terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
      // The path to where our Terraform code is located
      TerraformDir: "../",

      // Variables to pass to our Terraform code using -var options
      Vars: map[string]interface{}{
         "ami_prefix": expectedName,
         "bucket_prefix": expectedName,
      },

      // Environment variables to set when running Terraform
      EnvVars: map[string]string{
         "AWS_DEFAULT_REGION": awsRegion,
      },

      // Configure a plan file path so we can introspect the plan and make assertions about it.
      PlanFilePath: planFilePath,
   })

   // Run `terraform init`, `terraform plan`, and `terraform show` and fail the test if there are any errors
   plan := terraform.InitAndPlanAndShowWithStruct(t, terraformOptions)

   // Use the go struct to introspect the plan values in EC2 Instance.
   terraform.RequirePlannedValuesMapKeyExists(t, plan, "aws_instance.app_server")
   ec2Resource := plan.ResourcePlannedValuesMap["aws_instance.app_server"]
   ec2Tags := ec2Resource.AttributeValues["tags"].(map[string]interface{})

   // Use the go struct to introspect the plan values in S3 Bucket.
   terraform.RequirePlannedValuesMapKeyExists(t, plan, "aws_s3_bucket.flugel-s3-bucket")
   s3Resource := plan.ResourcePlannedValuesMap["aws_s3_bucket.flugel-s3-bucket"]
   s3Tags := s3Resource.AttributeValues["tags"].(map[string]interface{})

   // check if the tag name and owner are correct/matching
   assert.Equal(t, map[string]interface{}{"Name": expectedName, "Owner": expectedOwner}, s3Tags, ec2Tags)
   
}