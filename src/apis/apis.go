package apis

import (
	"todo_list/src/apis/auth"
	"todo_list/src/apis/health"

	"github.com/gofiber/fiber/v3"
)

func Router(app fiber.Router) {
	// Health Endpoint
	app.Route("/", health.Router, "Health")
	// Apis Endpoints
	api := app.Group("/api")
	{
		api.Route("/auth", auth.Router, "Auth")
	}
}
