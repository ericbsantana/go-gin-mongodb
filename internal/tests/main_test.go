package tests

import (
	"net/http"
	"net/http/httptest"
	config_test "rest-api/internal/tests/config"
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
