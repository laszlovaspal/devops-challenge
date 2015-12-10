devops-challenge
================

This is a simple application written in Go that utilizes the Amazon AWS SDK to create, list and delete AWS instances.

What it does:
  - Spins up two t2.micro EC2 instances with load balancing using Amazon CloudFormation
  - Sets up a LAMP stack on the instances
  - Installs Drupal, supported by an Amazon S3 bucket for file storage
  - Uses Amazon RDS MySQL for Database
  - Provides command line arguments and a REST API for creating the above mentioned CloudFormation stack, listing the running instances, and deleting the stack

How to build
-----------

With a set $GOPATH environment

    $ go get github.com/laszlovaspal/devops-challenge
    $ go install

Example usage
-------------

To create a new CloudFormation stack:

    $ devops-challenge -action=create -stackName=examplestack

To delete a CloudFormation stack:

    $ devops-challenge -action=delete -stackName=examplestack

To list CloudFormation stack events:

    $ devops-challenge -action=events -stackName=examplestack

To list running or pending EC2 instances:

    $ devops-challenge -action=list -stackName=examplestack

REST API
--------

To use the application via its REST API:

    $ devops-challenge -restPort 8080

Endpoints:

    http://localhost:8080/cloudformation/{stackId}/create
    http://localhost:8080/cloudformation/{stackId}/events
    http://localhost:8080/cloudformation/{stackId}/delete
    http://localhost:8080/cloudformation/list
