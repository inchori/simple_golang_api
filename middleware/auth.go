package middleware

import (
	"errors"
	"fmt"
	"strings"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     jwtware.SigningKey{Key: []byte("secret")},
		SuccessHandler: authentication,
		ErrorHandler:   jwtError,
	})
}

func validateToken(tokenString string) error {
	if tokenString == "" {
		return errors.New("empty token")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Claims.(jwt.MapClaims); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["sub"])
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

func authentication(c *fiber.Ctx) error {
	bearerToken := c.Get("Authorization")
	token := strings.TrimPrefix(bearerToken, "Bearer ")

	if err := validateToken(token); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Next()
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"message": "Missing or malformed JWT"})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"message": "Invalid or expired JWT"})
}
