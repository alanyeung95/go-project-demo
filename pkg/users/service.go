package users

import (
	"context"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// Service interface
type Service interface {
	CreateUser(ctx context.Context, r *http.Request, user *User) (*User, error)
	GetUserByID(ctx context.Context, r *http.Request, id string) (*User, error)
}

type service struct {
	repository Repository
}

// NewService start the new service
func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) CreateUser(ctx context.Context, r *http.Request, user *User) (*User, error) {
	var newID = uuid.NewV4().String()
	user.ID = newID
	return s.repository.Upsert(ctx, newID, *user)
}

func (s *service) GetUserByID(ctx context.Context, r *http.Request, id string) (*User, error) {
	return s.repository.Find(ctx, id)
}
