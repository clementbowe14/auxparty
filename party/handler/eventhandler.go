package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/clementbowe14/auxparty/db"
	"github.com/clementbowe14/auxparty/service/party"
)

type ErrorBody struct {
	Message *string `"json:error,omitempty"`
}

type EventHandler struct {
	partyService party.PartyServiceProvider
}

func NewEventHandler(db db.DynamoClient) EventHandler {
	var h EventHandler
	h.partyService = party.PartyServiceProvider{
		DatabaseClient: db,
	}

	return h
}

func (e *EventHandler) handle(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	partyName := req.QueryStringParameters["party_name"]
	description := req.QueryStringParameters["description"]
	partyCreator := req.QueryStringParameters["party_creator"]

	party, err := e.partyService.CreateParty(partyCreator, description, partyName)

	if err != nil {
		resp := ErrorBody{
			Message: aws.String(err.Error()),
		}

		return e.apiResponse(http.StatusBadRequest, resp)
	}

	return e.apiResponse(http.StatusOK, party)
}

func (e *EventHandler) apiResponse(httpStatus int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	resp.StatusCode = httpStatus
	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)

	return &resp, nil
}
