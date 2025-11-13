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

	user := api.Group("/user", middlewares.Protected())
	user.Get("/", handlers.GetUserProfile)
	user.Get("/:id", handlers.GetUserByID)
	user.Put("/update", handlers.UpdateUser)
	user.Delete("/:id", handlers.DeleteUser)

	alamat := api.Group("/alamat", middlewares.Protected())
	alamat.Get("/", handlers.GetAlamat)
	alamat.Post("/", handlers.CreateAlamat)
	alamat.Get("/:id", handlers.GetAlamatByID)
	alamat.Put("/:id", handlers.UpdateAlamat)
	alamat.Delete("/:id", handlers.DeleteAlamat)

	kategori := api.Group("/kategori", middlewares.Protected())
	kategori.Get("/", handlers.GetAllKategori)
	kategori.Post("/", handlers.CreateKategori)
	kategori.Get("/:id", handlers.GetKategoriByID)
	kategori.Put("/:id", handlers.UpdateKategori)
	kategori.Delete("/:id", handlers.DeleteKategori)


	produk := api.Group("/produk", middlewares.Protected())
	produk.Get("/", handlers.GetAllProduk)
	produk.Get("/:id", handlers.GetProdukByID)
	produk.Post("/", handlers.CreateProduk)
	produk.Put("/:id", handlers.UpdateProduk)
	produk.Delete("/:id", handlers.DeleteProduk)


	transaksi := api.Group("/transaksi")
	transaksi.Get("/", handlers.GetAllTransaksi)
	transaksi.Get("/:id", handlers.GetTransaksiByID)
	transaksi.Post("/", handlers.CreateTransaksi)
	transaksi.Put("/:id", handlers.UpdateTransaksi)
	transaksi.Delete("/:id", handlers.DeleteTransaksi)
}