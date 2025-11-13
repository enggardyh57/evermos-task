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

			if admin, ok := claims["is_admin"].(bool); ok {
				c.Locals("is_admin", admin)
			} else {
				c.Locals("is_admin", false) 
			}



			return c.Next()
		},
	})
}