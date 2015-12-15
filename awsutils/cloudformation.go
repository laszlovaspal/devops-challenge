package awsutils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

var cloudFormationClient = CreateNewCloudFormationClient()

// CreateNewCloudFormationClient creates new cloudformation client to communicate with AWS
func CreateNewCloudFormationClient() *cloudformation.CloudFormation {
	return cloudformation.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
}

// CreateNewCloudFormationStack creates a new cloudformation stack
func CreateNewCloudFormationStack(
	stackName string,
	cloudFormationTemplate string) (*cloudformation.CreateStackOutput, error) {

	params := &cloudformation.CreateStackInput{
		StackName: aws.String(stackName),
		Capabilities: []*string{
			aws.String("CAPABILITY_IAM"),
		},
		OnFailure: aws.String("DELETE"),
		Parameters: []*cloudformation.Parameter{
			{
				ParameterKey:     aws.String("InstanceType"),
				ParameterValue:   aws.String("t2.micro"),
				UsePreviousValue: aws.Bool(true),
			},
			{
				ParameterKey:     aws.String("KeyName"),
				ParameterValue:   aws.String("cheppers-challenge"),
				UsePreviousValue: aws.Bool(true),
			},
			{
				ParameterKey:     aws.String("SiteAdmin"),
				ParameterValue:   aws.String("admin"),
				UsePreviousValue: aws.Bool(true),
			},
			{
				ParameterKey:     aws.String("SiteEMail"),
				ParameterValue:   aws.String("admin@admin.hu"),
				UsePreviousValue: aws.Bool(true),
			},
			{
				ParameterKey:     aws.String("SitePassword"),
				ParameterValue:   aws.String("admin123"),
				UsePreviousValue: aws.Bool(true),
			},
			{
				ParameterKey:     aws.String("WebServerCapacity"),
				ParameterValue:   aws.String("2"),
				UsePreviousValue: aws.Bool(true),
			},
		},
		TemplateBody:     aws.String(cloudFormationTemplate),
		TimeoutInMinutes: aws.Int64(60),
	}
	return cloudFormationClient.CreateStack(params)
}

// GetCloudFormationStackEvents lists the events for a cloudformation stack
func GetCloudFormationStackEvents(stackName string) ([]*cloudformation.StackEvent, error) {

	descInput := &cloudformation.DescribeStackEventsInput{
		StackName: aws.String(stackName),
	}
	descOutput, err := cloudFormationClient.DescribeStackEvents(descInput)
	if err != nil {
		return nil, err
	}
	return descOutput.StackEvents, nil
}

// DeleteCloudFormationStack deletes a cloudformation stack
func DeleteCloudFormationStack(stackName string) (*cloudformation.DeleteStackOutput, error) {

	delInput := &cloudformation.DeleteStackInput{
		StackName: aws.String(stackName),
	}
	return cloudFormationClient.DeleteStack(delInput)
}

// GetURLOfCreatedStack returns the URL of the stack
func GetURLOfCreatedStack(stackName string) (string, error) {

	descInput := &cloudformation.DescribeStacksInput{
		StackName: aws.String(stackName),
	}
	descOutput, err := cloudFormationClient.DescribeStacks(descInput)
	if err != nil {
		return "", err
	}

	for _, stack := range descOutput.Stacks {
		for _, output := range stack.Outputs {
			return *output.OutputValue, nil
		}
	}
	return "", errors.New("Couldn't find URL for stack: " + stackName)
}

// IsDrupalRunningOnStack check whether the installed Drupal site is available on the stack
func IsDrupalRunningOnStack(stackName string) (bool, error) {

	url, err := GetURLOfCreatedStack(stackName)
	if err != nil {
		return false, err
	}

	response, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	return strings.Contains(string(responseBody), "Drupal"), nil
}
