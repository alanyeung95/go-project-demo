package grpc

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

// NewHandler return handler that serves the grpc service
func NewHandler(grpcClient *GreeterClient) http.Handler {
	r := chi.NewRouter()

	r.Get("/greet/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")

		message, err := grpcClient.SayHello(name)
		if err != nil {
			http.Error(w, "Failed to greet via gRPC: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": message})
	})

	return r
}
