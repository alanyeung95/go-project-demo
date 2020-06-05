package items

import "context"

// Repository is the item repo
type Repository interface {
	Upsert(ctx context.Context, id string, item Item) (*Item, error)
	Find(ctx context.Context, id string) (interface{}, error)
	Update(ctx context.Context, id string, model interface{}) error
}
