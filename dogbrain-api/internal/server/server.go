package server

import (
	"dogbrain-api/internal/db"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
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
	dbConn := db.NewDB()
	queries := db.New(dbConn.DB)

	sessionStorage := NewSessionStorage(queries)

	sessionStore := session.New(session.Config{
		Storage:        sessionStorage,
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
		DB:       dbConn,
		Emails:   NewEmailService(postmarkToken, accountToken, fromEmail),
		Sessions: sessionStore,
	}

	return server
}
