package repositories

import (
	"context"
	"go-gin-mongo/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *UserRepository) Find() ([]models.User, error) {
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

	return results, nil
}

func (r *UserRepository) Create(user models.User) (*mongo.InsertOneResult, error) {
	collection := r.db.Collection("users")

	result, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	collection := r.db.Collection("users")

	var user models.User

	err := collection.FindOne(context.Background(), bson.D{{Key: "email", Value: email}}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByID(id string) (*models.User, error) {
	collection := r.db.Collection("users")

	var user models.User

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	err = collection.FindOne(context.Background(), bson.D{{Key: "_id", Value: objectID}}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(id string, user models.User) (*mongo.UpdateResult, error) {
	collection := r.db.Collection("users")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	updateFields := bson.D{}

	if user.Email != "" {
		updateFields = append(updateFields, bson.E{Key: "email", Value: user.Email})
	}

	if user.Username != "" {
		updateFields = append(updateFields, bson.E{Key: "username", Value: user.Username})
	}

	result, err := collection.UpdateOne(context.Background(), bson.D{{Key: "_id", Value: objectID}}, bson.D{{Key: "$set", Value: updateFields}})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserRepository) Delete(id string) (*mongo.DeleteResult, error) {
	collection := r.db.Collection("users")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	result, err := collection.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: objectID}})

	if err != nil {
		return nil, err
	}

	return result, nil
}
