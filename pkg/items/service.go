package items

import (
	"context"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// Service interface
type Service interface {
	CreateItem(ctx context.Context, r *http.Request, item *Item) (*Item, error)
}

type service struct {
	repository Repository
}

// NewService start the new service
func NewService(repository Repository) (Service, error) {
	return &service{repository}, nil
}

func (s *service) CreateItem(ctx context.Context, r *http.Request, item *Item) (*Item, error) {
	var newID = uuid.NewV4().String()
	item.ID = newID
	return s.repository.Upsert(ctx, newID, *item)
}
