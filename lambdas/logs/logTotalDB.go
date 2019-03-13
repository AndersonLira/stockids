package main

import (
	"fmt"

	"github.com/andersonlira/stockids/db"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var logTotalTable = "skLogTotal"

//UpdateLogTotal for giving childID.
//If LogTotal not exists, it will create first
func UpdateLogTotal(childID string, score int) (logTotal LogTotal, err error) {
	ddb := db.GetDB()
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				N: aws.String(fmt.Sprintf("%d", score)),
			},
		},
		TableName: aws.String(logTotalTable),
		Key: map[string]*dynamodb.AttributeValue{
			"child_id": {
				S: aws.String(childID),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("ADD accumulated :r"),
	}

	_, err = ddb.UpdateItem(input)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	fmt.Println(err)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	return LogTotal{}, nil
}
