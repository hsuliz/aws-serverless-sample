package internal

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

func Handler(_ context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println(event)

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "kys",
	}

	return response, nil
}
