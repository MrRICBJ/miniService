package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"miniService/internal/db"
	"miniService/internal/entities"
)

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repo {
	return &repo{db: db}
}

func (r *repo) GetById(ctx context.Context, IDs []int) (*[]entities.Item, error) {
	items := make([]entities.Item, 0)

	for _, id := range IDs {
		item := entities.Item{}
		err := r.db.GetContext(ctx, &item, db.SelectById, id)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return &items, nil
}
