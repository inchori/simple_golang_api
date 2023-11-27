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

func NewLoginHandler(app fiber.Router, ctx context.Context, userService service.IUserService) {
	app.Post("/login", Login(ctx, userService))
}

// Login user login
//
//	@Summary	User Login
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		request	body		dto.LoginRequest	true	"Login"
//	@Success	200		{object}	dto.LoginResponse
//	@Router		/v1/auth/login [post]
func Login(ctx context.Context, userService service.IUserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		loginRequest := new(dto.LoginRequest)
		if err := c.BodyParser(&loginRequest); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			userByEmail, err := userService.GetUserByEmail(ctx, loginRequest.Email)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			if !utils.CheckPasswordHash(loginRequest.Password, userByEmail.Password) {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "invalid hashPassword",
				})
			}

			jwtToken, err := middleware.CreateAccessToken(strconv.Itoa(userByEmail.ID))
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			return c.JSON(
				dto.NewLoginResponse(jwtToken),
			)
		}
	}
}
