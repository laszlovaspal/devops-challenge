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

With a set `$GOPATH` environment

    $ export GO15VENDOREXPERIMENT=1
    $ go get github.com/laszlovaspal/devops-challenge
    $ go install laszlovaspal/devops-challenge

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

To check whether Drupal is available on the stack

    $ devops-challenge -action=check -stackName=examplestack

To simulate outage by terminating an EC2 instance in the stack

    $ devops-challenge -action=simulateOutage

REST API
--------

To use the application via its REST API:

    $ devops-challenge -restPort 8080

Endpoints:

    http://localhost:8080/cloudformation/{stackId}/create
    http://localhost:8080/cloudformation/{stackId}/events
    http://localhost:8080/cloudformation/{stackId}/check
    http://localhost:8080/cloudformation/{stackId}/delete
    http://localhost:8080/cloudformation/simulateOutage
    http://localhost:8080/cloudformation/list

To create a new CloudFormation stack:

    $ curl -X POST http://localhost:8080/cloudformation/examplestack/create

To delete a CloudFormation stack:

    $ curl -X POST http://localhost:8080/cloudformation/examplestack/delete

To list CloudFormation stack events:

    $ curl -X GET http://localhost:8080/cloudformation/examplestack/events

To list running or pending EC2 instances:

    $ curl -X GET http://localhost:8080/cloudformation/list

To check whether Drupal is available on the stack

    $ curl -X GET http://localhost:8080/cloudformation/examplestack/check

To simulate outage by terminating an EC2 instance in the stack

    $ curl -X POST http://localhost:8080/cloudformation/simulateOutage
