package awsutils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// CreateNewCloudFormationClient creates new cloudformation client to communicate with AWS
func CreateNewCloudFormationClient() *cloudformation.CloudFormation {
	return cloudformation.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
}

// CreateNewCloudFormationStack creates a new cloudformation stack
func CreateNewCloudFormationStack(
	cloudFormationClient *cloudformation.CloudFormation,
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
func GetCloudFormationStackEvents(
	cloudFormationClient *cloudformation.CloudFormation,
	stackName string) ([]*cloudformation.StackEvent, error) {

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
func DeleteCloudFormationStack(
	cloudFormationClient *cloudformation.CloudFormation,
	stackName string) (*cloudformation.DeleteStackOutput, error) {

	delInput := &cloudformation.DeleteStackInput{
		StackName: aws.String(stackName),
	}
	return cloudFormationClient.DeleteStack(delInput)
}
