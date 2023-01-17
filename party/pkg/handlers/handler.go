package handlers

import (
	"net/http"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
)



type HelloWorld struct {
	message *string `json:message, omitempty`
}


func CreateParty() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusOK, HelloWorld{
		aws.String("Hello World"),		
	})
}

