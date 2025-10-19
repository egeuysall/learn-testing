// handlers_with_email_test.go
package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/egeuysall/learn-testing/internal/repository"
	"github.com/egeuysall/learn-testing/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestSignupHandlerWithEmail_Success(t *testing.T) {
	pool := repository.SetupTestDB(t)
	defer pool.Close()

	// Create mock email sender
	mockEmail := &services.MockEmailSender{}

	// Create handler with mock
	handler := SignupHandlerWithEmail(pool, mockEmail)

	// Create request
	reqBody := SignupRequest{
		Email: "test@example.com",
		Name:  "Test User",
	}
	bodyBytes, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Call handler
	handler.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusCreated, w.Code)

	// Assert email was sent (THIS IS THE MAGIC!)
	assert.True(t, mockEmail.SendWelcomeEmailCalled)
	assert.Equal(t, "test@example.com", mockEmail.SentTo)
	assert.Equal(t, "Test User", mockEmail.SentName)
}

func TestSignupHandlerWithEmail_EmailFails(t *testing.T) {
	pool := repository.SetupTestDB(t)
	defer pool.Close()

	// Create mock email sender
	mockEmail := &services.MockEmailSender{
		ShouldReturnError: true,
	}

	// Create handler with mock
	handler := SignupHandlerWithEmail(pool, mockEmail)

	// Create request
	reqBody := SignupRequest{
		Email: "test@example.com",
		Name:  "Test User",
	}
	bodyBytes, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Call handler
	handler.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusCreated, w.Code)

	assert.True(t, mockEmail.SendWelcomeEmailCalled)
	assert.Equal(t, "test@example.com", mockEmail.SentTo)
	assert.Equal(t, "Test User", mockEmail.SentName)
}
