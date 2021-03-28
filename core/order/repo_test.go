package order

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func newOrder(id uuid.UUID, plate string, score int64) *Order {
	return &Order{
		ID:    id,
		Plate: plate,
		Score: score,
	}
}

func getDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", "../../data/restaurant_test.db")
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
		r := NewRepo(db, nil)

		defer clearAndClose(db, t)
		e := r.CreateOrder(ctx, *o)

		assert.Nil(t, e)
	})

	t.Run("failed when plate is empty", func(t *testing.T) {
		id := uuid.New()
		o := newOrder(id, "", 5)
		db := getDB(t)
		ctx := context.Background()
		r := NewRepo(db, nil)

		defer clearAndClose(db, t)
		e := r.CreateOrder(ctx, *o)

		assert.NotEmpty(t, e)
	})

	t.Run("failed when score is 0", func(t *testing.T) {
		id := uuid.New()
		o := newOrder(id, "Sant pieter", 0)
		db := getDB(t)
		ctx := context.Background()
		r := NewRepo(db, nil)

		defer clearAndClose(db, t)
		e := r.CreateOrder(ctx, *o)

		assert.NotEmpty(t, e)
	})
}
