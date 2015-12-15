package api

import (
	"io/ioutil"
	"log"

	"github.com/laszlovaspal/devops-challenge/awsutils"
)

// HandleCommandLineArguments takes care of command line arguments
func HandleCommandLineArguments(actionFlag string, stackNameFlag string) {
	cfClient := awsutils.CreateNewCloudFormationClient()
	ec2Client := awsutils.CreateNewEC2Client()

	switch actionFlag {

	case "events":
		log.Println(awsutils.GetCloudFormationStackEvents(cfClient, stackNameFlag))

	case "create":
		template, _ := ioutil.ReadFile("MultiAZ2.template")
		drupalMulitAZTemplate := string(template)
		log.Println(awsutils.CreateNewCloudFormationStack(cfClient,
			stackNameFlag, drupalMulitAZTemplate))

	case "delete":
		log.Println(awsutils.DeleteCloudFormationStack(cfClient, stackNameFlag))

	case "list":
		log.Println(awsutils.ListRunningEC2Instances(ec2Client))

	case "check":
		isRunning, err := awsutils.IsDrupalRunningOnStack(cfClient, stackNameFlag)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(isRunning)

	case "simulateOutage":
		awsutils.SimulateOutage(ec2Client)

	default:
		log.Println("Unknown action:", actionFlag)
	}
}
