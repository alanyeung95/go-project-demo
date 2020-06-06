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

// interface check
var _ items.Repository = (*ItemRepository)(nil)

// Upsert returns the item record being successfully created or updated
func (r *ItemRepository) Upsert(ctx context.Context, id string, item items.Item) (*items.Item, error) {
	var result items.Item
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": item,
	}
	if err := upsert(ctx, r.collection, filter, update, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Find returns an item record
func (r *ItemRepository) Find(ctx context.Context, id string) (*items.Item, error) {
	var result items.Item
	filter := bson.M{"_id": id}
	if err := r.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.NewResourceNotFound(err)
		}
		return nil, err
	}
	return &result, nil
}

// Update returns the item record being successfully updated
func (r *ItemRepository) Update(ctx context.Context, id string, model interface{}) error {
	filter := bson.M{"_id": id}
	return replaceOne(ctx, r.collection, filter, model)
}

/**
func unmarshalToModel(data bson.Raw) (interface{}, error) {
	var model items.Item
	if err := unmarshalBson(data, &model); err != nil {
		return nil, err
	}
	return model, nil
}

func unmarshalBson(data bson.Raw, val interface{}) error {
	return bson.Unmarshal(data, val)
}
**/
