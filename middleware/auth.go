package middleware

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"grpc_identity/dto"
	"strings"
	"time"
)

func NewLoginHandler(app fiber.Router) {
	app.Post("/login", Login())
}

type Authentication struct{}

func NewAuthentication() *Authentication {
	return &Authentication{}
}

func validateToken(tokenString string) error {
	if tokenString == "" {
		return errors.New("empty token")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	} else {
		return err
	}
}

func (a *Authentication) Authentication() fiber.Handler {
	return func(c *fiber.Ctx) error {
		bearerToken := c.Get("Authorization")
		token := strings.TrimPrefix(bearerToken, "Bearer ")

		if err := validateToken(token); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Next()
	}
}

func Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		loginRequest := new(dto.LoginRequest)
		if err := c.BodyParser(&loginRequest); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		} else {
			token := jwt.NewWithClaims(jwt.SigningMethod(jwt.SigningMethodHS256), &jwt.RegisteredClaims{
				Subject:   loginRequest.Email,
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			})

			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": err.Error(),
				})
			}
			return c.JSON(fiber.Map{
				"jwt": t,
			})
		}
	}
}
