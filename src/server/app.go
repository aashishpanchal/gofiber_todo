package server

import (
	"todo_list/src/apis"
	"todo_list/src/conf"
	"todo_list/src/middle"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func New() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      conf.Env.NAME,
		ErrorHandler: middle.ErrorHandler,
		BodyLimit:    conf.Env.BODY_LIMIT,
		// Proxy Config
		ProxyHeader: conf.Env.PROXY_HEADER,
		TrustProxy:  conf.Env.TRUST_PROXY,
	})

	// Stack trace only in dev
	app.Use(recover.New(recover.Config{
		EnableStackTrace: conf.Env.IS_DEV,
	}))
	// Logger Middleware
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	// Secure Header
	app.Use(helmet.New())
	// Rate limiting
	app.Use(middle.RateLimit(middle.RateOption{
		Limit:    conf.Env.RATE_LIMIT_MAX_REQ,
		WindowMs: conf.Env.RATE_LIMIT_WINDOWS,
	}))
	// Cors Origin Middleware
	app.Use(cors.New(cors.Config{
		MaxAge:           86400,
		AllowOrigins:     []string{conf.Env.CORS_ORIGIN},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
	}))

	// Api Router
	app.Route("/", apis.Router, "apis")

	// Notfound Handler
	app.Use(middle.NotFoundHandler)

	return app
}
