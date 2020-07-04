package items

import "context"

//go:generate mockgen -destination=../mocks/mock_items/mock_repository.go -package=mock_items github.com/alanyeung95/GoProjectDemo/pkg/items Repository

// Repository is the item repo
type Repository interface {
	Upsert(ctx context.Context, id string, item Item) (*Item, error)
	Find(ctx context.Context, id string) (*Item, error)
	Update(ctx context.Context, id string, model interface{}) error
}
