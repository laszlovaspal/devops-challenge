package main

import (
	"flag"
	"log"

	"github.com/laszlovaspal/devops-challenge/api"
)

var (
	actionFlag    = flag.String("action", "list", "create/list/delete CloudFormation stack")
	stackNameFlag = flag.String("stackName", "cheppers-challenge", "name of CloudFormation stack")
	restPort      = flag.Int("restPort", 0, "REST API port")
)

func main() {
	flag.Parse()

	log.Println("action", *actionFlag)
	log.Println("stackName", *stackNameFlag)

	if *restPort > 0 {
		api.StartRestAPI(*restPort)
	} else {
		api.HandleCommandLineArguments(*actionFlag, *stackNameFlag)
	}
}
