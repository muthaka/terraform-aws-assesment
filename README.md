# terraform-aws-assesment
Creating a bucket and an instance.


# Creating an S3 bucket and an EC2 instance

To create the two resources, make sure that you have installed Terraform.
On the command line, run the following commands on the root directory of the project:

1. terraform init  
Command Prepare your working directory for other commands.

2. terraform validate
Command checks whether the configuration is valid.
  
3. terraform plan
Command shows changes required by the current configuration

4. terraform apply
Command creates or updates infrastructure

5. terraform destroy
Command destroys previously-created infrastructure

# Running test automation for the Terraform code to validating that both resources are tagged properly

Testing is done using Terratest. 
Make sure that Go unit testing framework is installed

Navigate to the test directory and run the following command in your CLI:
1. go test -v  
The command will run the file validate_tagging_test.go which should return ok if the tagging is properly done. 


# How to run a simple HTTP service written in Python
HTTP service is done using flask framework and written in Python3 language.

To run the test, make sure you have installed the following in your Environment using pip:
1. Flask framework 
2. boto3

To start the test,on your CLI make sure that you have activated your enviroment and navigate to the http_service directory.
Run the command: python3 endpoints.py  
The command will start the service which will run on the url[machine ip address]:8080 

Testing endpoint /tags

endpoint - url[machine ip address]:8080/tags

The endpoint will return the Tag name and Tag owner.


Testing endpoint /shutdown

endpoint - url[machine ip address]:8080/shutdown

The endpoint will shutdown the instance