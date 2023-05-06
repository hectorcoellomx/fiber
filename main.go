package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hectorcoellomx/fiber/config"
	"github.com/hectorcoellomx/fiber/controllers"
	"github.com/hectorcoellomx/fiber/database"
	"github.com/hectorcoellomx/fiber/models"
)

func main() {

	db, err := database.OpenDB(config.Config{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "",
		DBName:   "fiber",
	})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})

	/*
		// Insertar datos de prueba

		users := []models.User{
			{Id: "1", Name: "José Saturnino Cardozo", Email: "pepe@futbol.com", Status: 1},
			{Id: "2", Name: "Antonio Naelson Sinha", Email: "sinha@futbol.com", Status: 1},
		}
		for _, user := range users {
			if err := db.Create(&user).Error; err != nil {
			}
		}

	*/

	app := fiber.New()
	api := app.Group("/api")
	api.Get("/users", controllers.IndexUser(db))
	api.Get("/users/:id", controllers.ShowUser(db))
	api.Post("/users", controllers.StoreUser(db))
	app.Listen(":3000")

}
