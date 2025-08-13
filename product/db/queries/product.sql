-- name: AddProduct :one
INSERT INTO products (name, description, price, quantity) VALUES ($1, $2, $3, $4) RETURNING *;
