package server

import (
	"dogbrain-api/internal/db"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	*fiber.App
	DB *db.DB
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "dogbrain-api",
			AppName:      "dogbrain-api",
		}),
		DB: db.NewDB(),
	}

	return server
}
