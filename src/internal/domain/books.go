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
	return b.store.FindBooks(ctx)
}

func (b Books) GetByID(ctx context.Context, id string) (*types.Book, error) {
	return b.store.GetBook(ctx, id)
}
