// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const checkUserExists = `-- name: CheckUserExists :one
SELECT EXISTS(
    SELECT 1 
    FROM users 
    WHERE LOWER(email) = LOWER($1)
) AS exists
`

func (q *Queries) CheckUserExists(ctx context.Context, lower string) (bool, error) {
	row := q.queryRow(ctx, q.checkUserExistsStmt, checkUserExists, lower)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createPasswordResetToken = `-- name: CreatePasswordResetToken :one
UPDATE users 
SET verification_token = $1,
    token_expiry = $2,
    updated_at = NOW()
WHERE email = $3 
  AND verified_at IS NOT NULL
RETURNING id
`

type CreatePasswordResetTokenParams struct {
	VerificationToken sql.NullString `json:"verification_token"`
	TokenExpiry       sql.NullTime   `json:"token_expiry"`
	Email             string         `json:"email"`
}

func (q *Queries) CreatePasswordResetToken(ctx context.Context, arg CreatePasswordResetTokenParams) (uuid.UUID, error) {
	row := q.queryRow(ctx, q.createPasswordResetTokenStmt, createPasswordResetToken, arg.VerificationToken, arg.TokenExpiry, arg.Email)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const createUser = `-- name: CreateUser :one
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
) RETURNING id, email, password, created_at, updated_at, verification_token, verified_at, token_expiry
`

type CreateUserParams struct {
	ID                uuid.UUID      `json:"id"`
	Email             string         `json:"email"`
	Password          string         `json:"password"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	VerificationToken sql.NullString `json:"verification_token"`
	TokenExpiry       sql.NullTime   `json:"token_expiry"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.ID,
		arg.Email,
		arg.Password,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.VerificationToken,
		arg.TokenExpiry,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.VerificationToken,
		&i.VerifiedAt,
		&i.TokenExpiry,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password, created_at, updated_at, verification_token, verified_at, token_expiry FROM users WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.VerificationToken,
		&i.VerifiedAt,
		&i.TokenExpiry,
	)
	return i, err
}

const resetPassword = `-- name: ResetPassword :one
UPDATE users 
SET password = $1,
    verification_token = NULL,
    token_expiry = NULL,
    updated_at = NOW()
WHERE verification_token = $2 
  AND token_expiry > NOW()
  AND verified_at IS NOT NULL
RETURNING id
`

type ResetPasswordParams struct {
	Password          string         `json:"password"`
	VerificationToken sql.NullString `json:"verification_token"`
}

func (q *Queries) ResetPassword(ctx context.Context, arg ResetPasswordParams) (uuid.UUID, error) {
	row := q.queryRow(ctx, q.resetPasswordStmt, resetPassword, arg.Password, arg.VerificationToken)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const verifyUser = `-- name: VerifyUser :one
UPDATE users 
SET 
    verified_at = $1,
    verification_token = NULL,
    token_expiry = NULL
WHERE 
    verification_token = $2 
    AND (token_expiry > $3 OR token_expiry IS NULL)
    AND verified_at IS NULL
RETURNING id, email, password, created_at, updated_at, verification_token, verified_at, token_expiry
`

type VerifyUserParams struct {
	VerifiedAt        sql.NullTime   `json:"verified_at"`
	VerificationToken sql.NullString `json:"verification_token"`
	Now               sql.NullTime   `json:"now"`
}

func (q *Queries) VerifyUser(ctx context.Context, arg VerifyUserParams) (User, error) {
	row := q.queryRow(ctx, q.verifyUserStmt, verifyUser, arg.VerifiedAt, arg.VerificationToken, arg.Now)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.VerificationToken,
		&i.VerifiedAt,
		&i.TokenExpiry,
	)
	return i, err
}
