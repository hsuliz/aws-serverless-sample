package types

import "context"

type Store interface {
	GetAll(context.Context, *string)
}
