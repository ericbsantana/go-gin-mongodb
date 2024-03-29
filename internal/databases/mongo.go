package databases

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitializeMongoDBConnection() (*mongo.Database, *mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	c, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatalf("Error connecting client to MongoDB: %v", err)
		return nil, nil, err
	}

	db := c.Database("rest-api-template")

	return db, c, nil
}
