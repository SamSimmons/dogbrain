package server

import (
	"dogbrain-api/internal/db"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
)

type FiberServer struct {
	*fiber.App
	DB       *db.DB
	Emails   *EmailService
	Sessions *session.Store
}

const (
	msgInvalidCredentials   = "Invalid credentials"
	msgServerError          = "An error occurred"
	msgLoginSuccess         = "Login successful"
	msgLogoutSuccess        = "Logged out successfully"
	msgUnverifiedEmail      = "Please verify your email before logging in"
	msgInvalidRequestFormat = "Invalid request format"
	msgResetEmailSent       = "If your email is valid, you will receive a password reset link shortly"
	msgPasswordReset        = "Password has been reset successfully"
)

func New(postmarkToken, accountToken, fromEmail string) *FiberServer {
	storage := postgres.New(postgres.Config{
		ConnectionURI: os.Getenv("DATABASE_URL"),
		Reset:         false,
	})

	sessionStore := session.New(session.Config{
		Storage:        storage,
		Expiration:     30 * 24 * time.Hour,
		KeyLookup:      "cookie:session",
		CookieHTTPOnly: true,
		CookieSecure:   true,
		CookieSameSite: "Lax",
		CookiePath:     "/",
	})

	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "dogbrain-api",
			AppName:      "dogbrain-api",
		}),
		DB:       db.NewDB(),
		Emails:   NewEmailService(postmarkToken, accountToken, fromEmail),
		Sessions: sessionStore,
	}

	return server
}
