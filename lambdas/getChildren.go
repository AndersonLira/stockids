package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type message struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	bk := &message{
		Title:   "Message",
		Message: "Showing message",
	}
	response, _ := json.Marshal(bk)
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil

}

func main() {
	lambda.Start(Create)
}
