package order

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/google/uuid"
)

type service struct {
	repostory Repository
	logger    log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}

func (s service) CreateOrder(ctx context.Context, plate string, score int64) (string, error) {
	logger := log.With(s.logger, "method", "CreateOrder")

	id := uuid.New()
	order := Order{
		ID:    id,
		Plate: plate,
		Score: score,
	}

	if err := s.repostory.CreateOrder(ctx, order); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create order", id)

	return "Success", nil
}

func (s service) GetOrder(ctx context.Context, id string) (*Order, error) {
	logger := log.With(s.logger, "method", "GetOrder")

	order, err := s.repostory.GetOrder(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return order, err
	}

	logger.Log("Get order", id)

	return order, nil
}
