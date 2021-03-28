package order

import (
	"context"

	"github.com/google/uuid"
)

// table order
// CREATE TABLE orders(
//   ID string not null primary key,
//   Plate text not null,
//   Score integer not null,
//   Hash text
// )

type Order struct {
	ID    uuid.UUID  `json:"uuid"`
	Plate string     `json:"plate"`
	Score int64      `json:"amount"`
	State OrderState `json:"state:omitempty"`
	Hash  string     `json:"hash:omitempty"`
}

type OrderState int

const (
	STATEWAIT     = 0
	STATEPREPARE  = 1
	STATEDELIVERY = 2
)

func (s OrderState) String() string {
	switch s {
	case STATEWAIT:
		return "esperando"
	case STATEPREPARE:
		return "preparando"
	case STATEDELIVERY:
		return "entregue"
	default:
		return "unknow"
	}
}

type Repository interface {
	CreateOrder(ctx context.Context, order Order) error
	GetOrder(ctx context.Context, id string) (*Order, error)
}
