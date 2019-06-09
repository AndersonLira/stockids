package main

import (
	"encoding/json"

	"github.com/andersonlira/stockids/lambdas"
	"github.com/aws/aws-lambda-go/events"
)

const childParam = "childId"

//UserHandler implements GenericHandler
type UserHandler struct {
}

//Get interface implementation
func (h UserHandler) Get(request events.APIGatewayProxyRequest, claims lambdas.Claims) (events.APIGatewayProxyResponse, error) {
	response, _ := json.Marshal(&claims)
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
}

//Create interface implementation
func (h UserHandler) Create(request events.APIGatewayProxyRequest, claims lambdas.Claims) (events.APIGatewayProxyResponse, error) {
	panic("Create not implemented yet")

}

//Update interface implementation
func (h UserHandler) Update(request events.APIGatewayProxyRequest, claims lambdas.Claims) (events.APIGatewayProxyResponse, error) {
	panic("Create not implemented yet")
}

//Delete interface implementation
func (h UserHandler) Delete(request events.APIGatewayProxyRequest, claims lambdas.Claims) (events.APIGatewayProxyResponse, error) {
	panic("Create not implemented yet")

}
