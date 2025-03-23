-- name: CreateCloth :one
INSERT INTO clothes (id, name, category, color, size, brand, location, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetClothByID :one
SELECT * FROM clothes WHERE id = $1;

-- name: ListClothes :many
SELECT * FROM clothes;

-- name: DeleteCloth :exec
DELETE FROM clothes WHERE id = $1;