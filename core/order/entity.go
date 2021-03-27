package order

import (
	"context"

	"github.com/google/uuid"
)

type Order struct {
	ID     uuid.UUID  `json:"uuid"`
	PLATE  string     `json:"plate"`
	AMOUNT int64      `json:"amount"`
	STATE  OrderState `json:"state:omitempty"`
	HASH   string     `json:"hash:omitempty"`
}

type OrderState int

const (
	STATEWAIT     = 0
	STATEPREPARE  = 1
	STATEDELIVERY = 2
)

func (t OrderState) String() string {
	switch t {
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
	GetOrder(ctx context.Context, uuid string) (string, error)
}
