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

func getLogs(childID string) []model.Log {

	queryInput := defaultLogQuery()
	queryInput.KeyConditionExpression = aws.String("child_id = :a")
	queryInput.ExpressionAttributeValues[":a"] = &dynamodb.AttributeValue{
		S: aws.String(childID),
	}
	return getLogsByQuery(queryInput)
}

func getLogsByQuery(queryInput *dynamodb.QueryInput) []model.Log {
	ddb := db.GetDB()
	result, err := ddb.Query(queryInput)

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
	return logs

}

func createLog(log model.Log) (model.Log, error) {

	if existLastMinutes(log.ChildID) {
		return model.Log{}, &lambdas.ConflictError{}
	}

	av, err := dynamodbattribute.MarshalMap(log)
	if err != nil {
		return model.Log{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(table),
	}

	ddb := db.GetDB()
	_, err = ddb.PutItem(input)

	if err != nil {
		return model.Log{}, err
	}
	UpdateLogTotal(log.ChildID, log.Score)
	return log, nil
}

func existLastMinutes(childID string) (exist bool) {
	now := time.Now()
	before := now.Add(-5 * time.Minute).Unix()
	queryInput := defaultLogQuery()

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

	return len(getLogsByQuery(queryInput)) > 0
}

func defaultLogQuery() *dynamodb.QueryInput {
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
