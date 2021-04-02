package restaurant

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateOrder endpoint.Endpoint
	UpdateOrder endpoint.Endpoint

	CreateCook endpoint.Endpoint
	UpdateCook endpoint.Endpoint
}

func MakeEndpoints(s RestaurantService) Endpoints {
	return Endpoints{
		CreateOrder: makeCreateOrderEndpoint(s),
		UpdateOrder: makeUpdateOrderEndpoint(s),

		CreateCook: makeCreateCookEndpoint(s),
		UpdateCook: makeUpdateCookEndpoint(s),
	}
}

func makeCreateOrderEndpoint(s RestaurantService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(OrderRequest)
		id, msg, err := s.CreateOrder(ctx, req.Plate, req.Score)
		if err != nil {
			return Response{Ok: "Invalid data"}, err
		}

		return Response{Id: id, Ok: msg}, err
	}
}

func makeUpdateOrderEndpoint(s RestaurantService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(OrderRequest)
		id, msg, err := s.UpdateOrder(ctx, req.Id, req.Plate, req.Score)
		if err != nil {
			return Response{Ok: "Invalid data"}, err
		}

		return Response{Id: id, Ok: msg}, err
	}
}

func makeCreateCookEndpoint(s RestaurantService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CookRequest)
		id, msg, err := s.CreateCook(ctx, req.Name, req.Score)
		if err != nil {
			return Response{Ok: "Invalid data"}, err
		}

		return Response{Id: id, Ok: msg}, err
	}
}

func makeUpdateCookEndpoint(s RestaurantService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CookRequest)
		id, msg, err := s.UpdateCook(ctx, req.Id, req.Score)
		if err != nil {
			return Response{Ok: "Invalid data"}, err
		}

		return Response{Id: id, Ok: msg}, err
	}
}
