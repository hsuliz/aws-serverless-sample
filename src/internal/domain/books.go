package domain

import (
	"context"
	"read-stats/internal/types"
)

type Books struct {
	store types.Store
}

func NewBooks(store types.Store) *Books {
	return &Books{store: store}
}

func (b Books) Find(ctx context.Context) (types.BookRange, error) {
	books, err := b.store.FindBooks(ctx)
	return books, err
}

func (b Books) GetByID(ctx context.Context, id string) (*types.Book, error) {
	book, err := b.store.GetBook(ctx, id)
	if err != nil {
		return nil, err
	}

	if book == nil {
		return nil, nil
	}

	book.BookDone = isDone(*book)
	return book, err
}

func isDone(book types.Book) bool {
	return book.PagesDone == book.Pages
}
