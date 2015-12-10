package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/laszlovaspal/devops-challenge/awsutils"
	"github.com/laszlovaspal/devops-challenge/rest"
)

var (
	actionFlag    = flag.String("action", "list", "create/list/delete CloudFormation stack")
	stackNameFlag = flag.String("stackName", "cheppers-challenge", "name of CloudFormation stack")
	restPort      = flag.Int("restPort", 0, "REST API port")
)

func handleCommandLineArguments() {
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

func handleRESTRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/cloudformation/{stackId}/create", rest.HandleCreateCloudformationStackRequest)
	router.HandleFunc("/cloudformation/{stackId}/events", rest.HandleListCloudformationStackEventsRequest)
	router.HandleFunc("/cloudformation/{stackId}/delete", rest.HandleDeleteCloudformationStackRequest)
	router.HandleFunc("/cloudformation/list", rest.HandleListEC2InstancesRequest)

	port := ":" + strconv.Itoa(*restPort)
	log.Printf("Listening on %s...", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func main() {
	flag.Parse()

	log.Println("action", *actionFlag)
	log.Println("stackName", *stackNameFlag)

	if *restPort > 0 {
		handleRESTRequests()
	} else {
		handleCommandLineArguments()
	}
}
