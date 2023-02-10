package handler

import (
	"github.com/clementbowe14/auxparty/tree/main/party/party"
)


type ErrorBody struct {
	Message *string `"json:error,omitempty"`
}

type EventHandler struct {
	partyService *party.PartyServiceProvider
}


func (e *EventHandler) handle(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	
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
