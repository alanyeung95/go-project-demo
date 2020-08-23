package users

import (
	"context"
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
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
	user.Password = hashAndSalt([]byte(user.Password))
	return s.repository.Upsert(ctx, newID, *user)
}

func (s *service) GetUserByID(ctx context.Context, r *http.Request, id string) (*User, error) {
	return s.repository.Find(ctx, id)
}

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	} // GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}
