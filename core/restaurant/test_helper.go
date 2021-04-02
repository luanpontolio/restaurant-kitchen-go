package restaurant

import (
	"database/sql"
	"testing"

	"github.com/google/uuid"
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
