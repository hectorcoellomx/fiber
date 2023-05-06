package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hectorcoellomx/fiber/models"
	"gorm.io/gorm"
)

func IndexUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var users []models.User
		if err := db.Find(&users).Error; err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(users)
	}
}

func ShowUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id := c.Params("id")

		var user models.User
		if err := db.First(&user, "id = ?", id).Error; err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(user)
	}
}

func StoreUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var user models.User

		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": err.Error()})
		}

		id := c.FormValue("id")
		name := c.FormValue("name")
		email := c.FormValue("email")

		userset := models.User{Id: id, Name: name, Email: email, Status: 1}

		if err := db.Create(&userset).Error; err != nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": "1"})
	}
}
