package middleware

import (
	"go-fiber-svelte/internal/lang"
	"go-fiber-svelte/internal/lib"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("token")
		if token == "" {
			authHeader := c.Get("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				token = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": lang.T.Get().UNAUTHORIZED_ACCESS,
			})
		}
		claims, err := lib.Jwt.Verify(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": lang.T.Get().UNAUTHORIZED_ACCESS,
			})
		}
		userId := int(claims["user_id"].(float64))
		c.Locals("user_id", userId)
		return c.Next()
	}
}
