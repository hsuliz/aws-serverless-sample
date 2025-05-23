package types

import "context"

type Store interface {
	FindBooks(context.Context) (BookRange, error)
	GetBook(context.Context, string) (*Book, error)
	CreateBook(context.Context, Book) error
	UpdateBookPagesDone(context.Context, string, int) error
	DeleteBook(context.Context, string) error
}
