package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type HelloWorld struct {
	Message string `json:"messsage"`
}
/**

- request body:
	party_name string
	party_creator string
	description string

- createParty
	- Party{
		PartyName: party_name
		PartyCreator: party_creator
		description: descript
	}
**/

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return apiResponse()
}

func apiResponse() (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	msg := HelloWorld {
		Message: "Hello Clement! Welcome to APIGateway.",
	}

	resp.StatusCode = 200
	stringBody, _ := json.Marshal(msg)
	resp.Body = string(stringBody)

	return &resp, nil
}
