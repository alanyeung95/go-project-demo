package users

import (
	"context"
	"log"
	"net/http"

	"github.com/alanyeung95/GoProjectDemo/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// Service interface
type Service interface {
	CreateUser(ctx context.Context, r *http.Request, user *User) (*User, error)
	GetUserByID(ctx context.Context, r *http.Request, id string) (*User, error)
	UserLogin(ctx context.Context, r *http.Request, loginInfo *UserLoginParam) (*LoginResponse, error)
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

func (s *service) UserLogin(ctx context.Context, r *http.Request, loginInfo *UserLoginParam) (*LoginResponse, error) {
	user, err := s.repository.FindByEmail(ctx, loginInfo.Email)
	if err != nil {
		return nil, errors.NewResourceNotFound(err)
	}

	var resp = LoginResponse{}
	pwdMatch := comparePasswords(user.Password, []byte(loginInfo.Password))
	if pwdMatch {
		resp.Status = true
	} else {
		resp.Status = false
	}

	return &resp, nil
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

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
