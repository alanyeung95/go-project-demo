package items

import (
	"context"
	"net/http"
)

// Service interface
type Service interface {
	CreateItem(ctx context.Context, r *http.Request) (*Item, error)
}

type service struct {
	repository Repository
}

// NewService start the new service
func NewService(repository Repository) (Service, error) {
	return &service{repository}, nil
}

func (s *service) CreateItem(ctx context.Context, r *http.Request) (*Item, error) {
	return s.repository.Upsert(ctx, "test_id", Item{Name: "testing", Price: 123})
}
