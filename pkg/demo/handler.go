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
	r.Get("/error_demo", h.handleGetErrorDemoSample)
	return r
}

type handlers struct {
	svc Service
}

func (h *handlers) handleGetDemoSample(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!"))
}

func (h *handlers) handleGetErrorDemoSample(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := h.svc.DemoError(ctx)
	w.Write([]byte("error demo: " + err.Error()))
}
