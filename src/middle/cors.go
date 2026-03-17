package middle

import (
	"todo_list/src/conf"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func CorsOrigin() fiber.Handler {
	return cors.New(cors.Config{
		MaxAge:           86400,
		AllowOrigins:     []string{conf.Env.CORS_ORIGIN},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
	})
}
