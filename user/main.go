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

func GetCreate() *httptransport.Server {
	var createEndpoint endpoint.Endpoint
	createEndpoint = makeCreateEndpoint(svc)
	createEndpoint = common.LoggingMiddleware(logger, "create")(createEndpoint)

	httptransport.ServerBefore()
	createHandler := httptransport.NewServer(createEndpoint, decodeCreateRequest, encodeResponse)
	return createHandler
}

func GetGetAll() *httptransport.Server {
	var getAllEndpoint endpoint.Endpoint
	getAllEndpoint = makeGetAllEndpoint(svc)
	getAllEndpoint = common.LoggingMiddleware(logger, "getAll")(getAllEndpoint)

	httptransport.ServerBefore()
	getAllHandler := httptransport.NewServer(getAllEndpoint, decodeGetAllRequest, encodeResponse)
	return getAllHandler
}

func main() {
	Db = common.GetDb("test.db")

	http.Handle("/", common.Method(http.MethodPost, GetCreate()))
	http.Handle("/", common.Method(http.MethodGet, GetGetAll()))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func init() {
	logger = log.NewLogfmtLogger(os.Stderr)
	svc = userService{}
}
