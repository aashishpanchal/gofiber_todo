package apis

import (
	"todo_list/src/apis/health"

	"github.com/gofiber/fiber/v3"
)

func Router(app fiber.Router) {
	// Health Endpoint
	app.Route("/", health.Router, "health")
}
