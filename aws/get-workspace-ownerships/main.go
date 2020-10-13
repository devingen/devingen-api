package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/devingen/devingen-api/aws"
	kimlik_aws "github.com/devingen/kimlik-api/aws"
)

func main() {
	serviceController := aws.GenerateController()
	lambda.Start(kimlik_aws.GenerateHandler(serviceController.GetWorkspaceOwnerships))
}
