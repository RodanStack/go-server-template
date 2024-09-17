-- name: GetUsers :many
SELECT id, username, email, created_at, updated_at, status
FROM users
ORDER BY id ASC LIMIT $1 OFFSET $2;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (username, password, email, status)
VALUES ($1, $2, $3, $4)
RETURNING *;