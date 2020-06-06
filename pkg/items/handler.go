package items

import (
	"encoding/json"
	"fmt"
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

	var model Item
	fmt.Printf("%+v\n", r.Body)

	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		kithttp.DefaultErrorEncoder(ctx, errors.NewBadRequest(err), w)
		return
	}

	// Note:
	// fmt.Printf("%+v\n", model)

	response, err := h.svc.CreateItem(ctx, r, &model)
	if err != nil {
		kithttp.DefaultErrorEncoder(ctx, errors.NewBadRequest(err), w)
		return
	}
	kithttp.EncodeJSONResponse(ctx, w, response)
}
