package db

import (
	"errors"
	"fmt"
	"serverless-todo/model"

	"github.com/kataras/golog"
	uuid "github.com/satori/go.uuid"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const tableName = "serverless-todo"

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

func (itemRepo ItemRepository) Delete(id string) (deletedItem *model.Item, err error) {
	// if id is not a valid value, return an error
	if id == "" {
		err = errors.New("Id is empty")
		golog.Errorf("Could not produce a uuid: %v", err)
		return nil, err
	}

	// get a DynamoDB client
	svc := getDynamoDbClient()

	// Delete item in table TodoList
	tableName := "serverless-todo"

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(tableName),
	}

	result, err := svc.DeleteItem(input)

	if err != nil {
		golog.Errorf("Could not delete item: %v", err)
		return nil, err
	}

	if result.Attributes != nil {
		if result.Attributes["Id"] != nil {
			golog.Debug("Id: is NOT null")
			golog.Debugf("Id: %v", result.Attributes["Id"].String())
		} else {
			golog.Debug("Id: is null")
		}
	}

	//golog.Debugf("Title: %v", result.Attributes["Title"].String())
	//golog.Debugf("Description: %v", result.Attributes["Description"].String())

	/* deletedItem = &model.Item{
		Id:          result.Attributes["Id"].GoString(),
		Title:       result.Attributes["Title"].GoString(),
		Description: result.Attributes["Description"].GoString(),
	} */

	return
}

func (itemRepo ItemRepository) GetAll() ([]model.Item, error) {
	return nil, nil
}
