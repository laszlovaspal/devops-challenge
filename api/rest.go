package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/laszlovaspal/devops-challenge/awsutils"
)

var (
	cfClient              = awsutils.CreateNewCloudFormationClient()
	template, _           = ioutil.ReadFile("Drupal_Multi_AZ_custom.template")
	drupalMulitAZTemplate = string(template)
)

// StartRestAPI sets up the routing and starts a http server
func StartRestAPI(restPort int) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/cloudformation/{stackId}/create", handleCreateCloudformationStackRequest)
	router.HandleFunc("/cloudformation/{stackId}/events", handleListCloudformationStackEventsRequest)
	router.HandleFunc("/cloudformation/{stackId}/delete", handleDeleteCloudformationStackRequest)
	router.HandleFunc("/cloudformation/list", handleListEC2InstancesRequest)

	port := ":" + strconv.Itoa(restPort)
	log.Printf("Listening on %s...", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// handleListEC2InstancesRequest returns a list of ec2 instances to the client
func handleListEC2InstancesRequest(w http.ResponseWriter, request *http.Request) {
	log.Println("Serving list EC2 request")
	ec2Client := awsutils.CreateNewEC2Client()
	instances, err := awsutils.ListRunningEC2Instances(ec2Client)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, instances)
}

// handleCreateCloudformationStackRequest creates a cloudformation stack
func handleCreateCloudformationStackRequest(w http.ResponseWriter, request *http.Request) {
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

// handleDeleteCloudformationStackRequest deletes a cloudformation stack
func handleDeleteCloudformationStackRequest(w http.ResponseWriter, request *http.Request) {
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

// handleListCloudformationStackEventsRequest return a list of cloudformation stack events to the client
func handleListCloudformationStackEventsRequest(w http.ResponseWriter, request *http.Request) {
	log.Println("Serving list CloudFormation events request")
	vars := mux.Vars(request)
	response, err := awsutils.GetCloudFormationStackEvents(cfClient, vars["stackId"])
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, response)
}
