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

func SetupRouter(db *mongo.Database) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	userRepository := repositories.UserRepositoryFromDatabase(db)
	userHandler := handlers.UserHandlerFromRepository(userRepository)

	r.GET("/users", userHandler.Find)
	r.GET("/users/:id", userHandler.FindByID)
	r.POST("/users", userHandler.Create)
	r.PATCH("/users/:id", userHandler.Update)
	r.DELETE("/users/:id", userHandler.Delete)

	return r
}

func main() {
	db, c, err := database.InitializeMongoDBConnection("mongodb://localhost:27017")

	if err != nil {
		panic(err)
	}

	r := SetupRouter(db)

	defer c.Disconnect(context.Background())

	r.Run(":8080")
}
