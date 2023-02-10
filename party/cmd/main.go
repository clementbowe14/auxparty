package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/clementbowe14/auxparty/tree/main/party/party"
	"github.com/clementbowe14/auxparty/tree/main/party/handler"
	"github.com/clementbowe14/auxparty/tree/main/party/db"

)

var (
	partyTableName       = "party"
	partyMemberTableName = "partyMember"
)

//add logging to this
type EventHandler struct {
	partyService *party.PartyServiceProvider
}

func main() {

	region := os.Getenv("AWS_REGION")

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		return
	}

	dynamoClient := dynamodb.New(awsSession)

	e := handler.EventHandler{
		partyService: &party.PartyServiceProvider{
			DatabaseClient: db.DynamoClient[Party]{
				tableName: partyTableName,
				client:    dynamoClient,
			},
		},
	}

	lambda.Start(e.handle)
}
