package server

import (
	"todo_list/boot/conf"
	"todo_list/pkgs/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      conf.Env.NAME,
		ErrorHandler: http.ErrorHandler,
		BodyLimit:    int(conf.Env.BODY_LIMIT),
	})
	// Logger Middleware
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	// Stack trace only in dev
	app.Use(recover.New(recover.Config{
		EnableStackTrace: conf.Env.IS_DEV,
	}))
	// Secure Header
	app.Use(helmet.New())
	// Cors Origin Middleware
	app.Use(cors.New(cors.Config{
		MaxAge:           86400,
		AllowOrigins:     []string{conf.Env.CORS_ORIGIN},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
	}))
	// Notfound Handler
	app.Use(http.NotFoundHandler)

	return app
}
