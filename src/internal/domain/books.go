package domain

import "read-stats/internal/store"

type Books struct {
	store store.DynamoDB
}
