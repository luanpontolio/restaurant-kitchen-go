package restaurant

import "context"

type OrderService interface {
	CreateOrder(ctx context.Context, plate string, amount int64) (string, string, error)
	UpdateOrder(ctx context.Context, id string, plate string, amount int64) (string, string, error)
}

type CookService interface {
	CreateCook(ctx context.Context, name string, score int64) (string, string, error)
	UpdateCook(ctx context.Context, id string, score int64) (string, string, error)
}

type RestaurantService interface {
	OrderService
	CookService
}
