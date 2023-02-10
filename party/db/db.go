package db

import (
	"errors"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	FailedToMarshalItem  = "Failed to marshal item"
	ErrorCouldNotPutItem = "Error, failed to enter new item"
)

/**
* dynamo client provider uses mutex to ensure safe concurrent operations
- read mutex ensures that no other process writes to the the same table name
- write mutex ensures that no other process
**/

type DynamoClient[T interface{}] struct {
	tableName string
	client    dynamodbiface.DynamoDBAPI
	mu sync.Mutex
	errChan chan error
}

/**
* How do I customize the construction this type parameter list
to use the generic values that I create?
**/
func (d *DynamoClient[T])InsertInto(item T) error {

	attributeValues, err := dynamodbattribute.MarshalMap(item)

	if err != nil {
		return errors.New(FailedToMarshalItem)
	}
	go func() {
		defer d.mu.Unlock()
	d.mu.Lock()
		input := &d.client.PutItemInput{
			Item:      attributeValues,
			TableName: aws.String(d.tableName),
		}

		d.errChan <- d.client.putItem(input)
	}()

	err <- d.errChan

	if err != nil {
		return errors.New(ErrorCouldNotPutItem)
	}

	return nil

}
