package types

import "context"

type Store interface {
	FindBooks(context.Context) (BookRange, error)
	GetBook(context.Context, string) (*Book, error)
}
