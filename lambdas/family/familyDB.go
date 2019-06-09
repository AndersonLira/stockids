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
)

const table = "skFamily"

func getFamilies(childID string) []model.Family {

	queryInput := defaultFamilyQuery()
	queryInput.KeyConditionExpression = aws.String("id = :a")
	queryInput.ExpressionAttributeValues[":a"] = &dynamodb.AttributeValue{
		S: aws.String(childID),
	}
	return getFamiliesByQuery(queryInput)
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

func deleteAllFamiliesOfUser(id string, userID string) bool {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
			"user_id": {
				S: aws.String(userID),
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

func defaultFamilyQuery() *dynamodb.QueryInput {
	return &dynamodb.QueryInput{
		TableName:                 aws.String(table),
		ScanIndexForward:          aws.Bool(false),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{},
	}
}
