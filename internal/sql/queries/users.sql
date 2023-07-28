-- name: CreateUser :one
INSERT INTO users (id, full_name, email, nick_name, age, password_hash) 
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;