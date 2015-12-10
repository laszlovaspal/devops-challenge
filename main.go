package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/laszlovaspal/devops-challenge/awsutils"
)

var actionFlag = flag.String("action", "list", "create/list/delete CloudFormation stack")
var stackNameFlag = flag.String("stackName", "cheppers-challenge", "name of CloudFormation stack")

func handleActionInput() {
	cfClient := awsutils.CreateNewCloudFormationClient()
	switch *actionFlag {

	case "events":
		fmt.Println(awsutils.GetCloudFormationStackEvents(cfClient, *stackNameFlag))

	case "create":
		template, _ := ioutil.ReadFile("Drupal_Multi_AZ_custom.template")
		drupalMulitAZTemplate := string(template)
		fmt.Println(awsutils.CreateNewCloudFormationStack(cfClient,
			*stackNameFlag, drupalMulitAZTemplate))

	case "delete":
		fmt.Println(awsutils.DeleteCloudFormationStack(cfClient, *stackNameFlag))

	case "list":
		ec2Client := awsutils.CreateNewEC2Client()
		fmt.Println(awsutils.ListRunningEC2Instances(ec2Client))

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
