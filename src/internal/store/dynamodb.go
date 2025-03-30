package store

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type DynamoDB struct {
	client    *dynamodb.Client
	tableName string
}
