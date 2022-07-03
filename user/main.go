package main

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"gorm.io/gorm"
	"mikro.host/common"
	"net/http"
	"os"
)

var Db *gorm.DB
var logger log.Logger
var svc userService

func CreateHandler() *httptransport.Server {
	var createEndpoint endpoint.Endpoint
	createEndpoint = makeCreateEndpoint(svc)
	createEndpoint = common.LoggingMiddleware(logger, "create")(createEndpoint)

	httptransport.ServerBefore()
	createHandler := httptransport.NewServer(createEndpoint, decodeCreateRequest, encodeResponse)
	return createHandler
}

func GetAllHandler() *httptransport.Server {
	var getAllEndpoint endpoint.Endpoint
	getAllEndpoint = makeGetAllEndpoint(svc)
	getAllEndpoint = common.LoggingMiddleware(logger, "getAll")(getAllEndpoint)

	httptransport.ServerBefore()
	getAllHandler := httptransport.NewServer(getAllEndpoint, decodeGetAllRequest, encodeResponse)
	return getAllHandler
}

func GetOneHandler() *httptransport.Server {
	var getOneEndpoint endpoint.Endpoint
	getOneEndpoint = makeGetOneEndpoint(svc)
	getOneEndpoint = common.LoggingMiddleware(logger, "getAll")(getOneEndpoint)

	httptransport.ServerBefore()
	getOneHandler := httptransport.NewServer(getOneEndpoint, decodeGetOneRequest, encodeResponse)
	return getOneHandler
}

func main() {
	Db = common.GetDb(nil)

	http.Handle("/create", CreateHandler())
	http.Handle("/get-all", GetAllHandler())
	http.Handle("/by-id", GetOneHandler())

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func init() {
	logger = log.NewLogfmtLogger(os.Stderr)
	svc = userService{}
}
