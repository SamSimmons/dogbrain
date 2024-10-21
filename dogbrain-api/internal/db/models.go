// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                uuid.UUID    `json:"id"`
	Email             string       `json:"email"`
	Password          string       `json:"password"`
	CreatedAt         time.Time    `json:"created_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
	VerificationToken string       `json:"verification_token"`
	VerifiedAt        sql.NullTime `json:"verified_at"`
}