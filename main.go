package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/laszlovaspal/devops-challenge/awsutils"
)

var actionFlag = flag.String("action", "list", "create/list/delete CloudFormation stack")
var stackNameFlag = flag.String("stackName", "cheppers-challenge", "name of CloudFormation stack")
var restPort = flag.String("restPort", "", "REST API port")

func handleActionInput() {
	cfClient := awsutils.CreateNewCloudFormationClient()
	switch *actionFlag {

	case "events":
		log.Println(awsutils.GetCloudFormationStackEvents(cfClient, *stackNameFlag))

	case "create":
		template, _ := ioutil.ReadFile("Drupal_Multi_AZ_custom.template")
		drupalMulitAZTemplate := string(template)
		log.Println(awsutils.CreateNewCloudFormationStack(cfClient,
			*stackNameFlag, drupalMulitAZTemplate))

	case "delete":
		log.Println(awsutils.DeleteCloudFormationStack(cfClient, *stackNameFlag))

	case "list":
		ec2Client := awsutils.CreateNewEC2Client()
		log.Println(awsutils.ListRunningEC2Instances(ec2Client))

	default:
		log.Println("Unknown action:", *actionFlag)
	}
}

func main() {
	flag.Parse()

	log.Println("action", *actionFlag)
	log.Println("stackName", *stackNameFlag)

	handleActionInput()
}
