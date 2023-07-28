-- name: CraeateTokens :one
INSERT INTO tokens (jwt_token, refresh_token) 
VALUES ($1, $2)
RETURNING *; 
