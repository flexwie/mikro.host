package common

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	stdhttp "net/http"
	"strings"
	"time"
)

func LoggingMiddleware(logger log.Logger, svcName string) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			start := time.Now()
			body, _ := json.Marshal(request)

			defer logger.Log("method", svcName, "duration", time.Since(start).String(), "body", string(body))
			return e(ctx, request)
		}
	}
}

func AuthMiddleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if ctx.Value(TokenContextKey) == "123" {
				return next(ctx, request)
			}

			return nil, NotAuthorized
		}
	}
}

var TokenContextKey = "Token"

func AuthBefore() http.RequestFunc {
	return func(ctx context.Context, request *stdhttp.Request) context.Context {

		token, ok := TokenValue(request.Header.Get("Authorization"))
		if !ok {
			return ctx
		}

		return context.WithValue(ctx, TokenContextKey, string(token))
	}
}

func TokenValue(val string) (token []byte, ok bool) {
	if len(val) < 8 || !strings.EqualFold(val[0:7], "BEARER ") {
		return nil, false
	}

	return []byte(val[7:]), true
}
