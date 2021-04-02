package restaurant

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

func NewRepo(db *sql.DB, logger log.Logger) RestaurantRespository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateOrder(ctx context.Context, order Order) error {
	sql := `
		INSERT INTO orders (id, plate, score, state)
		VALUES ($1, $2, $3, $4)`

	if order.Plate == "" || order.Score == 0 {
		return RepoErr
	}

	_, err := repo.db.ExecContext(ctx, sql, order.ID, order.Plate, order.Score, 0)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) UpdateOrder(ctx context.Context, order Order) error {
	sql := `UPDATE orders SET plate=$1, score=$2 where id=$3 and state=$4`

	if order.Plate == "" || order.Score == 0 {
		return nil
	}

	_, err := repo.db.ExecContext(ctx, sql, order.Plate, order.Score, order.ID, 0)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) CreateCook(ctx context.Context, cook Cook) error {
	sql := `
		INSERT INTO cooks (id, name, score)
		VALUES ($1, $2, $3)`

	if cook.Name == "" || cook.Score == 0 {
		return RepoErr
	}

	_, err := repo.db.ExecContext(ctx, sql, cook.ID, cook.Name, cook.Score)
	if err != nil {
		return err
	}
	return nil
}
