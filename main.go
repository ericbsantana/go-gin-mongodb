package main

import (
	"fmt"
	"net/http"
	mongo "rest-api/internal/repository"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Repository *mongo.Repository
}

func NewHandler(repo *mongo.Repository) *Handler {
	return &Handler{
		Repository: repo,
	}
}

func (h *Handler) GetAllUsersHandler(c *gin.Context) {
	users, err := h.Repository.FindAll()

	if err != nil {
		fmt.Printf("Error getting users from MongoDB: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	db, err := mongo.ConnectMongoDB()

	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v", err)
	}

	repo := mongo.NewRepository(db)

	handler := NewHandler(repo)

	r.GET("/users", handler.GetAllUsersHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
