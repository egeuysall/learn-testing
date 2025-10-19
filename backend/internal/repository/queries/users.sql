-- name: CreateUser :one
INSERT INTO users (email, name)
VALUES ($1, $2)
RETURNING id, email, name, created_at;

-- name: GetUserByEmail :one
SELECT id, email, name, created_at
FROM users
WHERE email = $1;
