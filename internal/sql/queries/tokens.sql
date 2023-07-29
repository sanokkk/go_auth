-- name: CraeateTokens :one
INSERT INTO tokens (expires_at, refresh_token) 
VALUES ($1, $2)
RETURNING *; 



-- name: DeleteToken :exec
DELETE FROM tokens
WHERE refresh_token=$1;
