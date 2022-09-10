package main

import (
	"github.com/Hawkk3y3/lambdaApp/stack"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(stack.Handler)
}
