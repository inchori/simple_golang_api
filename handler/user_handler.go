package handler

import (
	"context"
	"grpc_identity/dto"
	"grpc_identity/service"

	"github.com/gofiber/fiber/v2"
)

func NewUserHandler(app fiber.Router, ctx context.Context, userService service.IUserService) {
	app.Post("/signin", CreateUser(ctx, userService))
}

func CreateUser(ctx context.Context, userService service.IUserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRequest := new(dto.UserRequest)
		if err := c.BodyParser(&userRequest); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			user, err := userService.CreateUser(ctx, userRequest.Name, userRequest.Email, userRequest.Password)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			} else {
				return c.JSON(user)
			}
		}
	}
}
