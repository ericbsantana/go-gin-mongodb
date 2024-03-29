package config_test

import (
	"context"
	"fmt"
	"rest-api/internal/databases"
	"rest-api/internal/models"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	TestDBInstance *mongo.Database
)

type TestDB struct {
	Database  *mongo.Database
	Address   string
	Container testcontainers.Container
}

func SetupTestDB() *TestDB {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*60)

	container, db, addr, err := createMongoDBContainer(ctx)

	if err != nil {
		panic(fmt.Sprintf("failed to setup test database: %v", err))
	}

	return &TestDB{
		Container: container,
		Database:  db,
		Address:   addr,
	}
}

func (tdb *TestDB) TearDown() {
	_ = tdb.Container.Terminate(context.Background())
}

func createMongoDBContainer(ctx context.Context) (testcontainers.Container, *mongo.Database, string, error) {
	var env = map[string]string{
		"MONGO_INITDB_ROOT_USERNAME": "kiwi",
		"MONGO_INITDB_ROOT_PASSWORD": "kiwi",
		"MONGO_INITDB_DATABASE":      "rest-api-test-db",
	}
	var port = "27017/tcp"

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "mongo",
			ExposedPorts: []string{port},
			Env:          env,
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return container, nil, "", fmt.Errorf("failed to start container: %v", err)
	}

	p, err := container.MappedPort(ctx, "27017")
	if err != nil {
		return container, nil, "", fmt.Errorf("failed to get container external port: %v", err)
	}

	uri := fmt.Sprintf("mongodb://kiwi:kiwi@localhost:%s", p.Port())
	db, _, err := databases.InitializeMongoDBConnection(uri)

	if err != nil {
		return container, db, uri, fmt.Errorf("failed to establish database connection: %v", err)
	}

	return container, db, uri, nil
}

func SeedTestDatabase() {
	var users = []interface{}{
		models.User{
			ID: func() primitive.ObjectID {
				id, _ := primitive.ObjectIDFromHex("6607077651565dc6fbb91859")
				return id
			}(),
			Username: "oppenheimer",
			Email:    "oppenheimer@example.com",
		},
		models.User{
			ID:    primitive.NewObjectID(),
			Email: "aristotle@greek.com",
		},
	}

	_, err := TestDBInstance.Collection("users").InsertMany(context.Background(), users)
	if err != nil {
		panic(fmt.Sprintf("failed to populate database: %v", err))
	}
}

func ClearDB() {
	_, err := TestDBInstance.Collection("users").DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		panic(fmt.Sprintf("failed to clear database: %v", err))
	}
}
