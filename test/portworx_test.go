/**************************************************************************************************************
To Write a test file, use following link as a reference

https://github.com/terraform-ibm-modules/terraform-ibm-function/blob/main/test/cloud_function_test.go

***************************************************************************************************************/
/*
package test

import (
        "testing"

        "github.com/gruntwork-io/terratest/modules/terraform"
)

// An example of how to test the Terraform module to create cos instance in examples/instance using Terratest.
func TestAccIBMClusterE2E(t *testing.T) {
        t.Parallel()

        // Construct the terraform options with default retryable errors to handle the most common retryable errors in
        // terraform testing.
        terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
                // The path to where our Terraform code is located
                TerraformDir: "../examples/portworx",

                // Variables to pass to our Terraform code using -var options
                Vars: map[string]interface{}{
					ibmcloud_api_key        = "********"
					worker_nodes            = 4  // Number of workers
					storage_capacity        = 200  // In GBs
					storage_iops            = 10   // Must be a number, it will not be used unless a storage_profile is set to a custom profile
					storage_profile         = "10iops-tier"
					resource_group_name     = "default"
					region                  = "us-east"
					cluster_id              = "<cluster-id>"
					unique_id               = "roks-px-tf"
					create_external_etcd    = false
					etcd_username           = "portworxuser"
					etcd_password           = "portworxpassword"
					etcd_secret_name        = "px-etcd-certs" # don't change this
                },
        })

        // At the end of the test, run `terraform destroy` to clean up any resources that were created
		        defer terraform.Destroy(t, terraformOptions)

        // This will run `terraform init` and `terraform apply` and fail the test if there are any errors
        terraform.InitAndApply(t, terraformOptions)

        // Run `terraform output` to get the value of an output variable
        //portworx_flag := terraform.Output(t, terraformOptions, "portworx_is_ready")
        //if len(portworx_flag) <= 0 {
        //        t.Fatal("Wrong output")
        //}
        //fmt.Println("Portworx is installed: ", portworx_flag)
}
*/