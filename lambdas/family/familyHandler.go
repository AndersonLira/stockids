package main

import (
	"encoding/json"
	"net/http"

	"github.com/andersonlira/stockids/lambdas"
	"github.com/andersonlira/stockids/model"
	"github.com/aws/aws-lambda-go/events"
)

const childParam = "childId"

//FamilyHandler implements GenericHandler
type FamilyHandler struct {
}

//Get interface implementation
func (h FamilyHandler) Get(request events.APIGatewayProxyRequest, claims lambdas.Claims) (events.APIGatewayProxyResponse, error) {
	childID, _ := request.PathParameters[childParam]
	families := getFamilies(childID)
	response, _ := json.Marshal(&families)
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
}

//Create interface implementation
func (h FamilyHandler) Create(request events.APIGatewayProxyRequest, claims lambdas.Claims) (events.APIGatewayProxyResponse, error) {
	childID, errPath := request.PathParameters[childParam]
	if !errPath {
		return events.APIGatewayProxyResponse{Body: string(childID), StatusCode: http.StatusBadRequest}, nil
	}

	family := model.Family{}
	err := json.Unmarshal([]byte(request.Body), &family)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	family.UserID = "teste"

	family, err = createFamily(family)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response, _ := json.Marshal(family)
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: http.StatusCreated}, nil

}

//Update interface implementation
func (h FamilyHandler) Update(request events.APIGatewayProxyRequest, claims lambdas.Claims) (events.APIGatewayProxyResponse, error) {
	panic("Create not implemented yet")
}

//Delete interface implementation
func (h FamilyHandler) Delete(request events.APIGatewayProxyRequest, claims lambdas.Claims) (events.APIGatewayProxyResponse, error) {
	userID := "teste"
	id, _ := request.PathParameters["id"]
	if deleteAllFamiliesOfUser(id, userID) {
		return events.APIGatewayProxyResponse{Body: "true", StatusCode: http.StatusOK}, nil
	}
	return events.APIGatewayProxyResponse{Body: "false", StatusCode: http.StatusNotFound}, nil

}
