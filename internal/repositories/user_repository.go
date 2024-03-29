package repositories

import (
	"context"
	"rest-api/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Database
}

func UserRepositoryFromDatabase(db *mongo.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	collection := r.db.Collection("users")

	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var results []models.User

	if err := cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return []models.User{}, nil
}
