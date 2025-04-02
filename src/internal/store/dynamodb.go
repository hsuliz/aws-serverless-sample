package store

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"
	"read-stats/internal/types"
)

type DynamoDB struct {
	client    *dynamodb.Client
	tableName string
}

func NewDynamoDB(ctx context.Context, tableName string) *DynamoDB {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	client := dynamodb.NewFromConfig(cfg)
	return &DynamoDB{
		client:    client,
		tableName: tableName,
	}
}

func (d DynamoDB) FindAll(ctx context.Context) (types.BookRange, error) {
	bookRange := types.BookRange{
		Books: []types.Book{},
	}

	input := &dynamodb.ScanInput{
		TableName: &d.tableName,
		Limit:     aws.Int32(10),
	}

	result, err := d.client.Scan(ctx, input)

	if err != nil {
		return bookRange, fmt.Errorf("failed to get items from DynamoDB: %w", err)
	}
	err = attributevalue.UnmarshalListOfMaps(result.Items, &bookRange.Books)
	log.Println(bookRange)
	return bookRange, nil
}
