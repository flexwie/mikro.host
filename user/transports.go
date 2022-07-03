package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"mikro.host/models"
	"net/http"
)

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req models.CreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeGetAllRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeGetOneRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		return nil, errors.New("missing id in path")
	}

	return &models.GetOneRequest{vars["id"]}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
