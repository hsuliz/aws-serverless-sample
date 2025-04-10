package domain

import (
	"context"
	"fmt"
	"read-stats/internal/types"
	"time"
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

func (b Books) Create(ctx context.Context, book types.Book) (*types.Book, error) {
	book.ID = b.generateID()

	err := b.store.CreateBook(ctx, book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (b Books) UpdateBookPagesDone(ctx context.Context, bookID string, pagesDone int) error {
	return b.store.UpdateBookPagesDone(ctx, bookID, pagesDone)
}

func (b Books) generateID() string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	return fmt.Sprintf("%d", timestamp)
}

func isDone(book types.Book) bool {
	return book.PagesDone == book.Pages
}
