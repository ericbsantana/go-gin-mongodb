package mongo

import (
	"context"
	"log"
	user "rest-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		db: db,
	}
}

func ConnectMongoDB() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatalf("Error connecting client to MongoDB: %v", err)
		return nil, err
	}

	db := client.Database("rest-api-template")

	return db, nil
}

func (r *Repository) FindAll() ([]user.UserModel, error) {
	collection := r.db.Collection("users")

	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var results []user.UserModel

	if err := cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
