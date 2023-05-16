package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hectorcoellomx/fiber/auth"
	"github.com/hectorcoellomx/fiber/config"
	"github.com/hectorcoellomx/fiber/controllers"
	"github.com/hectorcoellomx/fiber/database"
	"github.com/hectorcoellomx/fiber/models"
)

func main() {

	/* db, err := database.OpenDB(config.Config{ Host: "localhost", Port: "3306", User: "root", Password: "", DBName: "fiber" }) */

	db, err := database.OpenDB(config.Config{
		Host:     "204.12.242.103",
		Port:     "1433",
		User:     "",
		Password: "",
		DBName:   "",
	})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})

	app := fiber.New()
	api := app.Group("/api")

	api.Get("/login", auth.GenerateToken)

	api.Get("/users", controllers.IndexUser(db))
	api.Get("/users/:id", controllers.ShowUser(db))
	api.Post("/users", JWTMiddleware, controllers.StoreUser(db))

	app.Listen(":3000")

}

func JWTMiddleware(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")
	validate := auth.ValidateToken(authHeader)

	if validate["success"] == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": validate["message"]})
	}

	return c.Next()
}
