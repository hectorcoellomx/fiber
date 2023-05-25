package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/hectorcoellomx/fiber/auth"
	"github.com/hectorcoellomx/fiber/config"
	"github.com/hectorcoellomx/fiber/controllers"
	"github.com/hectorcoellomx/fiber/database"
	"github.com/hectorcoellomx/fiber/models"
<<<<<<< HEAD
	"github.com/joho/godotenv"
=======
>>>>>>> 22f881b84c961f07505a670c837817ed369bef93
)

func main() {

<<<<<<< HEAD
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	db, err := database.OpenDB(config.Config{
		Host:     os.Getenv("DB_SERVER"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DATABASE"),
=======
	/* db, err := database.OpenDB(config.Config{ Host: "localhost", Port: "3306", User: "root", Password: "", DBName: "fiber" }) */

	db, err := database.OpenDB(config.Config{
		Host:     "204.12.242.103",
		Port:     "1433",
		User:     "",
		Password: "",
		DBName:   "",
>>>>>>> 22f881b84c961f07505a670c837817ed369bef93
	})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})

	app := fiber.New()
	api := app.Group("/api")

	api.Get("/login", controllers.LoginUser())
	api.Get("/refresh-token", controllers.RefreshToken())

	api.Get("/users", controllers.IndexUser(db))
	api.Get("/users/:id", controllers.ShowUser(db))
	api.Post("/users", JWTMiddleware, controllers.StoreUser(db))
	api.Put("/users/:id", JWTMiddleware, controllers.UpdateUser(db))
	api.Delete("/users/:id", JWTMiddleware, controllers.DestroyUser(db))

	app.Listen(":8080")

}

func JWTMiddleware(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")
<<<<<<< HEAD
	validate := auth.ValidateToken(c, authHeader)
	//claims := validate["claims"].(map[string]interface{})
	//id := claims["sub"].(string)

	if validate["success"] == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"success": false, "message": validate["message"], "error_code": validate["errorno"]})
=======
	validate := auth.ValidateToken(authHeader)

	if validate["success"] == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": validate["message"]})
>>>>>>> 22f881b84c961f07505a670c837817ed369bef93
	}

	return c.Next()
}
