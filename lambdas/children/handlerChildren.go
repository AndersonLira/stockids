package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andersonlira/stockids/lambdas"

	"github.com/andersonlira/goutils/str"
	"github.com/andersonlira/stockids/db"
	"github.com/andersonlira/stockids/model"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

//HandlerChildren implements GenericHandler
type HandlerChildren struct {
}

const table = "skChild"
const pathParam = "parentId"

//Get interface implementation
func (h HandlerChildren) Get(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	parentID, errPath := request.PathParameters[pathParam]
	if !errPath {
		return lambdas.InvalidPathParam()
	}

	ddb := db.GetDB()

	filt := expression.Name("parent_id").Equal(expression.Value(parentID))
	//proj := expression.NamesList(expression.Name("title"), expression.Name("year"), expression.Name("info.rating"))
	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	params := &dynamodb.ScanInput{
		ExpressionAttributeValues: expr.Values(),
		ExpressionAttributeNames:  expr.Names(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(table),
	}

	result, err := ddb.Scan(params)
	if err != nil {
		fmt.Println("Error getting children")
		fmt.Println(err.Error())
	}

	children := []model.Child{}
	for _, i := range result.Items {
		child := model.Child{}
		err = dynamodbattribute.UnmarshalMap(i, &child)
		children = append(children, child)
	}
	response, _ := json.Marshal(&children)
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 200}, nil
}

//Create interface implementation
func (h HandlerChildren) Create(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	parentID, errPath := request.PathParameters[pathParam]
	if !errPath {
		return events.APIGatewayProxyResponse{Body: string(parentID), StatusCode: http.StatusBadRequest}, nil
	}

	child := model.Child{}
	err := json.Unmarshal([]byte(request.Body), &child)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	child.ID = str.NewUUID()
	child.ParentID = parentID
	av, err := dynamodbattribute.MarshalMap(child)

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

	response, _ := json.Marshal(child)
	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: http.StatusCreated}, nil

}

//Update interface implementation
func (h HandlerChildren) Update(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	panic("Create not implemented yet")
}

//Delete interface implementation
func (h HandlerChildren) Delete(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	panic("Create not implemented yet")
}
