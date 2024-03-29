package config_test

import (
	"net/http"
	"rest-api/internal/handlers"
	"rest-api/internal/repositories"

	"github.com/gin-gonic/gin"
)

func SetupTestRouter() *gin.Engine {
	r := gin.Default()

	userRepository := repositories.UserRepositoryFromDatabase(TestDBInstance)
	userHandler := handlers.UserHandlerFromRepository(userRepository)

	r.GET("/users", userHandler.Find)
	r.POST("/users", userHandler.Create)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	return r
}
