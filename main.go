package main
import (

	"github.com/gin-gonic/gin"
)
func setupRouter() *gin.Engine {
	r := gin.Default()
}
func main() {
	r := setupRouter()

	r.Run(":8080")
}
