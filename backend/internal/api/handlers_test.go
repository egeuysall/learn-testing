// handlers_test.go
package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/egeuysall/learn-testing/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestSignupHandler_Success(t *testing.T) {
	// 1. Setup database
	pool := repository.SetupTestDB(t)
	defer pool.Close()

	// 2. Create handler
	handler := SignupHandler(pool)

	// 3. Create request body
	reqBody := SignupRequest{
		Email: "test@example.com",
		Name:  "Test User",
	}
	bodyBytes, _ := json.Marshal(reqBody)

	// 4. Create HTTP request
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	// 5. Create response recorder (captures the response)
	w := httptest.NewRecorder()

	// 6. Call the handler
	handler.ServeHTTP(w, req)

	// 7. Assert status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// 8. Parse and assert response body
	var resp SignupResponse
	json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "test@example.com", resp.Email)
	assert.Equal(t, "Test User", resp.Name)
	assert.Greater(t, resp.ID, int32(0))
}

func TestSignupHandler_InvalidJSON(t *testing.T) {
	// 1. Setup database
	pool := repository.SetupTestDB(t)
	defer pool.Close()

	// 2. Create handler
	handler := SignupHandler(pool)

	// 3. Create request body
	reqBody := []byte("invalid json")

	// 4. Create HTTP request
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// 5. Create response recorder (captures the response)
	w := httptest.NewRecorder()

	// 6. Call the handler
	handler.ServeHTTP(w, req)

	// 7. Assert status code
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSignupHandler_DuplicateEmail(t *testing.T) {
	pool := repository.SetupTestDB(t)
	defer pool.Close()

	handler := SignupHandler(pool)

	reqBody := SignupRequest{
		Email: "test@example.com",
		Name:  "Test User",
	}
	bodyBytes1, _ := json.Marshal(reqBody)
	bodyBytes2, _ := json.Marshal(reqBody)

	req1 := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(bodyBytes1))
	req1.Header.Set("Content-Type", "application/json")

	w1 := httptest.NewRecorder()
	handler.ServeHTTP(w1, req1)

	assert.Equal(t, http.StatusCreated, w1.Code)

	req2 := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(bodyBytes2))
	req2.Header.Set("Content-Type", "application/json")

	w2 := httptest.NewRecorder()
	handler.ServeHTTP(w2, req2)

	assert.Equal(t, http.StatusInternalServerError, w2.Code)
}
