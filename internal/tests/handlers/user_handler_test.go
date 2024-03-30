package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"go-gin-mongo/internal/models"
	config_test "go-gin-mongo/internal/tests/config"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMain(m *testing.M) {
	log.Println("setup is running")
	testDB := config_test.SetupTestDB()
	config_test.TestDBInstance = testDB.Database
	exitVal := m.Run()
	log.Println("teardown is running")
	_ = testDB.Container.Terminate(context.Background())
	os.Exit(exitVal)
}

func TestUserHandler(t *testing.T) {
	t.Cleanup(func() {
		config_test.ClearDB()
	})

	router := config_test.SetupTestRouter()

	t.Run("Create user", func(t *testing.T) {
		t.Run("should create user", func(t *testing.T) {
			defer config_test.ClearDB()
			w := httptest.NewRecorder()
			jsonStr := []byte(`{"username": "oppenheimer", "email": "oppenheimer@example.com"}`)
			req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
			router.ServeHTTP(w, req)

			assert.Equal(t, 201, w.Code)
		})

		t.Run("should not allow to create user with existing email", func(t *testing.T) {
			defer config_test.ClearDB()

			w1 := httptest.NewRecorder()
			jsonStr := []byte(`{"username": "oppenheimer", "email": "oppenheimer@example.com"}`)
			req1, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
			router.ServeHTTP(w1, req1)

			w2 := httptest.NewRecorder()
			req2, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
			router.ServeHTTP(w2, req2)

			assert.Equal(t, 400, w2.Code)
		})

		t.Run("should not allow to create user with incorrect dto", func(t *testing.T) {
			defer config_test.ClearDB()

			w := httptest.NewRecorder()
			jsonStr := []byte(`{"username": "oppenheimer"}`)
			req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
			router.ServeHTTP(w, req)

			assert.Equal(t, 400, w.Code)
		})
	})

	t.Run("Get user", func(t *testing.T) {
		t.Run("should get user", func(t *testing.T) {
			defer config_test.ClearDB()

			config_test.SeedTestDatabase()
			w2 := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/users/6607077651565dc6fbb91859", nil)
			router.ServeHTTP(w2, req)

			assert.Equal(t, 200, w2.Code)
			assert.JSONEq(t, `{"_id":"6607077651565dc6fbb91859","username":"oppenheimer","email":"oppenheimer@example.com"}`, w2.Body.String())
		})

		t.Run("should show user not found when user does not exist", func(t *testing.T) {
			defer config_test.ClearDB()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/users/"+primitive.NewObjectID().Hex(), nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, 404, w.Code)
		})
	})

	t.Run("Find users", func(t *testing.T) {
		config_test.SeedTestDatabase()
		defer config_test.ClearDB()

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/users", nil)
		router.ServeHTTP(w, req)

		var response []models.User
		err := json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, 200, w.Code)
		assert.NoError(t, err)
		assert.Len(t, response, 1)
	})

	t.Run("Update user", func(t *testing.T) {

		t.Run("should update user", func(t *testing.T) {
			config_test.SeedTestDatabase()
			defer config_test.ClearDB()

			w := httptest.NewRecorder()
			jsonStr := []byte(`{"username": "oppenheimer", "email": "abc@abc.com"}`)

			req, _ := http.NewRequest("PATCH", "/users/6607077651565dc6fbb91859", bytes.NewBuffer(jsonStr))
			router.ServeHTTP(w, req)

			assert.Equal(t, 200, w.Code)
			assert.JSONEq(t, `{"_id":"6607077651565dc6fbb91859","username":"oppenheimer","email":"abc@abc.com"}`, w.Body.String())
		})

		t.Run("should not update user that does not exist", func(t *testing.T) {
			config_test.SeedTestDatabase()
			defer config_test.ClearDB()

			w := httptest.NewRecorder()
			jsonStr := []byte(`{"username": "mathematics", "email": "abc@abc.com"}`)

			req, _ := http.NewRequest("PATCH", "/users/"+primitive.NewObjectID().Hex(), bytes.NewBuffer(jsonStr))
			router.ServeHTTP(w, req)

			assert.Equal(t, 404, w.Code)
		})

		t.Run("should not update user with incorrect payload", func(t *testing.T) {
			config_test.SeedTestDatabase()
			defer config_test.ClearDB()

			w := httptest.NewRecorder()

			req, _ := http.NewRequest("PATCH", "/users/6607077651565dc6fbb91859", nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, 400, w.Code)
		})
	})

	t.Run("Delete user", func(t *testing.T) {
		t.Run("should delete user", func(t *testing.T) {
			config_test.SeedTestDatabase()
			defer config_test.ClearDB()

			w1 := httptest.NewRecorder()
			req1, _ := http.NewRequest("DELETE", "/users/6607077651565dc6fbb91859", nil)
			router.ServeHTTP(w1, req1)

			w2 := httptest.NewRecorder()
			req2, _ := http.NewRequest("GET", "/users/6607077651565dc6fbb91859", nil)
			router.ServeHTTP(w2, req2)

			assert.Equal(t, 204, w1.Code)
			assert.Equal(t, 404, w2.Code)
		})

		t.Run("should not delete user that does not exist", func(t *testing.T) {
			config_test.SeedTestDatabase()
			defer config_test.ClearDB()

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", "/users/"+primitive.NewObjectID().Hex(), nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, 404, w.Code)
		})
	})
}
