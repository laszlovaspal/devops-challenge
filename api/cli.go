package api

import (
	"io/ioutil"
	"log"

	"github.com/laszlovaspal/devops-challenge/awsutils"
)

// HandleCommandLineArguments takes care of command line arguments
func HandleCommandLineArguments(actionFlag string, stackNameFlag string) {
	cfClient := awsutils.CreateNewCloudFormationClient()
	switch actionFlag {

	case "events":
		log.Println(awsutils.GetCloudFormationStackEvents(cfClient, stackNameFlag))

	case "create":
		template, _ := ioutil.ReadFile("Drupal_Multi_AZ_custom.template")
		drupalMulitAZTemplate := string(template)
		log.Println(awsutils.CreateNewCloudFormationStack(cfClient,
			stackNameFlag, drupalMulitAZTemplate))

	case "delete":
		log.Println(awsutils.DeleteCloudFormationStack(cfClient, stackNameFlag))

	case "list":
		ec2Client := awsutils.CreateNewEC2Client()
		log.Println(awsutils.ListRunningEC2Instances(ec2Client))

	default:
		log.Println("Unknown action:", actionFlag)
	}
}
