package awsutils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// CreateNewEC2Client creates an ec2 client to communicate with AWS
func CreateNewEC2Client() *ec2.EC2 {
	return ec2.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
}

// ListRunningEC2Instances lists running ec2 instances
func ListRunningEC2Instances(ec2Client *ec2.EC2) (*ec2.DescribeInstancesOutput, error) {
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String("running"),
					aws.String("pending"),
				},
			},
		},
	}
	return ec2Client.DescribeInstances(params)
}
