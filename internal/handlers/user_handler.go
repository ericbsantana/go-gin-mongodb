package handlers

import (
	"errors"
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

func ParseAndValidateDTO(c *gin.Context, dto interface{}) ([]string, error) {
	if c.Request.ContentLength == 0 {
		return nil, errors.New("request body cannot be empty")
	}

	if err := c.ShouldBindJSON(dto); err != nil {
		return nil, err
	}

	if err := validator.GetValidator().Struct(dto); err != nil {
		validationErrorMessages := validator.GetValidationErrorMessages(err)

		return validationErrorMessages, nil
	}

	return nil, nil
}

func (h *UserHandler) Create(c *gin.Context) {
	var createUserDTO dtos.CreateUserDTO

	messages, err := ParseAndValidateDTO(c, &createUserDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if messages != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": messages})
		return
	}

	user := models.User{
		Username: createUserDTO.Username,
		Email:    createUserDTO.Email,
	}

	_, rerr := h.Repository.Create(user)

	if rerr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
