-- name: AddOrder :one
INSERT INTO orders (product_id, quantity) VALUES ($1, $2) RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders WHERE id = $1 LIMIT 1;