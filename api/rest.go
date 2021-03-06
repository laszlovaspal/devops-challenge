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
	template, _           = ioutil.ReadFile("MultiAZ2.template")
	drupalMulitAZTemplate = string(template)
)

// StartRestAPI sets up the routing and starts a http server
func StartRestAPI(restPort int) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/cloudformation/{stackId}/create", handleCreateCloudformationStackRequest).Methods("POST")
	router.HandleFunc("/cloudformation/{stackId}/events", handleListCloudformationStackEventsRequest).Methods("GET")
	router.HandleFunc("/cloudformation/{stackId}/check", handleCheckDrupalOnStackRequest).Methods("GET")
	router.HandleFunc("/cloudformation/{stackId}/delete", handleDeleteCloudformationStackRequest).Methods("POST")
	router.HandleFunc("/cloudformation/simulateOutage", handleSimulateOutageOnCloudformationStackRequest).Methods("POST")
	router.HandleFunc("/cloudformation/list", handleListEC2InstancesRequest).Methods("GET")

	port := ":" + strconv.Itoa(restPort)
	log.Printf("Listening on %s...", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// handleListEC2InstancesRequest returns a list of ec2 instances to the client
func handleListEC2InstancesRequest(w http.ResponseWriter, request *http.Request) {
	log.Println("Serving list EC2 request")

	instances, err := awsutils.ListRunningEC2Instances()
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, instances)
}

// handleCreateCloudformationStackRequest creates a cloudformation stack
func handleCreateCloudformationStackRequest(w http.ResponseWriter, request *http.Request) {
	log.Println("Serving create CloudFormation stack request")
	vars := mux.Vars(request)
	response, err := awsutils.CreateNewCloudFormationStack(vars["stackId"], drupalMulitAZTemplate)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, response)
}

// handleCheckDrupalOnStackRequest checks whether Drupal is running on a cloudformation stack
func handleCheckDrupalOnStackRequest(w http.ResponseWriter, request *http.Request) {
	log.Println("Serving check Drupal on CloudFormation stack request")
	vars := mux.Vars(request)
	isRunning, err := awsutils.IsDrupalRunningOnStack(vars["stackId"])
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, isRunning)
}

// handleDeleteCloudformationStackRequest deletes a cloudformation stack
func handleDeleteCloudformationStackRequest(w http.ResponseWriter, request *http.Request) {
	log.Println("Serving delete CloudFormation stack request")
	vars := mux.Vars(request)
	response, err := awsutils.DeleteCloudFormationStack(vars["stackId"])
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, response)
}

// handleSimulateOutageOnCloudformationStackRequest simulates an outage on a cloudformation stack by terminating an instance
func handleSimulateOutageOnCloudformationStackRequest(w http.ResponseWriter, request *http.Request) {
	log.Println("Serving simulate outage on CloudFormation stack request")
	awsutils.SimulateOutage()
	fmt.Fprintln(w, "Simulating outage")
}

// handleListCloudformationStackEventsRequest return a list of cloudformation stack events to the client
func handleListCloudformationStackEventsRequest(w http.ResponseWriter, request *http.Request) {
	log.Println("Serving list CloudFormation events request")
	vars := mux.Vars(request)
	response, err := awsutils.GetCloudFormationStackEvents(vars["stackId"])
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintln(w, response)
}
