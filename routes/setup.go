package routes

import (
	"github.com/gofiber/fiber/v2"
	"evormos-task/handlers"
	"evormos-task/middlewares"
	
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/register", handlers.Register)
	api.Post("/login", handlers.Login)

	protected := api.Group("/toko", middlewares.Protected())
	protected.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Kamu berhasil mengakses endpoint toko yang dilindungi JWT"})
	})
}