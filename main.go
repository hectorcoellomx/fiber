package main

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gofiber/fiber/v2"
	"github.com/hectorcoellomx/fiber/config"
	"github.com/hectorcoellomx/fiber/controllers"
	"github.com/hectorcoellomx/fiber/database"
	"github.com/hectorcoellomx/fiber/models"
)

const (
	secretKey = "abc34Ns1t34pp2012s" 
)

func main() {
	
	/* db, err := database.OpenDB(config.Config{ Host: "localhost", Port: "3306", User: "root", Password: "", DBName: "fiber" }) */

	db, err := database.OpenDB(config.Config{
		Host:     "204.12.242.103",
		Port:     "1433",
		User:     "sa",
		Password: "Data4142",
		DBName:   "ZeusTest",
	})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{}) //db.AutoMigrate(&models.User{}, &models.Post{})

	app := fiber.New()
	api := app.Group("/api")

	api.Get("/login", generateToken)

	api.Get("/users", controllers.IndexUser(db))
	api.Get("/users/:id", controllers.ShowUser(db))
	api.Post("/users", protectedRoute, controllers.StoreUser(db))

	app.Listen(":3000")

}

func protectedRoute(c *fiber.Ctx) error {
	
	authHeader := c.Get("Authorization")
	validate := validateToken(authHeader);

	if validate["success"] == false {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{ "message": validate["message"] })
	}

	return c.Next()
}

func generateToken(c *fiber.Ctx) error {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "user_id"         
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() 

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{ "token": tokenString })
}


func validateToken(tokenString string) map[string]interface{} {
    res := make(map[string]interface{})

	if tokenString == "" {
		res["success"] = false
		res["message"] = "No se proporcionó el token de autenticación"
		return res
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		res["success"] = false
		res["message"] = "Token de autenticación inválido"
		return res
	}

	/* 
	
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		res["success"] = false
		res["message"] = "Token de autenticación inválido"
		return res
	}

	expirationTime := claims["exp"].(float64)
	expiration := time.Unix(int64(expirationTime), 0)

	if time.Now().After(expiration) {
		res["success"] = false
		res["message"] = "El token ha expirado"
	}
	
	*/

	res["success"] = true
	res["message"] = "ok"

    return res
}




