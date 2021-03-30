package order

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/google/uuid"
	"github.com/luanpontolio/restaurant-kitchen-go/domain"
)

type orderService struct {
	repostory domain.OrderRepository
	logger    log.Logger
}

func NewService(rep domain.OrderRepository, logger log.Logger) domain.OrderService {
	return &orderService{
		repostory: rep,
		logger:    logger,
	}
}

func (s orderService) CreateOrder(ctx context.Context, plate string, score int64) (string, string, error) {
	logger := log.With(s.logger, "method", "CreateOrder")

	id := uuid.New()
	order := domain.Order{
		ID:    id,
		Plate: plate,
		Score: score,
	}

	if err := s.repostory.CreateOrder(ctx, order); err != nil {
		level.Error(logger).Log("err", err)
		return "", "", err
	}

	logger.Log("create order", id)

	return id.String(), "Success", nil
}

func (s orderService) UpdateOrder(ctx context.Context, id string, plate string, score int64) (string, error) {
	logger := log.With(s.logger, "method", "UpdateOrder")

	uid := uuid.MustParse(id)
	if uid.String() == "" {
		level.Error(logger).Log("err: invalid id %s", id)
		return "Invalid Id", nil
	}

	order := domain.Order{
		ID:    uid,
		Plate: plate,
		Score: score,
	}

	if err := s.repostory.UpdateOrder(ctx, order); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("update order", id)

	return "Success", nil
}

func (s orderService) GetOrder(ctx context.Context, id string) (*domain.Order, error) {
	logger := log.With(s.logger, "method", "GetOrder")

	order, err := s.repostory.GetOrder(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return order, err
	}

	logger.Log("Get order", id)

	return order, nil
}
