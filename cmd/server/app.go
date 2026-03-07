package server

import (
	"todo_list/boot/conf"
	"todo_list/pkgs/http"

	"github.com/gofiber/fiber/v3"
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
	// Notfound Handler
	app.Use(func(ctx fiber.Ctx) error {
		path := ctx.Path()
		method := ctx.Method()
		// BadRequest Error
		err := http.BadRequestError(
			"Wrong Path",
			"NOT_FOUND",
			http.WithMeta("path", path),
			http.WithMeta("method", method),
		)
		return err.ToJSON(ctx)
	})
	return app
}
