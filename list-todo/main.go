package main

import (
	"encoding/json"
	"fmt"
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

func listTodo() (resp Response, returnError error) {
	itemRepo := db.ItemRepository{}

	savedItems, returnError := itemRepo.GetAll()
	if returnError != nil {
		fmt.Println("Got error saving the item:")
		fmt.Println(returnError.Error())
		return
	}

	body, returnError := json.Marshal(savedItems)
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

	golog.Debugf("req: %v", req)

	return listTodo()
}

func main() {
	lambda.Start(Handler)
}
