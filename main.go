package main

import (
	"flag"
	"fmt"
)

var (
	actionFlag    string
	stackNameFlag string
)

func initFlags() {
	flag.StringVar(&actionFlag, "action", "list", "create/list/delete CloudFormation stack")
	flag.StringVar(&stackNameFlag, "stackName", "cheppers-challenge", "name of CloudFormation stack")
	flag.Parse()
}

func main() {
	fmt.Println("hello")
	initFlags()

	fmt.Println("action ", actionFlag)
	fmt.Println("stackName", stackNameFlag)
}
