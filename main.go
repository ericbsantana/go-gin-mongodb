package main

import (
	"context"
	database "go-gin-mongo/internal/databases"
	"go-gin-mongo/internal/handlers"
	"go-gin-mongo/internal/repositories"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGODB_URI")

	db, c, err := database.InitializeMongoDBConnection(mongoURI)

	if err != nil {
		panic(err)
	}

	r := SetupRouter(db)

	defer c.Disconnect(context.Background())

	r.Run(":8080")
}
