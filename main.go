package main

import (
	"context"
	"net/http"
	database "rest-api/internal/databases"
	"rest-api/internal/handlers"
	"rest-api/internal/repositories"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func setupRouter(db *mongo.Database) *gin.Engine {
	r := gin.Default()

	userRepository := repositories.UserRepositoryFromDatabase(db)
	userHandler := handlers.UserHandlerFromRepository(userRepository)

	r.GET("/users", userHandler.Find)
	r.POST("/users", userHandler.Create)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	return r
}

func main() {
	db, c, err := database.InitializeMongoDBConnection()

	if err != nil {
		panic(err)
	}

	r := setupRouter(db)

	defer c.Disconnect(context.Background())

	r.Run(":8080")
}
