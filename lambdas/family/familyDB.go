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

const table = "skFamily"

func getFamilies(userID string) []model.Family {

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
	families := []model.Family{}
	for _, i := range result.Items {
		family := model.Family{}

		dynamodbattribute.UnmarshalMap(i, &family)

		families = append(families, family)

	}
	return families
}

func getFamiliesByQuery(queryInput *dynamodb.QueryInput) []model.Family {
	ddb := db.GetDB()
	result, err := ddb.Query(queryInput)

	if err != nil {
		fmt.Println("Got error querying families")
		fmt.Println(err.Error())
	}

	families := []model.Family{}
	for _, i := range result.Items {
		family := model.Family{}
		err = dynamodbattribute.UnmarshalMap(i, &family)
		families = append(families, family)
	}
	return families

}

func createFamily(family model.Family) (model.Family, error) {
	family.ID = str.NewUUID()
	family.CreatedAt = time.Now().Unix()
	av, err := dynamodbattribute.MarshalMap(family)
	if err != nil {
		return model.Family{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(table),
	}
	ddb := db.GetDB()
	_, err = ddb.PutItem(input)
	if err != nil {
		return model.Family{}, err
	}
	return family, nil
}

//DeleteFamily ...
func DeleteFamily(ID string) bool {
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
func UpdateFamily(ID string) (model.Family, error) {
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
		return GetFamily(ID)
	}
	return model.Family{}, nil
}

//GetFamily returns family with giving id
func GetFamily(ID string) (model.Family, error) {
	result, _ := db.GetDB().GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(ID),
			},
		},
	})
	family := model.Family{}
	err := dynamodbattribute.UnmarshalMap(result.Item, &family)
	return family, err
}

func defaultFamilyQuery() *dynamodb.QueryInput {
	return &dynamodb.QueryInput{
		TableName:                 aws.String(table),
		ScanIndexForward:          aws.Bool(false),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{},
	}
}
