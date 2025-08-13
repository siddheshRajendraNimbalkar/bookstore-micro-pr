-- name: AddOrder :one
INSERT INTO orders (product_id, quantity) VALUES ($1, $2) RETURNING *;

