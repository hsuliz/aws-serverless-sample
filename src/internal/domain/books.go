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

func (b Books) FindAll(ctx context.Context) (types.BookRange, error) {
	return b.store.FindAll(ctx)
}
