package model

import (
	"errors"
	"google/uuid"
	"time"
	"github.com/clementbowe14/auxparty/tree/main/db"
)

var (
	EmptyPartyNameError = "Cannot create a party using an empty string"
)

type PartyServiceProvider struct {
	DatabaseClient db.DynamoClient[Party]
}

type DefaultParty interface {
	func(partyCreator string, description string, partyName string) (*Party, error)
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

func (p *PartyServiceProvider) CreateParty(partyCreator string, description string, partyName string) (*Party, error) {
	now := time.Now()

	if len(partyName) == 0 {
		return nil, errors.New(EmptyPartyNameError)
	}

	party := Party{
		PartyId:      uuid.NewString(),
		PartyName:    partyName,
		PartyCreator: partyCreator,
		DateCreated:  now.Unix(),
		IsActive:     true,
		Description:  description,
	}

	err := p.DatabaseClient.InsertInto(party)

	if err != nil {
		return nil, err
	}

	return &party, nil
}
