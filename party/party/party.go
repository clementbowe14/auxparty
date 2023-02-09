package model

import (
	"errors"
	"google/uuid"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	EmptyPartyNameError = "Cannot create a party using an empty string"
	FailedToMarshalItem = "Failed to marshal item"
	ErrorCouldNotPutItem = "Error, failed to enter new item"
	tableName = "party"
)

type DefaultParty interface {
	func (partyCreator string, description string, partyName string) (*Party, error)
}

type Party struct {
	PartyId                 string `json:"party_id"`
	PartyName               string `json:"party_name"`
	PartyCreator            string `json:"user_id"`
	DateCreated             int64  `json:"date_created"`
	Description             string `json:"description"`
	IsActive                bool   `json:"is_active"`
	TotalMembers            int64  `json:"total_members"`
	TotalMusicListeningTime int64  `json:"total_music_listening_time`
}

func CreateParty(partyCreator string, description string, partyName string) (*Party, error) {
	now := time.Now()

	if len(partyName) == 0 {
		return nil, errors.New(EmptyPartyNameError)
	}

	p := Party{
		PartyId:      uuid.NewString(),
		PartyName:    partyName,
		PartyCreator: partyCreator,
		DateCreated:  now.Unix(),
		IsActive:     true,
		Description:  description,
	}

	attributeValues, err := dynamodbattribute.MarshalMap(p)

	if err != nil {
		return nil, errors.New(FailedToMarshalItem)
	}

	input := &dynamodb.PutItemInput{
		Item:      attributeValues,
		TableName: aws.String(tableName),
	}

	_, err = dynamoClient.putItem(input)

	if err != nil {
		return nil, errors.New(ErrorCouldNotPutItem)
	}

	return &p, nil

}

/**
db
item Party
insertInto(item interface{}, tableName) (error) {
	reflectItem -> Party
	reflectItem -> attributeValues 
	attributeValues -> input
	input -> putItemInput

}


**
/
