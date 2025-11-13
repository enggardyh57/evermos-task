package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"evormos-task/database"
	"evormos-task/models"
)


func GetAllProduk(c *fiber.Ctx) error {
	var produk []models.Produk
	if err := database.DB.Preload("Toko").Preload("Kategori").Find(&produk).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data produk"})
	}
	return c.JSON(produk)
}


func GetProdukByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var produk models.Produk
	if err := database.DB.Preload("Toko").Preload("Kategori").First(&produk, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Produk tidak ditemukan"})
	}
	return c.JSON(produk)
}


func CreateProduk(c *fiber.Ctx) error {
	var input models.Produk
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Request tidak valid"})
	}

	input.CreatedAt = time.Now()
	input.UpdatedAt = time.Now()

	if err := database.DB.Create(&input).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menambahkan produk", "detail": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Produk berhasil ditambahkan",
		"produk":  input,
	})
}


func UpdateProduk(c *fiber.Ctx) error {
	id := c.Params("id")
	var produk models.Produk

	if err := database.DB.First(&produk, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Produk tidak ditemukan"})
	}

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Request tidak valid"})
	}

	data["updated_at"] = time.Now()
	if err := database.DB.Model(&produk).Updates(data).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal memperbarui produk"})
	}

	return c.JSON(fiber.Map{
		"message": "Produk berhasil diperbarui",
		"produk":  produk,
	})
}


func DeleteProduk(c *fiber.Ctx) error {
	id := c.Params("id")
	var produk models.Produk

	if err := database.DB.First(&produk, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Produk tidak ditemukan"})
	}

	if err := database.DB.Delete(&produk).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus produk"})
	}

	return c.JSON(fiber.Map{"message": "Produk berhasil dihapus"})
}