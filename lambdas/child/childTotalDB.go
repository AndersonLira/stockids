package main

import (
	"fmt"

	"github.com/andersonlira/stockids/db"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var childTotalTable = "skChildTotal"

//UpdateChildTotal for giving childID.
//If ChildTotal not exists, it will create first
func UpdateChildTotal(childID string, score int) (childTotal ChildTotal, err error) {
	ddb := db.GetDB()
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				N: aws.String(fmt.Sprintf("%d", score)),
			},
		},
		TableName: aws.String(childTotalTable),
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
	return ChildTotal{}, nil
}

func getChildTotal(childID string) ChildTotal {
	ddb := db.GetDB()
	result, err := ddb.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(childTotalTable),
		Key: map[string]*dynamodb.AttributeValue{
			"child_id": {
				S: aws.String(childID),
			},
		},
	})
	if err != nil {
		fmt.Println("Got error querying childTotals")
		fmt.Println(err.Error())
	}

	childTotal := ChildTotal{}
	dynamodbattribute.UnmarshalMap(result.Item, &childTotal)
	return childTotal
}
