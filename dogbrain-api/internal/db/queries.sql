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

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1 LIMIT 1;

-- name: CreatePasswordResetToken :one
UPDATE users 
SET verification_token = $1,
    token_expiry = $2,
    updated_at = NOW()
WHERE email = $3 
  AND verified_at IS NOT NULL
RETURNING id;

-- name: ResetPassword :one
UPDATE users 
SET password = $1,
    verification_token = NULL,
    token_expiry = NULL,
    updated_at = NOW()
WHERE verification_token = $2 
  AND token_expiry > NOW()
  AND verified_at IS NOT NULL
RETURNING id;

-- name: CreateSession :exec
INSERT INTO sessions (
    id,
    user_id,
    data,
    expires_at,
    created_at
) VALUES (
    $1, $2, $3, $4, NOW()
);

-- name: GetSession :one
SELECT data 
FROM sessions 
WHERE id = $1 AND expires_at > NOW();

-- name: UpdateSession :exec
UPDATE sessions 
SET data = $2,
    expires_at = $3
WHERE id = $1;

-- name: DeleteSession :exec
DELETE FROM sessions 
WHERE id = $1;

-- name: DeleteUserSessions :exec
DELETE FROM sessions 
WHERE user_id = $1;

-- name: DeleteExpiredSessions :exec
DELETE FROM sessions 
WHERE expires_at < NOW();