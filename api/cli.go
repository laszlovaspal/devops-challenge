package api

import (
	"io/ioutil"
	"log"

	"github.com/laszlovaspal/devops-challenge/awsutils"
)

// HandleCommandLineArguments takes care of command line arguments
func HandleCommandLineArguments(actionFlag string, stackNameFlag string) {

	switch actionFlag {

	case "events":
		log.Println(awsutils.GetCloudFormationStackEvents(stackNameFlag))

	case "create":
		template, _ := ioutil.ReadFile("MultiAZ2.template")
		drupalMulitAZTemplate := string(template)
		log.Println(awsutils.CreateNewCloudFormationStack(stackNameFlag, drupalMulitAZTemplate))

	case "delete":
		log.Println(awsutils.DeleteCloudFormationStack(stackNameFlag))

	case "list":
		log.Println(awsutils.ListRunningEC2Instances())

	case "check":
		isRunning, err := awsutils.IsDrupalRunningOnStack(stackNameFlag)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(isRunning)

	case "simulateOutage":
		awsutils.SimulateOutage()

	default:
		log.Println("Unknown action:", actionFlag)
	}
}
