-- name: CreateUser :one
INSERT INTO users (id, email, first_name, last_name, password_hash, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;
