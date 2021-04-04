package restaurant

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type (
	OrderRequest struct {
		Id         string `json:"id:omitempty"`
		Plate      string `json:"plate"`
		Score      int64  `json:"score"`
		State      int64  `json:"state"`
		DeliveryAt bool   `json:"delivery_at"`
	}
	CookRequest struct {
		Id    string `json:"id:omitempty"`
		Name  string `json:"name"`
		Score int64  `json:"score"`
	}
	Response struct {
		Id   string   `json:"id"`
		Ok   string   `json:"ok"`
		Data []*Order `json:"data"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeFilterParamsReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req OrderRequest
	req.State, _ = strconv.ParseInt(r.FormValue("state"), 10, 64)
	req.DeliveryAt, _ = strconv.ParseBool(r.FormValue("delivery_at"))

	return req, nil
}

func decodeOrderReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req OrderRequest

	req.Id = mux.Vars(r)["id"]
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		panic(err)
	}

	return req, nil
}

func decodeCookReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CookRequest
	req.Id = mux.Vars(r)["id"]
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
