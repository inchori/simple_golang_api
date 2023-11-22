package handler

import (
	"context"
	"grpc_identity/dto"
	"grpc_identity/service"
	"grpc_identity/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func NewLoginHandler(app fiber.Router, ctx context.Context, userService service.IUserService) {
	app.Post("/login", Login(ctx, userService))
}

func Login(ctx context.Context, userService service.IUserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		loginRequest := new(dto.LoginRequest)
		if err := c.BodyParser(&loginRequest); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			userByEmail, hashPassword, err := userService.GetUserByEmail(ctx, loginRequest.Email)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			if !utils.CheckPasswordHash(loginRequest.Password, hashPassword) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "invalid hashPassword",
				})
			}

			token := jwt.New(jwt.SigningMethodHS256)

			claims := token.Claims.(jwt.MapClaims)
			claims["sub"] = strconv.Itoa(userByEmail.ID)
			claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			return c.JSON(fiber.Map{
				"token": t,
			})
		}
	}
}
