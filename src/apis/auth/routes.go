package auth

import "github.com/gofiber/fiber/v3"

func Router(app fiber.Router) {
	handler := Handler{}
	// Register Auth Endpoints
	app.Post("/register", handler.Register)
}
