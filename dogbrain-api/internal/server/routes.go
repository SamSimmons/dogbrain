package server

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"

	"dogbrain-api/internal/db"
)

const (
	argonTime    = 1
	argonMemory  = 64 * 1024
	argonThreads = 4
	argonKeyLen  = 32
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)
	s.App.Post("/register", s.registerUser)
	s.App.Get("/verify/:token", s.verifyEmail)
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func (s *FiberServer) registerUser(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Generate a random salt
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error generating salt"})
	}

	// Hash the password using Argon2id
	hash := argon2.IDKey([]byte(input.Password), salt, argonTime, argonMemory, argonThreads, argonKeyLen)

	// Combine salt and hash, and encode as base64
	encodedHash := base64.RawStdEncoding.EncodeToString(append(salt, hash...))

	verificationToken := uuid.New().String()

	user, err := s.DB.CreateUser(context.Background(), db.CreateUserParams{
		ID:                uuid.New(),
		Email:             input.Email,
		Password:          encodedHash,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		VerificationToken: verificationToken,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating user"})
	}
	fmt.Println(user)

	// TODO: Implement email sending logic
	// sendVerificationEmail(user.Email, verificationToken)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully. Please check your email to verify your account."})
}

func (s *FiberServer) verifyEmail(c *fiber.Ctx) error {
	token := c.Params("token")

	_, err := s.DB.VerifyUser(context.Background(), db.VerifyUserParams{
		VerificationToken: token,
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