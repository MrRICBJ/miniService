package repositories

import (
	"context"
	"miniService/internal/entities"
)

type Repo interface {
	GetById(ctx context.Context, IDs []int) (*[]entities.Item, error)
}
