package sqlite_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

	"github.com/luanpontolio/restaurant-kitchen-go/domain"
	sqliteOrderRepo "github.com/luanpontolio/restaurant-kitchen-go/order/repository/sqlite"
)

func newOrder(id uuid.UUID, plate string, score int64) *domain.Order {
	return &domain.Order{
		ID:    id,
		Plate: plate,
		Score: score,
	}
}

func getDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", "../../../db/restaurant_test.db")
	assert.Nil(t, err)

	return db
}

func clearAndClose(db *sql.DB, t *testing.T) {
	tx, err := db.Begin()
	assert.Nil(t, err)

	_, err = tx.Exec("delete from orders")
	assert.Nil(t, err)

	if err != nil {
		tx.Rollback()
	}

	tx.Commit()
	db.Close()
}

func TestCreateOrder(t *testing.T) {
	t.Run("success to create a new order", func(t *testing.T) {
		id := uuid.New()
		o := newOrder(id, "Parmegiana", 5)
		db := getDB(t)
		ctx := context.Background()
		r := sqliteOrderRepo.NewRepo(db, nil)

		defer clearAndClose(db, t)
		e := r.CreateOrder(ctx, *o)

		assert.Nil(t, e)
	})

	t.Run("failed when plate is empty", func(t *testing.T) {
		id := uuid.New()
		o := newOrder(id, "", 5)
		db := getDB(t)
		ctx := context.Background()
		r := sqliteOrderRepo.NewRepo(db, nil)

		defer clearAndClose(db, t)
		e := r.CreateOrder(ctx, *o)

		assert.NotEmpty(t, e)
	})

	t.Run("failed when score is 0", func(t *testing.T) {
		id := uuid.New()
		o := newOrder(id, "Saint pieter", 0)
		db := getDB(t)
		ctx := context.Background()
		r := sqliteOrderRepo.NewRepo(db, nil)

		defer clearAndClose(db, t)
		e := r.CreateOrder(ctx, *o)

		assert.NotEmpty(t, e)
	})
}

func TestUpdateOrder(t *testing.T) {
	t.Run("not execute when params are empty", func(t *testing.T) {
		o := newOrder(uuid.Nil, "", 0)
		db := getDB(t)
		ctx := context.Background()
		r := sqliteOrderRepo.NewRepo(db, nil)

		defer clearAndClose(db, t)
		e := r.UpdateOrder(ctx, *o)

		assert.Nil(t, e)
	})

	t.Run("success update a onder", func(t *testing.T) {
		uid := uuid.New()
		o1 := newOrder(uid, "Parmegiana", 5)
		o2 := newOrder(uid, "Parmegiana de Frango", 10)
		db := getDB(t)
		ctx := context.Background()
		r := sqliteOrderRepo.NewRepo(db, nil)

		defer clearAndClose(db, t)
		r.CreateOrder(ctx, *o1)
		e := r.UpdateOrder(ctx, *o2)

		assert.Nil(t, e)
	})
}

func TestGetOrder(t *testing.T) {

	t.Run("failed when id is empty", func(t *testing.T) {
		db := getDB(t)
		ctx := context.Background()
		r := sqliteOrderRepo.NewRepo(db, nil)

		defer clearAndClose(db, t)
		_, e := r.GetOrder(ctx, "")

		assert.Contains(t, e.Error(), "Unable to handle")
	})

	t.Run("failed when id is invalid", func(t *testing.T) {
		db := getDB(t)
		ctx := context.Background()
		r := sqliteOrderRepo.NewRepo(db, nil)

		defer clearAndClose(db, t)
		_, e := r.GetOrder(ctx, "1234")

		assert.Contains(t, e.Error(), "Unable to handle")
	})

	t.Run("success get order", func(t *testing.T) {
		id := uuid.New()
		o := newOrder(id, "Saint pieter", 5)
		db := getDB(t)
		ctx := context.Background()
		r := sqliteOrderRepo.NewRepo(db, nil)

		defer clearAndClose(db, t)
		r.CreateOrder(ctx, *o)
		result, _ := r.GetOrder(ctx, id.String())

		assert.Contains(t, result.Plate, "Saint pieter")
		assert.Contains(t, result.State.String(), "esperando")
	})
}
