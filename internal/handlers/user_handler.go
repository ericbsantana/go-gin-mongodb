package handlers

import (
	"net/http"
	"rest-api/internal/dtos"
	"rest-api/internal/models"
	"rest-api/internal/repositories"
	"rest-api/internal/validator"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Repository *repositories.UserRepository
}

func UserHandlerFromRepository(repo *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		Repository: repo,
	}
}

func (h *UserHandler) Find(c *gin.Context) {
	users, err := h.Repository.Find()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Create(c *gin.Context) {
	var createUserDTO dtos.CreateUserDTO

	messages, err := validator.ParseAndValidateDTO(c, &createUserDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if messages != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": messages})
		return
	}

	_, err = h.Repository.FindByEmail(createUserDTO.Email)

	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with email already exists"})
		return
	}

	user := models.User{
		Username: createUserDTO.Username,
		Email:    createUserDTO.Email,
	}

	_, err = h.Repository.Create(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
