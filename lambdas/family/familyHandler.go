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
	families := getFamilies(claims.Email)
	response, _ := json.Marshal(&families)
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
}

//Create interface implementation
func (h FamilyHandler) Create(request events.APIGatewayProxyRequest, claims lambdas.Claims) (events.APIGatewayProxyResponse, error) {

	family := model.Family{}
	err := json.Unmarshal([]byte(request.Body), &family)

	if err != nil {
		return lambdas.BadRequest()
	}
	family.UserID = claims.Email

	family, err = createFamily(family)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response, _ := json.Marshal(family)
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: http.StatusCreated}, nil

}

//Update interface implementation
func (h FamilyHandler) Update(request events.APIGatewayProxyRequest, claims lambdas.Claims) (events.APIGatewayProxyResponse, error) {
	ID, errPath := request.PathParameters["id"]
	if !errPath {
		return lambdas.InvalidPathParam()
	}
	family := model.Family{}
	err := json.Unmarshal([]byte(request.Body), &family)

	if err != nil {
		return lambdas.BadRequest()
	}

	family, err = UpdateFamily(ID)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response, _ := json.Marshal(family)
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: http.StatusOK}, nil
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
