-- name: AddProduct :one
INSERT INTO products (name, description, price, quantity, image) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products WHERE id = $1 LIMIT 1;