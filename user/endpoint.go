package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"mikro.host/models"
)

func makeCreateEndpoint(svc UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(models.CreateRequest)
		v, err := svc.Create(req)

		if err != nil {
			return nil, err
		}

		return models.CreateResponse{"", v}, nil
	}
}

func makeGetAllEndpoint(svc UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v, err := svc.GetAll()

		if err != nil {
			return nil, err
		}

		return models.GetAllResponse{"", v}, nil
	}
}

func makeGetOneEndpoint(svc UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(models.GetOneRequest)

		v, err := svc.Get(req.Id)

		if err != nil {
			return nil, err
		}

		return models.GetOneResponse{Value: v}, nil
	}
}

func makeUpdateEndpoint(svc UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(models.UpdateRequest)

		v, err := svc.Update(req)

		if err != nil {
			return nil, err
		}

		return models.UpdateResponse{v}, nil
	}
}
