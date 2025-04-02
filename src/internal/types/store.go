package types

import "context"

type Store interface {
	FindAll(context.Context) (BookRange, error)
}
