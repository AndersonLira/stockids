package dbsetup

import (
	"log"
	"os"

	"github.com/andersonlira/stockids/db"

	"github.com/andersonlira/stockids/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var tables = []model.Tableable{
	&model.Family{},
}

//CreateTables of application
func CreateTables() {

	for _, t := range tables {
		props := GetModelProps(t)
		tableName := t.GetTableName()

		log.Printf("Creating table %s", tableName)

		attributeDefinitions := []*dynamodb.AttributeDefinition{}
		keySchema := []*dynamodb.KeySchemaElement{}

		for _, p := range props {

			if p.FieldIndex {
				field := dynamodb.AttributeDefinition{
					AttributeName: aws.String(p.FieldName),
					AttributeType: aws.String(p.FieldType),
				}

				attributeDefinitions = append(attributeDefinitions, &field)
				key := dynamodb.KeySchemaElement{
					AttributeName: aws.String(p.FieldName),
					KeyType:       aws.String(p.FieldKeyType),
				}
				keySchema = append(keySchema, &key)
				log.Printf("%s --> %v", tableName, field.AttributeName)
			}

		}

		input := &dynamodb.CreateTableInput{
			AttributeDefinitions: attributeDefinitions,
			KeySchema:            keySchema,
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(1),
				WriteCapacityUnits: aws.Int64(1),
			},
			TableName: aws.String(tableName),
		}

		svc := db.GetDB()
		_, err := svc.CreateTable(input)
		if err != nil {
			log.Printf("Got error calling CreateTable: %s", tableName)
			log.Println(err.Error())
			os.Exit(1)
		}
		log.Printf("Table %s created", tableName)
	}
}
