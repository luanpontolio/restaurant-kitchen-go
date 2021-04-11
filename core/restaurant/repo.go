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

func (repo *repo) GetAllOrder(ctx context.Context, state int64, delivery_at bool) ([]*Order, error) {
	var result []*Order
	var sort string

	if delivery_at {
		sort += " order by deliveryat"
	} else {
		sort += " order by createdat"
	}

	sql := "select * from orders where state=$1 " + sort
	rows, err := repo.db.QueryContext(ctx, sql, state)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var o Order
		err = rows.Scan(&o.ID, &o.Plate, &o.Score, &o.State, &o.Hash, &o.DeliveryAt, &o.CreatedAt, &o.UpdatedAt)
		if err != nil {
			return nil, err
		}
		result = append(result, &o)
	}
	return result, nil
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
	sql := `UPDATE orders SET plate=$1, score=$2, state=$3 where id=$4`

	if order.Plate == "" || order.Score == 0 {
		return nil
	}

	_, err := repo.db.ExecContext(ctx, sql, order.Plate, order.Score, order.State, order.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) UpdateOrderHash(ctx context.Context, order Order) error {
	sql := `UPDATE orders SET hash=$1 where id=$2 and state=2`

	_, err := repo.db.ExecContext(ctx, sql, order.Hash, order.ID)
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

func (repo *repo) UpdateCook(ctx context.Context, cook Cook) error {
	sql := `UPDATE cooks SET score=$1, state=$2 where id=$3`

	if cook.Score == 0 {
		return nil
	}

	_, err := repo.db.ExecContext(ctx, sql, cook.Score, cook.State, cook.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetCookByScore(ctx context.Context, score int64) (*Cook, error) {
	var c Cook

	sql := "select id, score from cooks where state=0 and score >= $1 limit 1"
	err := repo.db.QueryRowContext(ctx, sql, score).Scan(&c.ID, &c.Score)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
