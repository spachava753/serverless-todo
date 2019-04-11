package main

import (
	//"bytes"
	"context"
	"encoding/json"
	"fmt"
	"serverless-todo/db"
	"serverless-todo/model"

	"github.com/kataras/golog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

func saveTodo(item *model.Item) (resp Response, returnError error) {
	itemRepo := db.ItemRepository{}

	savedItem, returnError := itemRepo.Insert(item)
	if returnError != nil {
		fmt.Println("Got error saving the item:")
		fmt.Println(returnError.Error())
		return
	}

	//var buf bytes.Buffer

	body, returnError := json.Marshal(savedItem)
	if returnError != nil {
		return Response{StatusCode: 404}, returnError
	}
	//json.Indent(&buf, body, "", "\t")

	resp = Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		//Body:            buf.String(),
		Body:            string(body),
		Headers: map[string]string{
			"Content-Type": "application/json",
			//"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(req events.APIGatewayProxyRequest) (resp Response, returnError error) {

	golog.SetLevel("debug")

	golog.Debugf("Valid json: %v", json.Valid(req.Body))

	requestItem := model.Item{}

	json.Unmarshal(req.Body, requestItem)

	return saveTodo(&requestItem)
}

func main() {
	lambda.Start(Handler)
}
