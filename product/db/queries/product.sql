-- name: AddProduct :one
INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products WHERE id = $1 LIMIT 1;