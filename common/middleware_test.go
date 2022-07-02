package common

import (
	"bytes"
	"context"
	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestTokenValue(t *testing.T) {
	value, ok := TokenValue("Bearer 123")

	assert.True(t, ok)

	if string(value) != "123" {
		t.Fatalf("expected 123 got %s", string(value))
	}
	assert.Equalf(t, "123", string(value), "expected 123 but got %s", string(value))
}

func TestAuthBefore(t *testing.T) {
	reqFunc := AuthBefore()

	ctx := context.Background()
	req := http.Request{Header: map[string][]string{"Authorization": {"Bearer 123"}}}

	ctx = reqFunc(ctx, &req)
	value := ctx.Value(TokenContextKey)

	assert.Equal(t, "123", value)
}

func TestAuthMiddleware(t *testing.T) {
	authEndpoint := AuthMiddleware()(func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return ctx, nil
	})

	// no token
	_, err := authEndpoint(context.Background(), nil)
	assert.Error(t, err)

	// wrong token
	_, err = authEndpoint(context.WithValue(context.Background(), TokenContextKey, "1"), nil)
	assert.Equal(t, err, NotAuthorized)

	// right token
	r, err := authEndpoint(context.WithValue(context.Background(), TokenContextKey, "123"), nil)
	assert.Nil(t, err)
	assert.NotNil(t, r)
}

func TestLoggingMiddleware(t *testing.T) {
	var b bytes.Buffer
	logger := log.NewLogfmtLogger(&b)

	loggingEndpoint := LoggingMiddleware(logger, "log")(func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return ctx, nil
	})

	_, err := loggingEndpoint(context.Background(), nil)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(b.String(), "method=log"))
}
