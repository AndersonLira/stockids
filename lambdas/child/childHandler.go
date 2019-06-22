package main

import (
	"encoding/json"
	"net/http"

	"github.com/andersonlira/goutils/str"

	"github.com/andersonlira/godyn/db"

	"github.com/andersonlira/stockids/lambdas"
	"github.com/andersonlira/stockids/model"
	"github.com/aws/aws-lambda-go/events"
)

const childParam = "childId"

//ChildHandler implements GenericHandler
type ChildHandler struct {
}

//Get interface implementation
func (h ChildHandler) Get(request events.APIGatewayProxyRequest, claims lambdas.Claims) (events.APIGatewayProxyResponse, error) {
	fhildren := getChildren(claims.Username)
	response, _ := json.Marshal(&fhildren)
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
}

//Create interface implementation
func (h ChildHandler) Create(request events.APIGatewayProxyRequest, claims lambdas.Claims) (events.APIGatewayProxyResponse, error) {

	child := model.Child{}
	err := json.Unmarshal([]byte(request.Body), &child)

	if err != nil {
		return lambdas.BadRequest()
	}
	child.UserID = claims.Username
	child.ID = str.NewUUID()
	err = db.Create(&child)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response, _ := json.Marshal(child)
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: http.StatusCreated}, nil

}

//Update interface implementation
func (h ChildHandler) Update(request events.APIGatewayProxyRequest, claims lambdas.Claims) (events.APIGatewayProxyResponse, error) {
	ID, errPath := request.PathParameters["id"]

	if !errPath {
		return lambdas.InvalidPathParam()
	}
	child := model.Child{}
	err := json.Unmarshal([]byte(request.Body), &child)
	child.ID = ID
	child.UserID = claims.Username

	if err != nil {
		return lambdas.BadRequest()
	}

	err = db.Update(&child)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	response, _ := json.Marshal(child)
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: http.StatusOK}, nil
}

//Delete interface implementation
func (h ChildHandler) Delete(request events.APIGatewayProxyRequest, claims lambdas.Claims) (events.APIGatewayProxyResponse, error) {
	ID, errPath := request.PathParameters["id"]

	if !errPath {
		return lambdas.InvalidPathParam()
	}

	err := db.Delete(&model.Child{}, []interface{}{ID, claims.Username}...)

	result := make(map[string]bool)
	result["message"] = err == nil
	response, _ := json.Marshal(result)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: string(response), StatusCode: http.StatusOK}, nil
	}
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: http.StatusOK}, nil

}
