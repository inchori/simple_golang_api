package handler

import (
	"context"
	"grpc_identity/dto"
	"grpc_identity/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NewUserHandler(app fiber.Router, ctx context.Context, userService service.IUserService) {
	app.Post("/signIn", CreateUser(ctx, userService))
	app.Get("/:id", GetUserByID(ctx, userService))
	app.Get("/:name", GetUserByName(ctx, userService))
	app.Delete("/:id", DeleteUserByID(ctx, userService))
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

func GetUserByID(ctx context.Context, userService service.IUserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParams := c.Params("id")
		id, err := strconv.Atoi(idParams)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			userByID, err := userService.GetUserByID(ctx, id)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			} else {
				return c.JSON(userByID)
			}
		}
	}
}

func GetUserByName(ctx context.Context, userService service.IUserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")
		userByName, err := userService.GetUserByName(ctx, name)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			return c.JSON(userByName)
		}
	}
}

func DeleteUserByID(ctx context.Context, userService service.IUserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idParams := c.Params("id")
		id, err := strconv.Atoi(idParams)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			err := userService.DeleteByID(ctx, id)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			} else {
				return c.JSON(fiber.Map{
					"message": "successfully delete user",
				})
			}
		}
	}
}
