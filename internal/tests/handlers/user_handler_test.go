package handlers_test

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	config_test "rest-api/internal/tests/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	log.Println("setup is running")
	testDB := config_test.SetupTestDB()
	config_test.TestDBInstance = testDB.Database
	config_test.SeedTestDatabase()
	exitVal := m.Run()
	log.Println("teardown is running")
	_ = testDB.Container.Terminate(context.Background())
	os.Exit(exitVal)
}

func TestInitializeRouter(t *testing.T) {
	router := config_test.SetupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}

func TestCreateUser(t *testing.T) {
	defer config_test.ClearDB()
	router := config_test.SetupTestRouter()

	w := httptest.NewRecorder()

	jsonStr := []byte(`{"username": "John Doe", "email": "johndoe@example.com"}`)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestCreateUserWithExistingEmail(t *testing.T) {
	router := config_test.SetupTestRouter()

	w1 := httptest.NewRecorder()
	jsonStr := []byte(`{"username": "John Doe", "email": "johndoe@example.com"}`)
	req1, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w1, req1)

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w2, req2)

	assert.Equal(t, 400, w2.Code)
}