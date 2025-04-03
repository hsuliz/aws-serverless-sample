package handlers

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log"
	"net/http"
	"read-stats/internal/domain"
)

type APIGatewayV2 struct {
	booksDomain *domain.Books
}

func NewAPIGatewayV2(books *domain.Books) *APIGatewayV2 {
	return &APIGatewayV2{booksDomain: books}
}

func (g APIGatewayV2) Find(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	books, err := g.booksDomain.Find(ctx)
	if err != nil {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}

	return response(http.StatusOK, books), nil
}

func (g APIGatewayV2) Get(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	bookID, _ := request.PathParameters["id"]
	log.Println(bookID)

	book, err := g.booksDomain.GetByID(ctx, bookID)
	if err != nil {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}
	return response(http.StatusOK, book), nil
}

func response(code int, object interface{}) events.APIGatewayV2HTTPResponse {
	marshalled, err := json.Marshal(object)
	if err != nil {
		return errResponse(http.StatusInternalServerError, err.Error())
	}

	return events.APIGatewayV2HTTPResponse{
		StatusCode: code,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body:            string(marshalled),
		IsBase64Encoded: false,
	}
}

func errResponse(status int, body string) events.APIGatewayV2HTTPResponse {
	message := map[string]string{
		"message": body,
	}

	messageBytes, _ := json.Marshal(&message)

	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(messageBytes),
	}
}
