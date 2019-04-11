package db

import (
	"fmt"
	"serverless-todo/model"

	"github.com/kataras/golog"
	uuid "github.com/satori/go.uuid"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func getDynamoDbClient() (svc *dynamodb.DynamoDB) {
	// Initialize a session that the SDK will use to load
	sess := session.Must(session.NewSession())

	// Create DynamoDB client
	svc = dynamodb.New(sess)

	return
}

type ItemRepository struct{}

func (itemRepo ItemRepository) Close() {}

func (itemRepo ItemRepository) Insert(item *model.Item) (result *model.Item, returnError error) {

	// if they are inserting a new item, then we need a new id
	if item.Id == "" {
		uuid, returnError := uuid.NewV4()
		if returnError != nil {
			golog.Errorf("Could not produce a uuid: %v", returnError)
			return nil, returnError
		}
		item.Id = uuid.String()
	}

	// get a DynamoDB client
	svc := getDynamoDbClient()

	av, returnError := dynamodbattribute.MarshalMap(item)
	if returnError != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(returnError.Error())
		return
	}

	// Create item in table TodoList
	tableName := "serverless-todo"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, returnError = svc.PutItem(input)
	if returnError != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(returnError.Error())
		return
	}

	result = item

	return
}

func (itemRepo ItemRepository) Delete(id int) error {
	return nil
}

func (itemRepo ItemRepository) GetAll() ([]model.Item, error) {
	return nil, nil
}
