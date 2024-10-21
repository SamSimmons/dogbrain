-- name: CreateUser :one
INSERT INTO users (
    id, email, password, created_at, updated_at, verification_token
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: VerifyUser :one
UPDATE users
SET verified_at = $2
WHERE verification_token = $1 AND verified_at IS NULL
RETURNING *;