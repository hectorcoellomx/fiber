package auth

import (
    "time"

    "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

const (
	secretKey = "abc34Ns1t34pp2012s" 
)

func GenerateToken(c *fiber.Ctx) error {

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


func ValidateToken(tokenString string) map[string]interface{} {
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