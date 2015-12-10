package rest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/laszlovaspal/devops-challenge/awsutils"
)

var (
	cfClient              = awsutils.CreateNewCloudFormationClient()
	template, _           = ioutil.ReadFile("Drupal_Multi_AZ_custom.template")
	drupalMulitAZTemplate = string(template)
)

// HandleListEC2InstancesRequest returns a list of ec2 instances to the client
func HandleListEC2InstancesRequest(w http.ResponseWriter, request *http.Request) {
	log.Println("Serving list EC2 request")
	ec2Client := awsutils.CreateNewEC2Client()
	instances, err := awsutils.ListRunningEC2Instances(ec2Client)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, instances)
}

// HandleCreateCloudformationStackRequest creates a cloudformation stack
func HandleCreateCloudformationStackRequest(w http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		log.Println("Serving create CloudFormation stack request")
		vars := mux.Vars(request)
		response, err := awsutils.CreateNewCloudFormationStack(cfClient,
			vars["stackId"], drupalMulitAZTemplate)
		if err != nil {
			fmt.Fprintln(w, err)
		}
		fmt.Fprintln(w, response)
	}
}

// HandleDeleteCloudformationStackRequest deletes a cloudformation stack
func HandleDeleteCloudformationStackRequest(w http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		log.Println("Serving delete CloudFormation stack request")
		vars := mux.Vars(request)
		response, err := awsutils.DeleteCloudFormationStack(cfClient, vars["stackId"])
		if err != nil {
			fmt.Fprintln(w, err)
		}
		fmt.Fprintln(w, response)
	}
}

// HandleListCloudformationStackEventsRequest return a list of cloudformation stack events to the client
func HandleListCloudformationStackEventsRequest(w http.ResponseWriter, request *http.Request) {
	log.Println("Serving list CloudFormation events request")
	vars := mux.Vars(request)
	response, err := awsutils.GetCloudFormationStackEvents(cfClient, vars["stackId"])
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, response)
}
