package main

import (
	"encoding/json"
	"fmt"

	"github.com/andersonlira/stockids/db"
	"github.com/andersonlira/stockids/model"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// func Create(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	svc := GetDB()

// 	items := []model.Account{model.Account{ID: str.NewUUID(), Name: "Account 1"}}

// 	// Add each item to Movies table:
// 	for _, item := range items {
// 		av, err := dynamodbattribute.MarshalMap(item)

// 		if err != nil {
// 			fmt.Println("Got error marshalling map:")
// 			fmt.Println(err.Error())
// 			os.Exit(1)
// 		}

// 		// Create item in table Movies
// 		input := &dynamodb.PutItemInput{
// 			Item:      av,
// 			TableName: aws.String("accounts"),
// 		}

// 		_, err = svc.PutItem(input)

// 		if err != nil {
// 			fmt.Println("Got error calling PutItem:")
// 			fmt.Println(err.Error())
// 			os.Exit(1)
// 		}

// 		fmt.Println("Successfully added '", item.Name)
// 	}

// 	bk := &message{
// 		Title:   "Message",
// 		Message: "Showing message",
// 	}
// 	response, _ := json.Marshal(bk)
// 	return events.APIGatewayProxyResponse{Body: string(response), StatusCode: 201}, nil
// }

//GetChildren returns childrens from giving parent
func GetChildren(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	ddb := db.GetDB()
	result, err := ddb.Scan(&dynamodb.ScanInput{
		TableName: aws.String("skChild"),
	})
	if err != nil {
		fmt.Println("Got error calling PutItem:")
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
