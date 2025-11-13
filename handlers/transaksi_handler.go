package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"evormos-task/database"
	"evormos-task/models"
)


func GetAllTransaksi(c *fiber.Ctx) error {
	var transaksi []models.Transaksi
	if err := database.DB.Preload("User").Preload("Alamat").Find(&transaksi).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(transaksi)
}


func GetTransaksiByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var trx models.Transaksi

	if err := database.DB.Preload("User").Preload("Alamat").First(&trx, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Transaksi tidak ditemukan"})
	}
	return c.JSON(trx)
}


func CreateTransaksi(c *fiber.Ctx) error {
	var trx models.Transaksi
	if err := c.BodyParser(&trx); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	
	trx.KodeInvoice = fmt.Sprintf("INV-%d", time.Now().Unix())
	trx.CreatedAt = time.Now()
	trx.UpdatedAt = time.Now()

	
	if err := database.DB.Create(&trx).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":  "Gagal membuat transaksi",
			"detail": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Transaksi berhasil dibuat",
		"data":    trx,
	})
}


func UpdateTransaksi(c *fiber.Ctx) error {
	id := c.Params("id")
	var trx models.Transaksi

	if err := database.DB.First(&trx, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Transaksi tidak ditemukan"})
	}

	var input models.Transaksi
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	trx.ID_User = input.ID_User
	trx.ID_Alamat = input.ID_Alamat
	trx.HargaTotal = input.HargaTotal
	trx.MethodBayar = input.MethodBayar
	trx.UpdatedAt = time.Now()

	if err := database.DB.Save(&trx).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Transaksi berhasil diperbarui",
		"data":    trx,
	})
}


func DeleteTransaksi(c *fiber.Ctx) error {
	id := c.Params("id")
	var trx models.Transaksi

	if err := database.DB.First(&trx, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Transaksi tidak ditemukan"})
	}

	if err := database.DB.Delete(&trx).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Transaksi berhasil dihapus"})
}
