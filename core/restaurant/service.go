package restaurant

import "context"

type OrderService interface {
	GetAllOrder(ctx context.Context, state int64, delivery_at bool) ([]*Order, error)
	CreateOrder(ctx context.Context, plate string, amount int64) (string, string, error)
	UpdateOrder(ctx context.Context, id string, plate string, amount int64, state OrderState) (string, string, error)
	UpdateOrderHash(ctx context.Context, id string) error
}

type CookService interface {
	GetCookByScore(ctx context.Context, score int64) (*Cook, error)
	CreateCook(ctx context.Context, name string, score int64) (string, string, error)
	UpdateCook(ctx context.Context, id string, score int64, state int64) (string, string, error)
}

type RestaurantService interface {
	OrderService
	CookService
}
