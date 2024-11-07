-- name: CreateUser :one
INSERT INTO users (
    id, 
    email, 
    password, 
    created_at, 
    updated_at, 
    verification_token, 
    token_expiry
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: VerifyUser :one
UPDATE users 
SET 
    verified_at = $1,
    verification_token = NULL,
    token_expiry = NULL
WHERE 
    verification_token = $2 
    AND (token_expiry > sqlc.arg(now) OR token_expiry IS NULL)
    AND verified_at IS NULL
RETURNING *;

-- name: CheckUserExists :one
SELECT EXISTS(
    SELECT 1 
    FROM users 
    WHERE LOWER(email) = LOWER($1)
) AS exists;