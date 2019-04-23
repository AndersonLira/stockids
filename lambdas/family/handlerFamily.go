package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/andersonlira/stockids/lambdas"

	"github.com/andersonlira/stockids/model"
	"github.com/aws/aws-lambda-go/events"
)

const childParam = "childId"
const table = "skFamily"

//HandlerFamily implements GenericHandler
type HandlerFamily struct {
}

//Get interface implementation
func (h HandlerFamily) Get(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	childID, _ := request.PathParameters[childParam]
	payload := map[string]interface{}{
		"families":  getFamilies(childID),
		"total": getFamilyTotal(childID),
	}
	response, _ := json.Marshal(&payload)
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
}

//Create interface implementation
func (h HandlerFamily) Create(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	childID, errPath := request.PathParameters[childParam]
	if !errPath {
		return events.APIGatewayProxyResponse{Body: string(childID), StatusCode: http.StatusBadRequest}, nil
	}

	family := model.Family{}
	err := json.Unmarshal([]byte(request.Body), &family)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	family.ChildID = childID
	family.Date = time.Now().Unix()

	family, err = createFamily(family)

	if err != nil {

		if _, ok := err.(lambdas.ConflictError); ok {
			return lambdas.Conflict()
		}
		return events.APIGatewayProxyResponse{}, err
	}

	response, _ := json.Marshal(family)
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: http.StatusCreated}, nil

}

//Update interface implementation
func (h HandlerFamily) Update(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	panic("Create not implemented yet")
}

//Delete interface implementation
func (h HandlerFamily) Delete(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	panic("Create not implemented yet")
}
