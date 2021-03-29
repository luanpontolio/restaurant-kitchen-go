package order

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	OrderRequest struct {
		Id    string `json:"id:omitempty"`
		Plate string `json:"plate"`
		Score int64  `json:"score"`
	}
	OrderResponse struct {
		Id string `json:"id"`
		Ok string `json:"ok"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeOrderReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req OrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
