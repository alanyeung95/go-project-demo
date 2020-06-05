package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/alanyeung95/GoProjectDemo/pkg/errors"
	"github.com/alanyeung95/GoProjectDemo/pkg/items"
)

// NewItemRepository is the repo to store item model
func NewItemRepository(client *mongo.Client, database, collection string, enableSharding bool) (*ItemRepository, error) {
	c, err := newCollection(client, database, collection, enableSharding)
	if err != nil {
		return nil, err
	}

	return &ItemRepository{c}, nil
}

type ItemRepository struct {
	collection *mongo.Collection
}

var _ items.Repository = (*ItemRepository)(nil)

func (r *ItemRepository) Find(ctx context.Context, id string) (bson.Raw, error) {
	// todo: return model rather than raw bson
	//	func (r *ItemRepository) Find(ctx context.Context, id string) (interface{}, error) {

	filter := bson.M{"_id": id}
	result, err := r.collection.FindOne(ctx, filter).DecodeBytes()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.NewResourceNotFound(err)
		}
		return nil, err
	}

	//return unmarshalToModel(result)
	return result, nil
}

func (r *ItemRepository) Update(ctx context.Context, id string, model interface{}) error {
	filter := bson.M{"_id": id}
	return replaceOne(ctx, r.collection, filter, model)
}
