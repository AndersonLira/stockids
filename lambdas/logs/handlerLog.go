package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/andersonlira/goutils/str"
	"github.com/andersonlira/stockids/db"
	"github.com/andersonlira/stockids/model"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const childParam = "childId"
const table = "skLog"

//HandlerLog implements GenericHandler
type HandlerLog struct {
}

//Get interface implementation
func (h HandlerLog) Get(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ddb := db.GetDB()
	result, err := ddb.Query(&dynamodb.QueryInput{
		TableName:              aws.String(table),
		KeyConditionExpression: aws.String("child_id = :a"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {
				S: aws.String("31ed531d-0df4-4fd8-98e0-02bb1d9e68b5"),
			},
		},
	})
	if err != nil {
		fmt.Println("Got error querying logs")
		fmt.Println(err.Error())
	}

	logs := []model.Log{}
	for _, i := range result.Items {
		log := model.Log{}
		err = dynamodbattribute.UnmarshalMap(i, &log)
		logs = append(logs, log)
	}
	response, _ := json.Marshal(&logs)
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

	log.ID = str.NewUUID()
	log.ChildID = childID
	log.Date = time.Now()
	av, err := dynamodbattribute.MarshalMap(log)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	// Create item in table Movies
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(table),
	}

	ddb := db.GetDB()
	_, err = ddb.PutItem(input)

	if err != nil {
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
