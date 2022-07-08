package common

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	NotAuthorized = errors.New("not authorized")
	NotFound      = errors.New("not found")
)

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError called without error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
