package main

import (
	"encoding/json"
	"serverless-todo/db"

	"github.com/kataras/golog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

func deleteTodo(id string) (resp Response, returnError error) {
	itemRepo := db.ItemRepository{}

	deletedItem, returnError := itemRepo.Delete(id)
	if returnError != nil {
		golog.Error("Got error deleting the item:")
		golog.Error(returnError.Error())
		return
	}

	body, returnError := json.Marshal(deletedItem)
	if returnError != nil {
		return Response{StatusCode: 404}, returnError
	}

	resp = Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Access-Control-Allow-Origin": "*",
		},
	}

	return
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(req events.APIGatewayProxyRequest) (resp Response, returnError error) {

	golog.SetLevel("debug")

	body := req.PathParameters["id"]

	golog.Debugf("req: %v", req)
	golog.Debugf("input: %v", req.PathParameters["id"])

	return deleteTodo(body)
}

func main() {
	lambda.Start(Handler)
}
