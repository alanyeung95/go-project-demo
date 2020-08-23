package users

import (
	"encoding/json"
	"net/http"

	"github.com/alanyeung95/GoProjectDemo/pkg/errors"
	"github.com/go-chi/chi"

	kithttp "github.com/go-kit/kit/transport/http"
)

// NewHandler return handler that serves the users service
func NewHandler(srv Service) http.Handler {
	h := handlers{srv}
	r := chi.NewRouter()
	r.Get("/{id}", h.handleGetUser)
	r.Post("/", h.handleCreateUser)
	r.Post("/login", h.handleUserLogin)
	return r
}

type handlers struct {
	svc Service
}

func (h *handlers) handleGetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	response, err := h.svc.GetUserByID(ctx, r, id)
	if err != nil {
		kithttp.DefaultErrorEncoder(ctx, errors.NewBadRequest(err), w)
		return
	}

	kithttp.EncodeJSONResponse(ctx, w, response)
}

func (h *handlers) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var model User
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		kithttp.DefaultErrorEncoder(ctx, errors.NewBadRequest(err), w)
		return
	}

	response, err := h.svc.CreateUser(ctx, r, &model)
	if err != nil {
		kithttp.DefaultErrorEncoder(ctx, errors.NewBadRequest(err), w)
		return
	}

	kithttp.EncodeJSONResponse(ctx, w, response)
}

func (h *handlers) handleUserLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request UserLoginParam
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		kithttp.DefaultErrorEncoder(ctx, errors.NewBadRequest(err), w)
		return
	}

	response, err := h.svc.UserLogin(ctx, r, &request)
	if err != nil {
		kithttp.DefaultErrorEncoder(ctx, err, w)
		return
	}

	kithttp.EncodeJSONResponse(ctx, w, response)
}
