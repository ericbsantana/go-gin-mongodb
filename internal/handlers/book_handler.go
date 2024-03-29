package handlers

import (
	"net/http"
	"rest-api/internal/repositories"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	Repository *repositories.BookRepository
}

func NewBookHandler(repo *repositories.BookRepository) *BookHandler {
	return &BookHandler{
		Repository: repo,
	}
}

func (h *BookHandler) GetAllBooksHandler(c *gin.Context) {
	books, err := h.Repository.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}

	c.JSON(http.StatusOK, books)
}
