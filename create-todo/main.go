package main

import (
	"serverless-todo/model"
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

func getDynamoDbClient() (svc *dynamodb.DynamoDB) {
	// Initialize a session that the SDK will use to load
	sess := session.Must(session.NewSession())

	// Create DynamoDB client
	svc = dynamodb.New(sess)

	return
}

func saveTodo(item *model.Item) (resp Response, returnError error) {
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

	var buf bytes.Buffer

	body, returnError := json.Marshal(av)
	if returnError != nil {
		return Response{StatusCode: 404}, returnError
	}
	json.HTMLEscape(&buf, body)

	resp = Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
			//"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (resp Response, returnError error) {

	

	return saveTodo(&model.Item{
		Title:       "The Big New Movie",
		Description: "Nothing happens at all.",
	})
}

func main() {
	lambda.Start(Handler)
}
