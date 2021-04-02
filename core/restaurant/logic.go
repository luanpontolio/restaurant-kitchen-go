package restaurant

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/google/uuid"
)

type service struct {
	repostory RestaurantRespository
	logger    log.Logger
}

func NewService(rep RestaurantRespository, logger log.Logger) RestaurantService {
	return &service{
		repostory: rep,
		logger:    logger,
	}
}

func (s service) CreateOrder(ctx context.Context, plate string, score int64) (string, string, error) {
	logger := log.With(s.logger, "method", "CreateOrder")

	id := uuid.New()
	order := Order{
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

func (s service) UpdateOrder(ctx context.Context, id string, plate string, score int64) (string, string, error) {
	logger := log.With(s.logger, "method", "UpdateOrder")

	if id == "" {
		level.Error(logger).Log("err: invalid id %s", id)
		return "", "Invalid Id", nil
	}

	order := Order{
		ID:    uuid.MustParse(id),
		Plate: plate,
		Score: score,
	}
	fmt.Printf("object Order %v", order)
	if err := s.repostory.UpdateOrder(ctx, order); err != nil {
		level.Error(logger).Log("err", err)
		return "", "", err
	}

	logger.Log("update order", id)

	return id, "Success", nil
}

func (s service) CreateCook(ctx context.Context, name string, score int64) (string, string, error) {
	logger := log.With(s.logger, "method", "CreateCook")

	id := uuid.New()
	cook := Cook{
		ID:    id,
		Name:  name,
		Score: score,
	}

	if err := s.repostory.CreateCook(ctx, cook); err != nil {
		level.Error(logger).Log("err", err)
		return "", "", err
	}

	logger.Log("create cook", id)

	return id.String(), "Success", nil
}

func (s service) UpdateCook(ctx context.Context, id string, score int64) (string, string, error) {
	logger := log.With(s.logger, "method", "UpdateCook")

	if id == "" {
		level.Error(logger).Log("err: invalid id %s", id)
		return "", "Invalid Id", nil
	}

	cook := Cook{
		ID:    uuid.MustParse(id),
		Score: score,
	}
	fmt.Printf("object Order %v", cook)
	if err := s.repostory.UpdateCook(ctx, cook); err != nil {
		level.Error(logger).Log("err", err)
		return "", "", err
	}

	logger.Log("update order", id)

	return id, "Success", nil
}
