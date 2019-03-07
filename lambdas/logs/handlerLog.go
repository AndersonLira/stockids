package main

import (
	"encoding/json"
	"fmt"

	"github.com/andersonlira/stockids/db"
	"github.com/andersonlira/stockids/model"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//HandlerLog implements GenericHandler
type HandlerLog struct {
}

//Get interface implementation
func (h HandlerLog) Get(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ddb := db.GetDB()
	result, err := ddb.Query(&dynamodb.QueryInput{
		TableName:              aws.String("skLog"),
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
	panic("Create not implemented yet")
}

//Update interface implementation
func (h HandlerLog) Update(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	panic("Create not implemented yet")
}

//Delete interface implementation
func (h HandlerLog) Delete(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	panic("Create not implemented yet")
}
