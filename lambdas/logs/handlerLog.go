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
const table = "skLog"

//HandlerLog implements GenericHandler
type HandlerLog struct {
}

//Get interface implementation
func (h HandlerLog) Get(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	childID, _ := request.PathParameters[childParam]
	payload := map[string]interface{}{
		"logs":  getLogs(childID),
		"total": getLogTotal(childID),
	}
	response, _ := json.Marshal(&payload)
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
}

//Create interface implementation
func (h HandlerLog) Create(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	childID, errPath := request.PathParameters[childParam]
	if !errPath {
		return events.APIGatewayProxyResponse{Body: string(childID), StatusCode: http.StatusBadRequest}, nil
	}

	log := model.Log{}
	err := json.Unmarshal([]byte(request.Body), &log)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	log.ChildID = childID
	log.Date = time.Now().Unix()

	log, err = createLog(log)

	if err != nil {

		if _, ok := err.(lambdas.ConflictError); ok {
			return lambdas.Conflict()
		}
		return events.APIGatewayProxyResponse{}, err
	}

	response, _ := json.Marshal(log)
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: http.StatusCreated}, nil

}

//Update interface implementation
func (h HandlerLog) Update(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	panic("Create not implemented yet")
}

//Delete interface implementation
func (h HandlerLog) Delete(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	panic("Create not implemented yet")
}
