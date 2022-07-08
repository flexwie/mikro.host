package http

import (
	"context"
	"encoding/json"
	"mikro.host/common"
	"mikro.host/models"
	"net/http"
)

type Errorer interface {
	Error() error
}

func decodeCreateClusterRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req models.CreateClusterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(Errorer); ok && err.Error() != nil {
		common.EncodeError(ctx, err.Error(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
