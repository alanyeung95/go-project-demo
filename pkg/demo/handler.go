package demo

import (
	"net/http"

	"github.com/go-chi/chi"
)

// NewHandler return handler that serves the demo service
func NewHandler(srv Service) http.Handler {
	h := handlers{srv}
	r := chi.NewRouter()
	r.Get("/", h.handleGetDemoSample)
	return r
}

type handlers struct {
	svc Service
}

func (h *handlers) handleGetDemoSample(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!"))
}
