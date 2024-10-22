package server

import (
	"dogbrain-api/internal/db"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	*fiber.App
	DB *db.DB
	Emails *EmailService
}

func New(postmarkToken, accountToken, fromEmail string) *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "dogbrain-api",
			AppName:      "dogbrain-api",
		}),
		DB: db.NewDB(),
		Emails: NewEmailService(postmarkToken, accountToken, fromEmail),		
	}

	return server
}
