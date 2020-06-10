package items

import (
	"encoding/json"
	"net/http"

	"github.com/alanyeung95/GoProjectDemo/pkg/errors"
	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

// NewHandler return handler that serves the items service
func NewHandler(srv Service) http.Handler {
	h := handlers{srv}
	r := chi.NewRouter()
	r.Get("/{id}", h.handleGetItemsSample)
	r.Post("/", h.handleCreateItem)
	r.Get("/{id}/raw", h.handleGetItemTextByID)
	return r
}

type handlers struct {
	svc Service
}

func (h *handlers) handleGetItemsSample(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	response, err := h.svc.GetItemByID(ctx, r, id)
	if err != nil {
		kithttp.DefaultErrorEncoder(ctx, errors.NewBadRequest(err), w)
		return
	}

	kithttp.EncodeJSONResponse(ctx, w, response)
}

func (h *handlers) handleCreateItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model Item
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		kithttp.DefaultErrorEncoder(ctx, errors.NewBadRequest(err), w)
		return
	}

	response, err := h.svc.CreateItem(ctx, r, &model)
	if err != nil {
		kithttp.DefaultErrorEncoder(ctx, errors.NewBadRequest(err), w)
		return
	}

	kithttp.EncodeJSONResponse(ctx, w, response)
}

func (h *handlers) handleGetItemTextByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	response, err := h.svc.GetItemTextByID(ctx, r, id)
	if err != nil {
		kithttp.DefaultErrorEncoder(ctx, errors.NewBadRequest(err), w)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(response))
}
