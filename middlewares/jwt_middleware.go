package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte("supersecretkey"),
		ContextKey: "jwt",

		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized atau token tidak valid",
			})
		},

		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("jwt").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)

			
			if uid, ok := claims["user_id"].(float64); ok {
				c.Locals("user_id", uint(uid))
			}

			adminVal := claims["is_admin"]

			switch v := adminVal.(type) {
			case bool:
				c.Locals("is_admin", v)
			case float64:
				c.Locals("is_admin", v == 1)
			case int:
				c.Locals("is_admin", v == 1)
			default:
				c.Locals("is_admin", false)
			}



			return c.Next()
		},
	})
}