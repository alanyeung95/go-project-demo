package items

import (
	"net/http"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

// NewHandler return handler that serves the items service
func NewHandler(srv Service) http.Handler {
	h := handlers{srv}
	r := chi.NewRouter()
	r.Get("/", h.handleGetItemsSample)
	r.Post("/{id}", h.handleCreateItem)
	return r
}

type handlers struct {
	svc Service
}

func (h *handlers) handleGetItemsSample(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!"))
}

func (h *handlers) handleCreateItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	response, err := h.svc.CreateItem(ctx, r)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	kithttp.EncodeJSONResponse(ctx, w, response)
}
