package order

import "context"

type Service interface {
	CreateOrder(ctx context.Context, plate string, amount int64) (string, error)
	GetOrder(ctx context.Context, id string) (*Order, error)
}
