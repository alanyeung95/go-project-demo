package items

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Repository interface {
	Find(ctx context.Context, id string) (bson.Raw, error)
	// todo
	//Find(ctx context.Context, id string) (interface{}, error)
	Update(ctx context.Context, id string, model interface{}) error
}
