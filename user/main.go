package main

import (
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
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

func GetGetOne() *httptransport.Server {
	var getOneEndpoint endpoint.Endpoint
	getOneEndpoint = makeGetOneEndpoint(svc)
	getOneEndpoint = common.LoggingMiddleware(logger, "getAll")(getOneEndpoint)

	httptransport.ServerBefore()
	getOneHandler := httptransport.NewServer(getOneEndpoint, decodeGetOneRequest, encodeResponse)
	return getOneHandler
}

func main() {
	Db = common.GetDb(nil)

	r := mux.NewRouter()

	r.HandleFunc("/", GetCreate().ServeHTTP).Methods(http.MethodPost)
	r.HandleFunc("/", GetGetAll().ServeHTTP).Methods(http.MethodGet)
	r.HandleFunc("/{id}", GetGetOne().ServeHTTP).Methods(http.MethodGet)

	//http.Handle("/", common.Method(http.MethodPost, GetCreate()))
	//http.Handle("/", common.Method(http.MethodGet, GetGetAll()))
	//http.Handle("/:id", common.Method(http.MethodGet, GetGetOne()))

	http.Handle("/", r)
}

func init() {
	logger = log.NewLogfmtLogger(os.Stderr)
	svc = userService{}
}
