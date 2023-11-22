package handler

import (
	"context"
	"grpc_identity/dto"
	"grpc_identity/service"
	"grpc_identity/utils"

	"github.com/gofiber/fiber/v2"
)

func NewUserHandler(app fiber.Router, ctx context.Context, userService service.IUserService, authentication fiber.Handler) {
	app.Post("/signUp", CreateUser(ctx, userService))
	app.Get("/:id", GetUserByID(ctx, userService))
	app.Get("/:name", GetUserByName(ctx, userService))
	app.Delete("/:id", authentication, DeleteUserByID(ctx, userService))
	app.Put("/:id", authentication, UpdateUser(ctx, userService), authentication)
}

func CreateUser(ctx context.Context, userService service.IUserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRequest := new(dto.UserCreateRequest)
		if err := c.BodyParser(&userRequest); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			password, err := utils.HashPassword(userRequest.Password)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			user, err := userService.CreateUser(ctx, userRequest.Name, userRequest.Email, password)
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
		id, err := c.ParamsInt("id")
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
		id, err := c.ParamsInt("id")
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

func UpdateUser(ctx context.Context, userService service.IUserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			userUpdateRequest := new(dto.UserUpdateRequest)
			if err := c.BodyParser(&userUpdateRequest); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			password, err := utils.HashPassword(userUpdateRequest.Password)
			userResponse, err := userService.UpdateUser(ctx, userUpdateRequest.Name, password, id)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			return c.JSON(userResponse)
		}
	}
}
