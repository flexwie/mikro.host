package http

import (
	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"mikro.host/cluster/transport"
	"mikro.host/common"
	"net/http"
)

func NewService(endpoints transport.Endpoints, options []kithttp.ServerOption, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	errorLogger := kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger))
	errorEncoder := kithttp.ServerErrorEncoder(common.EncodeError)
	options = append(options, errorLogger, errorEncoder)

	r.Methods(http.MethodPost).Path("/api/cluster").Handler(kithttp.NewServer(
		endpoints.CreateCluster,
		decodeCreateClusterRequest,
		encodeResponse,
		options...,
	))

	return r
}
