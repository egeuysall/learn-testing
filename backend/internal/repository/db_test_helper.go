package repository

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func SetupTestDB(t *testing.T) *pgxpool.Pool {
	ctx := context.Background()

	// Start PostgreSQL container
	pgContainer, err := postgres.Run(ctx,
		"docker.io/postgres:16-alpine",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpass"),
	)
	if err != nil {
		t.Fatal(err)
	}

	// Clean up container when test ends
	t.Cleanup(func() {
		pgContainer.Terminate(ctx)
	})

	// Get connection string with sslmode parameter
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Connection string: %s", connStr)

	var pool *pgxpool.Pool
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		pool, err = pgxpool.New(ctx, connStr)
		if err == nil {
			// Test the connection
			err = pool.Ping(ctx)
			if err == nil {
				break
			}
			pool.Close()
		}
		if i < maxRetries-1 {
			t.Logf("Connection attempt %d failed, retrying in 1 second...", i+1)
			time.Sleep(1 * time.Second)
		}
	}
	if err != nil {
		t.Fatalf("failed to connect to database after %d attempts: %v", maxRetries, err)
	}

	// Run migrations
	runMigrations(t, pool)

	return pool
}

func runMigrations(t *testing.T, pool *pgxpool.Pool) {
	ctx := context.Background()

	// Actual SQL to create the table
	migration := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) UNIQUE NOT NULL,
		name VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	);
	`

	_, err := pool.Exec(ctx, migration)
	if err != nil {
		t.Fatal(err)
	}
}
