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

	createdUser, err := h.Repository.Create(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": createdUser.InsertedID, "username": user.Username, "email": user.Email})
}

func (h *UserHandler) FindByID(c *gin.Context) {
	id := c.Param("id")

	user, err := h.Repository.FindByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var updateUserDTO dtos.UpdateUserDTO

	messages, err := validator.ParseAndValidateDTO(c, &updateUserDTO)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if messages != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": messages})
		return
	}

	_, err = h.Repository.FindByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	user := models.User{
		Username: updateUserDTO.Username,
		Email:    updateUserDTO.Email,
	}

	_, err = h.Repository.Update(id, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	updatedUser, _ := h.Repository.FindByID(id)

	c.JSON(http.StatusOK, gin.H{"id": updatedUser.ID, "username": updatedUser.Username, "email": updatedUser.Email})
}

func (h *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	_, err := h.Repository.FindByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{`message`: `User not found`})
		return
	}

	_, err = h.Repository.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
