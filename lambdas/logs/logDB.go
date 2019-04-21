package main

import (
	"fmt"

	"github.com/andersonlira/stockids/db"
	"github.com/andersonlira/stockids/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func getLogs(childID string) []model.Log {

	ddb := db.GetDB()
	result, err := ddb.Query(&dynamodb.QueryInput{
		TableName:              aws.String(table),
		KeyConditionExpression: aws.String("child_id = :a"),
		Limit:            aws.Int64(30),
		ScanIndexForward: aws.Bool(false),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":a": {
				S: aws.String(childID),
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
	return logs
}

func createLog(log model.Log) (model.Log, error) {
	log.Message = "altered"
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
