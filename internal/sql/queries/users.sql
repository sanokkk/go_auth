-- name: CreateUser :one
INSERT INTO users (id, full_name, email, nick_name, age, password_hash) 
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE nick_name=$1 AND password_hash=$2
LIMIT 1;