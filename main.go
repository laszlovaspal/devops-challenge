package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/laszlovaspal/devops-challenge/cloudformationutils"
)

var actionFlag = flag.String("action", "list", "create/list/delete CloudFormation stack")
var stackNameFlag = flag.String("stackName", "cheppers-challenge", "name of CloudFormation stack")

func handleActionInput() {
	cfClient := cloudformationutils.CreateNewCloudFormationClient()
	switch *actionFlag {

	case "events":
		fmt.Println(cloudformationutils.GetCloudFormationStackEvents(cfClient, *stackNameFlag))

	case "create":
		template, _ := ioutil.ReadFile("Drupal_Multi_AZ_custom.template")
		drupalMulitAZTemplate := string(template)
		fmt.Println(cloudformationutils.CreateNewCloudFormationStack(cfClient,
			*stackNameFlag, drupalMulitAZTemplate))

	case "delete":
		fmt.Println(cloudformationutils.DeleteCloudFormationStack(cfClient, *stackNameFlag))

	default:
		fmt.Println("Unknown action:", *actionFlag)
	}
}

func main() {
	flag.Parse()

	fmt.Println("action", *actionFlag)
	fmt.Println("stackName", *stackNameFlag)

	handleActionInput()
}
