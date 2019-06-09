package main

import (
	"github.com/andersonlira/stockids/lambdas"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	gh := lambdas.GenericHandler{Handlerable: UserHandler{}}
	lambda.Start(gh.Handler)
}
