package restaurant

import (
	"context"
	"testing"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

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
		o := newOrder(id, "Saint pieter", 0)
		db := getDB(t)
		ctx := context.Background()
		r := NewRepo(db, nil)

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
		r := NewRepo(db, nil)

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
		r := NewRepo(db, nil)

		defer clearAndClose(db, t)
		r.CreateOrder(ctx, *o1)
		e := r.UpdateOrder(ctx, *o2)

		assert.Nil(t, e)
	})
}

func TestCreateCook(t *testing.T) {
	t.Run("success to create a new cooker", func(t *testing.T) {
		id := uuid.New()
		c := newCook(id, "Alfredo", 5)
		db := getDB(t)
		ctx := context.Background()
		r := NewRepo(db, nil)

		defer clearAndClose(db, t)
		e := r.CreateCook(ctx, *c)

		assert.Nil(t, e)
	})

	t.Run("failed when name is empty", func(t *testing.T) {
		id := uuid.New()
		c := newCook(id, "", 5)
		db := getDB(t)
		ctx := context.Background()
		r := NewRepo(db, nil)

		defer clearAndClose(db, t)
		e := r.CreateCook(ctx, *c)

		assert.NotEmpty(t, e)
	})

	t.Run("failed when score is 0", func(t *testing.T) {
		id := uuid.New()
		c := newCook(id, "Alfredo", 0)
		db := getDB(t)
		ctx := context.Background()
		r := NewRepo(db, nil)

		defer clearAndClose(db, t)
		e := r.CreateCook(ctx, *c)

		assert.NotEmpty(t, e)
	})
}
