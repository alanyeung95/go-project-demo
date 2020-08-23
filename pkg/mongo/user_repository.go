package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/alanyeung95/GoProjectDemo/pkg/errors"
	"github.com/alanyeung95/GoProjectDemo/pkg/users"
)

// NewUserRepository is the repo to store User model
func NewUserRepository(client *mongo.Client, database, collection string, enableSharding bool) (*UserRepository, error) {
	c, err := newCollection(client, database, collection, enableSharding)
	if err != nil {
		return nil, err
	}

	return &UserRepository{c}, nil
}

type UserRepository struct {
	collection *mongo.Collection
}

// interface check
var _ users.Repository = (*UserRepository)(nil)

// Upsert returns the User record being successfully created or updated
func (r *UserRepository) Upsert(ctx context.Context, id string, User users.User) (*users.User, error) {
	var result users.User
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": User,
	}
	if err := upsert(ctx, r.collection, filter, update, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Find returns an User record
func (r *UserRepository) Find(ctx context.Context, id string) (*users.User, error) {
	var result users.User
	filter := bson.M{"_id": id}
	if err := r.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.NewResourceNotFound(err)
		}
		return nil, err
	}
	return &result, nil
}

// FindByEmail returns an User record
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*users.User, error) {
	var result users.User
	filter := bson.M{"email": email}
	if err := r.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.NewResourceNotFound(err)
		}
		return nil, err
	}
	return &result, nil
}
