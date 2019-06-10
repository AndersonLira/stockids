package dbsetup

import (
	"fmt"
	"os"

	"github.com/andersonlira/stockids/db"
	"github.com/andersonlira/stockids/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var tables = []model.Tableable{
	&model.Family{},
}

func CreateTables() {
	for index, _ := range tables {
		t := tables[index]
		props := GetModelProps(t)
		fmt.Println(t.GetTableName())
		for _, p := range props {
			fmt.Printf("   %v\n", p.FieldName)
		}
	}

	//createFamilyTable()
}

func createFamilyTable() {
	// Create table Movies
	tableName := "skFamily"

	svc := db.GetDB()

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("avatar"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Year"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Title"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Created the table", tableName)

}
