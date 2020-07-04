package items

import (
	"context"
	"net/http"
	"strconv"

	uuid "github.com/satori/go.uuid"

	"github.com/alanyeung95/GoProjectDemo/pkg/errors"
)

// Service interface
type Service interface {
	CreateItem(ctx context.Context, r *http.Request, item *Item) (*Item, error)
	GetItemByID(ctx context.Context, r *http.Request, id string) (*Item, error)
	GetItemTextByID(ctx context.Context, r *http.Request, id string) (string, error)
}

type service struct {
	repository Repository
}

// NewService start the new service
func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) CreateItem(ctx context.Context, r *http.Request, item *Item) (*Item, error) {
	var newID = uuid.NewV4().String()
	item.ID = newID
	return s.repository.Upsert(ctx, newID, *item)
}

func (s *service) GetItemByID(ctx context.Context, r *http.Request, id string) (*Item, error) {
	return s.repository.Find(ctx, id)
}

func (s *service) GetItemTextByID(ctx context.Context, r *http.Request, id string) (string, error) {
	item, err := s.repository.Find(ctx, id)
	if err != nil {
		return "", errors.NewBadRequest(err)
	}
	var itemText string
	itemText = "ID: " + item.ID + "\n" + "Name: " + item.Name + "\n" + "Price: " + strconv.Itoa(item.Price)
	return itemText, nil
}
