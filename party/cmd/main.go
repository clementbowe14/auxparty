package main

import (
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/clementbowe14/auxparty/db"
	"github.com/clementbowe14/auxparty/party/handler"
	"github.com/clementbowe14/auxparty/service/party"
)

var (
	partyTableName       = "party"
	partyMemberTableName = "partyMember"
)

func main() {

	region := os.Getenv("AWS_REGION")

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		return
	}

	dynamoClient := db.NewDynamoClient[party.Party](dynamodb.New(awsSession), partyTableName)

	e := handler.NewEventHandler(dynamoClient)

	lambda.Start(e.handle)
}
