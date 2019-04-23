package main

import (
	"fmt"

	"github.com/andersonlira/stockids/db"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var familyTotalTable = "skFamilyTotal"

//UpdateFamilyTotal for giving childID.
//If FamilyTotal not exists, it will create first
func UpdateFamilyTotal(childID string, score int) (familyTotal FamilyTotal, err error) {
	ddb := db.GetDB()
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				N: aws.String(fmt.Sprintf("%d", score)),
			},
		},
		TableName: aws.String(familyTotalTable),
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
	return FamilyTotal{}, nil
}

func getFamilyTotal(childID string) FamilyTotal {
	ddb := db.GetDB()
	result, err := ddb.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(familyTotalTable),
		Key: map[string]*dynamodb.AttributeValue{
			"child_id": {
				S: aws.String(childID),
			},
		},
	})
	if err != nil {
		fmt.Println("Got error querying familyTotals")
		fmt.Println(err.Error())
	}

	familyTotal := FamilyTotal{}
	dynamodbattribute.UnmarshalMap(result.Item, &familyTotal)
	return familyTotal
}
