package server

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"

	"dogbrain-api/internal/db"
)

const (
	argonTime    = 2
	argonMemory  = 19 * 1024
	argonThreads = 1
	argonKeyLen  = 32

	saltLen = 16

	minPasswordLength = 8
	maxPasswordLength = 128

	tokenValidityDuration = 24 * time.Hour
	maxEmailLength = 255
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Post("/register", s.registerUser)
	s.App.Get("/verify/:token", s.verifyEmail)
}

func validatePassword(password string) error {
	// "Passwords must be at least 8 characters long"
	if len(password) < minPasswordLength {
		return fmt.Errorf("password must be at least %d characters long", minPasswordLength)
	}

	// "Do not set the maximum password length too low. Anywhere around 64-256 characters is a good maximum"
	if len(password) > maxPasswordLength {
		return fmt.Errorf("password must not exceed %d characters", maxPasswordLength)
	}

	// "All valid Unicode characters should be allowed, including whitespace"
	if !utf8.ValidString(password) {
		return fmt.Errorf("password contains invalid characters")
	}

	return nil
}

func validateEmail(email string) error {
	// "Maximum of 255 characters"
	if len(email) > maxEmailLength {
		return fmt.Errorf("email must not exceed %d characters", maxEmailLength)
	}

	// "It does not start or end with a whitespace"
	if strings.TrimSpace(email) != email {
		return fmt.Errorf("email must not start or end with whitespace")
	}

	// "Includes at least 1 @ character"
	atIndex := strings.LastIndex(email, "@")
	if atIndex == -1 {
		return fmt.Errorf("email must contain @")
	}

	// "Has at least 1 character before the @"
	if atIndex == 0 {
		return fmt.Errorf("email must have at least one character before @")
	}

	domain := email[atIndex+1:]
	// "The domain part includes at least 1 . and has at least 1 character before it"
	dotIndex := strings.LastIndex(domain, ".")
	if dotIndex == -1 || dotIndex == 0 {
		return fmt.Errorf("invalid email domain")
	}

	return nil
}

func (s *FiberServer) registerUser(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := validateEmail(input.Email); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := validatePassword(input.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	normalizedEmail := strings.ToLower(input.Email)

	exists, err := s.DB.CheckUserExists(context.Background(), normalizedEmail)
	if err != nil {
		// Database error - don't expose internals to user
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	if exists {
		// User exists, but we send the same message as success to avoid leaking information
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "If your email is valid, you will receive a verification link shortly",
		})
	}

	// Generate a random salt
	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// Hash the password using Argon2id
	hash := argon2.IDKey(
		[]byte(input.Password),
		salt,
		argonTime,
		argonMemory,
		argonThreads,
		argonKeyLen,
	)

	encodedHash := base64.RawStdEncoding.EncodeToString(append(salt, hash...))

	verificationToken := uuid.New().String()
	tokenExpiry := time.Now().Add(tokenValidityDuration)

	_, err = s.DB.CreateUser(context.Background(), db.CreateUserParams{
		ID:                uuid.New(),
		Email:             normalizedEmail,
		Password:          encodedHash,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		VerificationToken: sql.NullString{
			String: verificationToken,
			Valid: true,
		},
		TokenExpiry:			 sql.NullTime{
			Time: tokenExpiry,
			Valid: true,
		},
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	if err := s.Emails.SendVerificationEmail(input.Email, verificationToken); err != nil {
		// Log the error but don't expose it to the user
		fmt.Printf("Failed to send verification email: %v\n", err)
		// Don't return an error to the user to avoid leaking information
}

	// Return the same message whether the email existed or not
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "If your email is valid, you will receive a verification link shortly",
	})
}

func (s *FiberServer) verifyEmail(c *fiber.Ctx) error {
	token := c.Params("token")

	_, err := s.DB.VerifyUser(context.Background(), db.VerifyUserParams{
		VerificationToken: sql.NullString{
			String:token,
			Valid: true,
		},
		VerifiedAt:        sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid or expired verification token"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error verifying email"})
	}

	return c.JSON(fiber.Map{"message": "Email verified successfully"})
}