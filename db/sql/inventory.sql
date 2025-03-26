-- name: CreateItem :one
INSERT INTO inventory (id, owner_id, name, category, color, size, brand, location, created_at, updated_at) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: GetItemByID :one
SELECT * FROM inventory WHERE id = $1 AND owner_id = $2;

-- name: ListItems :many
SELECT * FROM inventory WHERE owner_id = $1;

-- name: DeleteItem :exec
DELETE FROM inventory WHERE id = $1 AND owner_id = $2;

-- name: UpdateItemLocation :exec
UPDATE inventory SET location = $2 WHERE id = $1 AND owner_id = $3;
