package repository

import (
	"context"
	"testing"

	generated "github.com/egeuysall/learn-testing/internal/repository/generated"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Setup test database
	pool := SetupTestDB(t)
	defer pool.Close()

	// Create queries instance
	queries := generated.New(pool)
	ctx := context.Background()

	// Test: Create a user
	user, err := queries.CreateUser(ctx, generated.CreateUserParams{
		Email: "test@example.com",
		Name:  "Test User",
	})

	// Assert no error
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "test@example.com", user.Email)
	assert.Equal(t, "Test User", user.Name)
	assert.Greater(t, user.ID, int32(0))
}

func TestGetUserByEmail(t *testing.T) {
	pool := SetupTestDB(t)
	defer pool.Close()

	queries := generated.New(pool)
	ctx := context.Background()

	user, err := queries.CreateUser(ctx, generated.CreateUserParams{
		Email: "john@example.com",
		Name:  "John Doe",
	})

	user, err = queries.GetUserByEmail(ctx, "john@example.com")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "john@example.com", user.Email)
	assert.Equal(t, "John Doe", user.Name)
}
