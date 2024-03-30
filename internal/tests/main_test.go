package tests

import (
	config_test "go-gin-mongo/internal/tests/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplication(t *testing.T) {
	t.Run("should initialize application", func(t *testing.T) {
		router := config_test.SetupTestRouter()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, "OK", w.Body.String())
	})
}
