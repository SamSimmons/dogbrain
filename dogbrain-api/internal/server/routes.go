package server

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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
	maxEmailLength        = 255
)

func (s *FiberServer) RegisterFiberRoutes() {
	authLimiter := limiter.New(limiter.Config{
		Max:        20,            // 20 requests
		Expiration: 1 * time.Hour, // Per hour
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // Rate limit by IP
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests. Please try again later.",
			})
		},
	})
	v1 := s.App.Group("/api/v1")
	auth := v1.Group("", authLimiter)

	auth.Post("/register", s.registerUser)
	auth.Get("/verify/:token", s.verifyEmail)
	auth.Post("/forgot-password", s.forgotPassword)
	auth.Post("/reset-password", s.resetPassword)

	auth.Post("/login", s.logIn)
	auth.Post("/logout", s.logOut)

	s.Stack()
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

	encodedHash, err := hashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	verificationToken := uuid.New().String()
	tokenExpiry := time.Now().Add(tokenValidityDuration)

	_, err = s.DB.CreateUser(context.Background(), db.CreateUserParams{
		ID:        uuid.New(),
		Email:     normalizedEmail,
		Password:  encodedHash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		VerificationToken: sql.NullString{
			String: verificationToken,
			Valid:  true,
		},
		TokenExpiry: sql.NullTime{
			Time:  tokenExpiry,
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
			String: token,
			Valid:  true,
		},
		VerifiedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Now: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid or expired verification token"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error verifying email"})
	}

	return c.JSON(fiber.Map{"message": "Email verified successfully"})
}

func verifyPassword(storedHash string, password string) (bool, error) {
	hashBytes, err := base64.RawStdEncoding.DecodeString(storedHash)
	if err != nil {
		return false, err
	}

	salt := hashBytes[:saltLen]
	storedHashOnly := hashBytes[saltLen:]

	// Use same Argon2id parameters as registration
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		argonTime,
		argonMemory,
		argonThreads,
		argonKeyLen,
	)

	return subtle.ConstantTimeCompare(hash, storedHashOnly) == 1, nil
}

func hashPassword(password string) (string, error) {
	// Generate salt
	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("error generating salt: %w", err)
	}

	// Hash the password using Argon2id
	hash := argon2.IDKey(
		[]byte(password),
		salt,
		argonTime,
		argonMemory,
		argonThreads,
		argonKeyLen,
	)

	// Encode salt+hash
	return base64.RawStdEncoding.EncodeToString(append(salt, hash...)), nil
}

func (s *FiberServer) forgotPassword(c *fiber.Ctx) error {
	var input struct {
		Email string `json:"email"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": msgInvalidRequestFormat,
		})
	}

	if err := validateEmail(input.Email); err != nil {
		// Still return success to avoid leaking valid emails
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": msgResetEmailSent,
		})
	}

	email := strings.ToLower(strings.TrimSpace(input.Email))

	// Generate reset token
	resetToken := uuid.New().String()
	tokenExpiry := time.Now().Add(1 * time.Hour)

	// Update user with reset token
	_, err := s.DB.CreatePasswordResetToken(context.Background(), db.CreatePasswordResetTokenParams{
		VerificationToken: sql.NullString{
			String: resetToken,
			Valid:  true,
		},
		TokenExpiry: sql.NullTime{
			Time:  tokenExpiry,
			Valid: true,
		},
		Email: email,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			// User not found - return same message to avoid leaking valid emails
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": msgResetEmailSent,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": msgServerError,
		})
	}

	if err := s.Emails.SendPasswordResetEmail(email, resetToken); err != nil {
		// Log the error but don't expose it to the user
		fmt.Printf("Failed to send password reset email: %v\n", err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": msgResetEmailSent,
	})
}

func (s *FiberServer) resetPassword(c *fiber.Ctx) error {
	var input struct {
		Token    string `json:"token"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": msgInvalidRequestFormat,
		})
	}

	if err := validatePassword(input.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	encodedHash, err := hashPassword(input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": msgServerError,
		})
	}

	// Update password and invalidate token
	userID, err := s.DB.ResetPassword(context.Background(), db.ResetPasswordParams{
		Password: encodedHash,
		VerificationToken: sql.NullString{
			String: input.Token,
			Valid:  true,
		},
	})

	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid or expired reset token",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": msgServerError,
		})
	}

	if err := s.Sessions.Storage.(*SessionStorage).DeleteUserSessions(userID); err != nil {
		// Log the error but don't fail the password reset
		fmt.Printf("Failed to delete user sessions during password reset: %v\n", err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": msgPasswordReset,
	})
}

func (s *FiberServer) logIn(c *fiber.Ctx) error {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		if elapsed < time.Second {
			time.Sleep(time.Second - elapsed)
		}
	}()

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": msgInvalidRequestFormat,
		})
	}

	if err := validateEmail(input.Email); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": msgInvalidCredentials,
		})
	}
	if err := validatePassword(input.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": msgInvalidCredentials,
		})
	}

	email := strings.ToLower(strings.TrimSpace(input.Email))

	user, err := s.DB.GetUserByEmail(context.Background(), email)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": msgInvalidCredentials,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": msgServerError,
		})
	}

	valid, err := verifyPassword(user.Password, input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": msgServerError,
		})
	}

	if !valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": msgInvalidCredentials,
		})
	}

	if !user.VerifiedAt.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": msgUnverifiedEmail,
		})
	}

	sess, err := s.Sessions.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": msgServerError,
		})
	}

	sess.Set("user_id", user.ID.String())

	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": msgServerError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": msgLoginSuccess,
	})
}

func (s *FiberServer) logOut(c *fiber.Ctx) error {
	sess, err := s.Sessions.Get(c)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": msgLogoutSuccess,
		})
	}

	if err := sess.Destroy(); err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": msgLogoutSuccess,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": msgLogoutSuccess,
	})
}
