package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/alanyeung95/GoProjectDemo/pkg/errors"
)

// NewClient is the function to create an new mongo connection
func NewClient(addresses, username, password, database string) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s/%s", username, password, addresses, database)
	opts := options.Client().
		ApplyURI(uri)
	client, err := mongo.NewClient(opts)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client, nil
}

func newCollection(client *mongo.Client, database, collection string, enableSharding bool) (*mongo.Collection, error) {
	//func newCollection(client *mongo.Client, database, collection string, enableSharding bool, indexes []mongo.IndexModel) (*mongo.Collection, error) {
	c := client.Database(database).Collection(collection)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if enableSharding {
		if err := ShardCollection(ctx, client, database, collection); err != nil {
			return nil, err
		}
	}

	// EnsureIndexes without timeout
	//if err := EnsureIndexes(context.Background(), c, indexes); err != nil {
	//	return nil, err
	//}

	return c, nil
}

// ShardCollection shards the given collection with hashed shard key on the _id.
func ShardCollection(ctx context.Context, client *mongo.Client, databaseName string, collectionName string) error {
	_, err := client.Database(databaseName).Collection(collectionName).Indexes().
		CreateOne(ctx, mongo.IndexModel{Keys: bson.D{{"_id", "hashed"}}})
	if err != nil {
		return err
	}

	adminDatabase := client.Database("admin")
	result := adminDatabase.RunCommand(ctx,
		bson.D{
			{"shardCollection", fmt.Sprintf("%s.%s", databaseName, collectionName)},
			{"key", bson.D{{"_id", "hashed"}}}},
	)
	return result.Err()
}

// EnsureIndexes ensures indexes are existed or created if not exists
/**
func EnsureIndexes(ctx context.Context, c *mongo.Collection, idxs []mongo.IndexModel) error {
	if len(idxs) < 1 {
		return nil
	}
	_, err := c.Indexes().CreateMany(ctx, idxs)
	if driverErr, ok := err.(driver.Error); ok {
		if driverErr.Code == 85 || driverErr.Code == 86 {
			// Ignore IndexOptionsConflict (85) or IndexKeySpecsConflict (86)
			return nil
		}
	}
	return err
}
**/

func create(ctx context.Context, c *mongo.Collection, entity interface{}) error {
	_, err := c.InsertOne(ctx, entity)
	if isDuplicate, err := isDuplicateKeyErr(err); isDuplicate {
		return errors.NewError(errors.ServerError, err.Error())
	}
	return err
}

func createMany(ctx context.Context, c *mongo.Collection, entities []interface{}, opts ...*options.InsertManyOptions) error {
	_, err := c.InsertMany(ctx, entities, opts...)
	if isDuplicate, err := isDuplicateKeyErr(err); isDuplicate {
		return errors.NewError(errors.ServerError, err.Error())
	}
	return err
}

// upsert create or update the target document and return created or updated document
func upsert(ctx context.Context, c *mongo.Collection, filter interface{}, update interface{}, result interface{}) error {
	err := c.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetUpsert(true),
		options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func exists(ctx context.Context, c *mongo.Collection, id string) (bool, error) {
	filter := bson.M{"_id": id}
	return existsBy(ctx, c, filter)
}

func existsBy(ctx context.Context, c *mongo.Collection, filter interface{}) (bool, error) {
	err := c.FindOne(ctx, filter).Err()
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func findOne(ctx context.Context, c *mongo.Collection, filter interface{}, result interface{}) error {
	err := c.FindOne(ctx, filter).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.NewResourceNotFound(err)
		}
		return err
	}

	return nil
}

func findAll(ctx context.Context, c *mongo.Collection, filter interface{}, results interface{}, opts ...*options.FindOptions) error {
	cursor, err := c.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}

	return cursor.All(ctx, results)
}

func replaceOne(ctx context.Context, c *mongo.Collection, filter interface{}, entity interface{}) error {
	result, err := c.ReplaceOne(ctx, filter, entity)
	if err != nil {
		return err
	} else if result.MatchedCount == 0 {
		return errors.NewResourceNotFound(mongo.ErrNoDocuments)
	}
	return nil
}

func findAndReplaceOne(ctx context.Context, c *mongo.Collection, filter interface{}, entity interface{}, result interface{}) error {
	err := c.FindOneAndReplace(ctx, filter, entity, options.FindOneAndReplace().SetReturnDocument(options.After)).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.NewResourceNotFound(err)
		}
		return err
	}
	return nil
}

// findAndUpdateOne update the target document and return updated document
func findAndUpdateOne(ctx context.Context, c *mongo.Collection, filter interface{}, update interface{}, result interface{}) error {
	err := c.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.NewResourceNotFound(err)
		}
		return err
	}
	return nil
}
