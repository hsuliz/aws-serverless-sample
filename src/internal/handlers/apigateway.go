package handlers

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"read-stats/internal/domain"
	"read-stats/internal/types"
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

	book, err := g.booksDomain.GetByID(ctx, bookID)
	if err != nil {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}
	return response(http.StatusOK, book), nil
}

func (g APIGatewayV2) Post(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	book := types.Book{}
	if err := json.Unmarshal([]byte(request.Body), &book); err != nil {
		return errResponse(http.StatusBadRequest, err.Error()), nil
	}

	if book.Title == "" {
		return errResponse(http.StatusBadRequest, "Missing book title"), nil
	}

	createdBook, err := g.booksDomain.Create(ctx, book)
	if err != nil {
		return errResponse(http.StatusBadRequest, err.Error()), nil
	}

	return response(http.StatusOK, createdBook), nil
}

func (g APIGatewayV2) Patch(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	bookID := request.PathParameters["id"]
	if bookID == "" || request.Body == "" {
		return errResponse(http.StatusBadRequest, ""), nil
	}

	var payload struct {
		PagesDone int `json:"pages_done"`
	}
	if err := json.Unmarshal([]byte(request.Body), &payload); err != nil {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}

	if err := g.booksDomain.UpdateBookPagesDone(ctx, bookID, payload.PagesDone); err != nil {
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}

	return response(http.StatusOK, ""), nil
}

func (g APIGatewayV2) Delete(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	bookID := request.PathParameters["id"]
	if bookID == "" {
		return errResponse(http.StatusBadRequest, "book id is empty"), nil
	}

	err := g.booksDomain.Delete(ctx, bookID)
	if err != nil {
		//#TODO add error type
		return errResponse(http.StatusInternalServerError, err.Error()), nil
	}
	return response(http.StatusOK, ""), nil
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
