package order

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateOrder(ctx context.Context, order Order) error {
	sql := `
		INSERT INTO orders (id, plate, score)
		VALUES ($1, $2, $3)`

	if order.Plate == "" || order.Score == 0 {
		return RepoErr
	}

	_, err := repo.db.ExecContext(ctx, sql, order.ID, order.Plate, order.Score)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetOrder(ctx context.Context, id string) (*Order, error) {
	var order Order
	err := repo.db.QueryRow("SELECT * FROM orders WHERE id=$1", id).Scan(
		&order.ID, &order.Plate, &order.Score, &order.Hash,
	)

	if err != nil {
		return &order, RepoErr
	}

	return &order, nil
}
