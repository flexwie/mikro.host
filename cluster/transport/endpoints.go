package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"mikro.host/cluster/service"
	"mikro.host/models"
)

type Endpoints struct {
	CreateCluster endpoint.Endpoint
}

func MakeServerEndpoints(svc service.Service) Endpoints {
	return Endpoints{
		CreateCluster: makeCreateClusterEndpoint(svc),
	}
}

func makeCreateClusterEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(models.CreateClusterRequest)
		id, err := svc.CreateCluster(ctx, req)
		return models.CreateClusterResponse{Value: id}, err
	}
}
