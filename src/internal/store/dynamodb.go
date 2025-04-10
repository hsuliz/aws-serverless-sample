package store

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ddbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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
		log.Fatalf(err.Error())
	}
	client := dynamodb.NewFromConfig(cfg)
	return &DynamoDB{
		client:    client,
		tableName: tableName,
	}
}

func (d DynamoDB) FindBooks(ctx context.Context) (types.BookRange, error) {
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
	return bookRange, nil
}

func (d DynamoDB) GetBook(ctx context.Context, id string) (*types.Book, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]ddbtypes.AttributeValue{
			"id": &ddbtypes.AttributeValueMemberS{Value: id},
		},
		TableName: &d.tableName,
	}

	response, err := d.client.GetItem(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to get item from DynamoDB: %w", err)
	}

	if len(response.Item) == 0 {
		return nil, fmt.Errorf("book with id %s does not exist", id)
	}

	book := types.Book{}
	err = attributevalue.UnmarshalMap(response.Item, &book)
	if err != nil {
		return nil, fmt.Errorf("error getting item %w", err)
	}

	return &book, nil
}

func (d DynamoDB) CreateBook(ctx context.Context, book types.Book) error {
	item, err := attributevalue.MarshalMap(&book)
	if err != nil {
		return err
	}

	_, err = d.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &d.tableName,
		Item:      item,
	})
	if err != nil {
		return err
	}

	return nil
}

func (d DynamoDB) UpdateBookPagesDone(ctx context.Context, bookID string, pagesDone int) error {
	update := expression.Set(expression.Name("pages_done"), expression.Value(pagesDone))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return err
	}

	input := &dynamodb.UpdateItemInput{
		TableName: &d.tableName,
		Key: map[string]ddbtypes.AttributeValue{
			"id": &ddbtypes.AttributeValueMemberS{Value: bookID},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}

	_, err = d.client.UpdateItem(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (d DynamoDB) DeleteBook(ctx context.Context, bookID string) error {
	deleteInput := &dynamodb.DeleteItemInput{
		TableName: &d.tableName,
		Key: map[string]ddbtypes.AttributeValue{
			"id": &ddbtypes.AttributeValueMemberS{Value: bookID},
		},
	}

	_, err := d.client.DeleteItem(ctx, deleteInput)
	if err != nil {
		return err
	}

	return nil
}
