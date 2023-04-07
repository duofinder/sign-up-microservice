package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/duofinder/sign-up-microservice/handlers"
)

func main() {
	lambda.Start(handlers.Signup)
}
