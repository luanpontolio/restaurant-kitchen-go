package order

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateOrder endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreateOrder: makeCreateOrderEndpoint(s),
	}
}

func makeCreateOrderEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(OrderRequest)
		msg, err := s.CreateOrder(ctx, req.Plate, req.Score)
		if err != nil {
			return OrderResponse{Ok: "Invalid data"}, err
		}

		return OrderResponse{Ok: msg}, err
	}
}
