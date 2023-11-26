package handler

import (
	"context"
	"grpc_identity/dto"
	"grpc_identity/middleware"
	"grpc_identity/service"
	"grpc_identity/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func NewUserHandler(app fiber.Router, ctx context.Context, userService service.IUserService, protected fiber.Handler) {
	app.Post("/signIn", CreateUser(ctx, userService))
	app.Get("/:id", GetUserByID(ctx, userService))
	app.Get("/:name", GetUserByName(ctx, userService))
	app.Delete("/:id", protected, DeleteUserByID(ctx, userService))
	app.Put("/:id", protected, UpdateUser(ctx, userService))
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
				return c.JSON(dto.NewUserResponse(user))
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
				return c.JSON(dto.NewUserResponse(userByID))
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
			return c.JSON(dto.NewUserResponse(userByName))
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
			tokenClaimsID, err := middleware.ExtractTokenMetadata(c)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			if strconv.Itoa(id) != tokenClaimsID {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "unauthorized",
				})
			}

			err = userService.DeleteByID(ctx, id)
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
			tokenClaimsID, err := middleware.ExtractTokenMetadata(c)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			if strconv.Itoa(id) != tokenClaimsID {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "unauthorized",
				})
			}

			userUpdateRequest := new(dto.UserUpdateRequest)
			if err := c.BodyParser(&userUpdateRequest); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			password, err := utils.HashPassword(userUpdateRequest.Password)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			userResponse, err := userService.UpdateUser(ctx, userUpdateRequest.Name, password, id)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			return c.JSON(dto.NewUserResponse(userResponse))
		}
	}
}
