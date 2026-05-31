package auth

import (
	"todo_list/src/db"
	"todo_list/src/db/repos"
	"todo_list/src/lib/core"

	"github.com/gofiber/fiber/v3"
)

type Handler struct{}

func (h *Handler) Register(ctx fiber.Ctx) error {
	result, err := db.Q.CreateUser(
		ctx.Context(), repos.CreateUserParams{
			Name:     "Joker",
			Email:    "joker@gmail.com",
			Password: "joker123",
		},
	)
	if err != nil {
		return core.BadRequestError("user not created", "USER_ERR", core.WithCause(err))
	}
	return ctx.JSON(fiber.Map{
		"id":    result.ID,
		"Email": result.Email,
	})
}
