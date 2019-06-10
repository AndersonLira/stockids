package db

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//GetDB returns a dynamo db connection
func GetDB() *dynamodb.DynamoDB {
	env := os.Getenv("SK_ENVIRONMENT")

	switch env {
	case "DEV":
		return getDevDB()
	default:
		return getProdDB()
	}
}

func getDevDB() *dynamodb.DynamoDB {
	log.Println("Using local database")
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://localhost:8000")},
	)

	if err != nil {
		fmt.Println("Error creating session:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	return svc

}

func getProdDB() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		fmt.Println("Error creating session:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	return svc

}
