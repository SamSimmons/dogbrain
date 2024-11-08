package server

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"dogbrain-api/internal/db"

	"github.com/google/uuid"
)

type SessionStorage struct {
	queries db.Querier
}

func NewSessionStorage(queries db.Querier) *SessionStorage {
	return &SessionStorage{queries: queries}
}

// Get retrieves the session data for a given key
func (s *SessionStorage) Get(key string) ([]byte, error) {
	data, err := s.queries.GetSession(context.Background(), key)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get session: %w", err)
	}

	// Decode the base64 data
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode session data: %w", err)
	}

	return decoded, nil
}

// Set stores session data for a given key with expiration
func (s *SessionStorage) Set(key string, val []byte, exp time.Duration) error {
	// Extract UUID from gob-encoded data
	var userID uuid.NullUUID
	if strings.Contains(string(val), "user_id") {
		parts := strings.Split(string(val), "$")
		for _, part := range parts {
			if id, err := uuid.Parse(strings.TrimSpace(part)); err == nil {
				userID = uuid.NullUUID{
					UUID:  id,
					Valid: true,
				}
				break
			}
		}
	}

	// Encode the binary data as base64
	encodedData := base64.StdEncoding.EncodeToString(val)

	err := s.queries.CreateSession(context.Background(), db.CreateSessionParams{
		ID:        key,
		UserID:    userID,
		Data:      encodedData,
		ExpiresAt: time.Now().Add(exp),
	})

	if err != nil {
		err = s.queries.UpdateSession(context.Background(), db.UpdateSessionParams{
			ID:        key,
			Data:      encodedData,
			ExpiresAt: time.Now().Add(exp),
		})
		if err != nil {
			return fmt.Errorf("failed to update session: %w", err)
		}
	}

	return nil
}

// Delete removes a session by key
func (s *SessionStorage) Delete(key string) error {
	fmt.Printf("Deleting session with key: %s\n", key)
	err := s.queries.DeleteSession(context.Background(), key)
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}
	fmt.Printf("Session deleted successfully\n")
	return nil
}

// Reset removes all expired sessions
func (s *SessionStorage) Reset() error {
	return s.queries.DeleteExpiredSessions(context.Background())
}

// Close implements the Storage interface (no-op in our case)
func (s *SessionStorage) Close() error {
	return nil
}

// DeleteUserSessions removes all sessions for a specific user
func (s *SessionStorage) DeleteUserSessions(userID uuid.UUID) error {
	return s.queries.DeleteUserSessions(context.Background(), uuid.NullUUID{
		UUID:  userID,
		Valid: true,
	})
}
