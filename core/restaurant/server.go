package restaurant

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("GET").Path("/v1/order").Queries("state", "{state}").Handler(httptransport.NewServer(
		endpoints.ListAllOrder,
		decodeFilterParamsReq,
		encodeResponse,
	))

	r.Methods("POST").Path("/v1/order").Handler(httptransport.NewServer(
		endpoints.CreateOrder,
		decodeOrderReq,
		encodeResponse,
	))

	r.Methods("PUT").Path("/v1/order/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateOrder,
		decodeOrderReq,
		encodeResponse,
	))

	r.Methods("POST").Path("/v1/cook").Handler(httptransport.NewServer(
		endpoints.CreateCook,
		decodeCookReq,
		encodeResponse,
	))

	r.Methods("PUT").Path("/v1/cook/{id}").Handler(httptransport.NewServer(
		endpoints.UpdateCook,
		decodeCookReq,
		encodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
