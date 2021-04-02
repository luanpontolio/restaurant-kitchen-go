package restaurant

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// table order
// CREATE TABLE orders(ID string not null primary key, Plate text not null, Score integer not null, State integer, Hash text, CreatedAt timestamp  NOT NULL DEFAULT current_timestamp, UpdatedAt timestamp  NOT NULL DEFAULT current_timestamp);
// CREATE TABLE cooks(ID string not null primary key, Name text not null, Score integer not null, CreatedAt timestamp  NOT NULL DEFAULT current_timestamp, UpdatedAt timestamp  NOT NULL DEFAULT current_timestamp);
type Order struct {
	ID        uuid.UUID  `json:"uuid"`
	Plate     string     `json:"plate"`
	Score     int64      `json:"amount"`
	State     OrderState `json:"state:omitempty"`
	Hash      string     `json:"hash:omitempty"`
	CreatedAt time.Time  `json:"created_at:omitempty"`
	UpdatedAt time.Time  `json:"updated_at:omitempty"`
}

type Cook struct {
	ID        uuid.UUID `json:"uuid"`
	Name      string    `json:"plate"`
	Score     int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at:omitempty"`
	UpdatedAt time.Time `json:"updated_at:omitempty"`
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

type OrderRepository interface {
	CreateOrder(ctx context.Context, order Order) error
	UpdateOrder(ctx context.Context, order Order) error
}

type CookRepository interface {
	CreateCook(ctx context.Context, cook Cook) error
}

type RestaurantRespository interface {
	OrderRepository
	CookRepository
}
