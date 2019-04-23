package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/andersonlira/stockids/lambdas"

	"github.com/andersonlira/stockids/db"
	"github.com/andersonlira/stockids/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func getFamilies(childID string) []model.Family {

	queryInput := defaultFamilyQuery()
	queryInput.KeyConditionExpression = aws.String("child_id = :a")
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

	if existLastMinutes(family.ChildID) {
		return model.Family{}, lambdas.ConflictError{}
	}

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
	UpdateFamilyTotal(family.ChildID, family.Score)
	return family, nil
}

func existLastMinutes(childID string) (exist bool) {
	now := time.Now()
	before := now.Add(-5 * time.Minute).Unix()
	queryInput := defaultFamilyQuery()

	d := "date"
	queryInput.ExpressionAttributeNames = map[string]*string{
		"#d": &d,
	}
	queryInput.KeyConditionExpression = aws.String("child_id = :a and #d >= :d")

	queryInput.ExpressionAttributeValues[":a"] = &dynamodb.AttributeValue{
		S: aws.String(childID),
	}
	queryInput.ExpressionAttributeValues[":d"] = &dynamodb.AttributeValue{
		N: aws.String(strconv.FormatInt(before, 10)),
	}

	return len(getFamiliesByQuery(queryInput)) > 0
}

func defaultFamilyQuery() *dynamodb.QueryInput {
	return &dynamodb.QueryInput{
		TableName: aws.String(table),
		//KeyConditionExpression: aws.String("child_id = :a"),
		Limit:                     aws.Int64(30),
		ScanIndexForward:          aws.Bool(false),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			// ":a": {
			// 	S: aws.String(childID),
			// },
		},
	}
}
