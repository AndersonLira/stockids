package main

import (
	"fmt"
	"time"

	"github.com/andersonlira/goutils/str"
	"github.com/andersonlira/stockids/db"
	"github.com/andersonlira/stockids/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

const table = "skChild"

func getChildren(userID string) []model.Child {

	filt := expression.Name("user_id").Equal(expression.Value(userID))
	expr, _ := expression.NewBuilder().WithFilter(filt).Build()
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(table),
	}

	result, _ := db.GetDB().Scan(params)
	fhildren := []model.Child{}
	for _, i := range result.Items {
		child := model.Child{}

		dynamodbattribute.UnmarshalMap(i, &child)

		fhildren = append(fhildren, child)

	}
	return fhildren
}

func getChildrenByQuery(queryInput *dynamodb.QueryInput) []model.Child {
	ddb := db.GetDB()
	result, err := ddb.Query(queryInput)

	if err != nil {
		fmt.Println("Got error querying fhildren")
		fmt.Println(err.Error())
	}

	fhildren := []model.Child{}
	for _, i := range result.Items {
		child := model.Child{}
		err = dynamodbattribute.UnmarshalMap(i, &child)
		fhildren = append(fhildren, child)
	}
	return fhildren

}

func createChild(child model.Child) (model.Child, error) {
	child.ID = str.NewUUID()
	child.CreatedAt = time.Now().Unix()
	av, err := dynamodbattribute.MarshalMap(child)
	if err != nil {
		return model.Child{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(table),
	}
	ddb := db.GetDB()
	_, err = ddb.PutItem(input)
	if err != nil {
		return model.Child{}, err
	}
	return child, nil
}

//DeleteChild ...
func DeleteChild(ID string) bool {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(ID),
			},
		},
		TableName: aws.String(table),
	}
	ddb := db.GetDB()
	_, err := ddb.DeleteItem(input)

	if err != nil {
		return false
	}
	return true
}

//
func UpdateChild(ID string) (model.Child, error) {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":d": {
				S: aws.String("changed"),
			},
		},
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(ID),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set description = :d"),
	}

	_, err := db.GetDB().UpdateItem(input)
	if err == nil {
		return GetChild(ID)
	}
	return model.Child{}, nil
}

//GetChild returns child with giving id
func GetChild(ID string) (model.Child, error) {
	result, _ := db.GetDB().GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(ID),
			},
		},
	})
	child := model.Child{}
	err := dynamodbattribute.UnmarshalMap(result.Item, &child)
	return child, err
}

func defaultChildQuery() *dynamodb.QueryInput {
	return &dynamodb.QueryInput{
		TableName:                 aws.String(table),
		ScanIndexForward:          aws.Bool(false),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{},
	}
}
