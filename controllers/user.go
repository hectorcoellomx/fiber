package controllers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hectorcoellomx/fiber/auth"
	"github.com/hectorcoellomx/fiber/models"
	"gorm.io/gorm"
)

func LoginUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := auth.GenerateToken(c, "1", "hector@gmail.com", "1", 1.0)
		token_refresh := auth.GenerateToken(c, "1", "hector@gmail.com", "1", 168.0, "refresh")

		data := fiber.Map{
			"token":   token,
			"refresh": token_refresh,
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Ok", "data": data})

	}
}

func RefreshToken() fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")
		validate := auth.ValidateToken(c, authHeader, "refresh")

		if validate["success"] == false {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": validate["message"], "error_code": validate["errorno"]})
		}

		claims := validate["claims"].(map[string]interface{})
		id := claims["sub"].(string)
		email := claims["email"].(string)
		role := claims["role"].(string)

		token := auth.GenerateToken(c, id, email, role, 1.0)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Ok", "data": token})
	}
}

func IndexUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var users []models.User
		if err := db.Find(&users).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error(), "error_code": 500})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Ok", "data": users})
	}
}

func ShowUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id := c.Params("id")

		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "message": err.Error(), "error_code": 404})
			} else {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error(), "error_code": 500})
			}
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Ok", "data": user})

	}
}

func StoreUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var user models.User

		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": err.Error(), "error_code": 400})
		}

		id := c.FormValue("id")
		name := c.FormValue("name")
		email := c.FormValue("email")
		statusStr := c.FormValue("status")
		status, err := strconv.Atoi(statusStr)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error(), "error_code": 500})
		}

		userset := models.User{Id: id, Name: name, Email: email, Status: status}

		if err := db.Create(&userset).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error(), "error_code": 500})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Record created", "data": userset})
	}
}

func UpdateUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "message": err.Error(), "error_code": 404})
			} else {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error(), "error_code": 500})
			}
		}

		name := c.FormValue("name")
		email := c.FormValue("email")
		statusStr := c.FormValue("status")
		status, err := strconv.Atoi(statusStr)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error(), "error_code": 500})
		}

		user.Name = name
		user.Email = email
		user.Status = status
		db.Save(&user)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Record updated", "data": user})
	}
}

func DestroyUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"success": false, "message": err.Error(), "error_code": 404})
			} else {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": err.Error(), "error_code": 500})
			}
		}

		db.Delete(&user)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": true, "message": "Record deleted", "data": user})

	}
}
