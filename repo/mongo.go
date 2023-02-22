package repo

import (
	"context"
	"fmt"

	"github.com/karalef/shlink/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

// MongoConfig contains mongodb connection options.
type MongoConfig struct {
	Addr       string
	User, Pass string
	Database   string
	Collection string
}

// NewMongoRepository establishes a new mongodb connection and
// creates a new repository with mongodb collection.
func NewMongoRepository(ctx context.Context, cfg MongoConfig) (*MongoRepository, error) {
	opts := options.Client()
	opts.ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s", cfg.User, cfg.Pass, cfg.Addr))

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return &MongoRepository{
		collection: client.Database(cfg.Database).Collection(cfg.Collection),
	}, nil
}

var _ Repository = (*MongoRepository)(nil)

// MongoRepository is a mongodb storage.
type MongoRepository struct {
	collection *mongo.Collection
}

// Store stores an URL.
func (r *MongoRepository) Store(ctx context.Context, short, origin string) error {
	_, err := r.collection.InsertOne(ctx, model.NewURL(short, origin))
	if err != nil && !mongo.IsDuplicateKeyError(err) {
		return err
	}
	return nil
}

// Get returns an URL by short.
func (r *MongoRepository) Get(ctx context.Context, short string) (*model.URL, error) {
	var u model.URL
	err := r.collection.FindOne(ctx, bson.M{"_id": short}).Decode(&u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// Update updates views count.
func (r *MongoRepository) Update(ctx context.Context, short string, viewsDelta int64) error {
	_, err := r.collection.UpdateByID(ctx, short, bson.M{"$inc": bson.M{"views": viewsDelta}})
	return err
}

// Delete deletes an URL by short.
func (r *MongoRepository) Delete(ctx context.Context, short string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": short})
	return err
}
