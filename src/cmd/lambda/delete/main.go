package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
	"read-stats/internal/domain"
	"read-stats/internal/handlers"
	"read-stats/internal/store"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	tableName, exists := os.LookupEnv("TABLE_NAME")
	if !exists {
		panic("Need TABLE_NAME environment variable")
	}

	dynamoDB := store.NewDynamoDB(context.TODO(), tableName)
	booksDomain := domain.NewBooks(dynamoDB)
	handler := handlers.NewAPIGatewayV2(booksDomain)
	lambda.Start(handler.Delete)
}
