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
			return models.CreateResponse{err.Error(), v}, nil
		}

		return models.CreateResponse{"", v}, nil
	}
}

func makeGetAllEndpoint(svc UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v, _ := svc.GetAll()

		return models.GetAllResponse{"", v}, nil
	}
}
