package sqlite

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/luanpontolio/restaurant-kitchen-go/domain"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type sqliteOrderRepo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) domain.OrderRepository {
	return &sqliteOrderRepo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *sqliteOrderRepo) CreateOrder(ctx context.Context, data domain.Order) error {
	sql := `
		INSERT INTO orders (id, plate, score)
		VALUES ($1, $2, $3)`

	if data.Plate == "" || data.Score == 0 {
		return RepoErr
	}

	_, err := repo.db.ExecContext(ctx, sql, data.ID, data.Plate, data.Score)
	if err != nil {
		return err
	}
	return nil
}

func (repo *sqliteOrderRepo) UpdateOrder(ctx context.Context, data domain.Order) error {
	sql := `UPDATE orders SET plate=$1, score=$2 where id=$3 and state=$4`

	if data.Plate == "" || data.Score == 0 {
		return nil
	}

	_, err := repo.db.ExecContext(ctx, sql, data.Plate, data.Score, data.ID, 0)
	if err != nil {
		return err
	}
	return nil
}

func (repo *sqliteOrderRepo) GetOrder(ctx context.Context, id string) (*domain.Order, error) {
	var order domain.Order
	err := repo.db.QueryRow("SELECT id, plate, score, hash FROM orders WHERE id=$1", id).Scan(
		&order.ID, &order.Plate, &order.Score, &order.Hash,
	)

	if err != nil {
		return &order, RepoErr
	}
	return &order, nil
}
