package handlers

import (
	"fmt"
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

	createdUser, err := h.Repository.Create(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": createdUser.InsertedID, "username": user.Username, "email": user.Email})
}

func (h *UserHandler) FindByID(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)

	user, err := h.Repository.FindByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, user)
}
