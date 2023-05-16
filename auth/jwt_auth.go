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

	duration := 1.0 // hours
	durationInHours := time.Duration(duration * float64(time.Hour))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = "user_id"         
	claims["exp"] = time.Now().Add(durationInHours).Unix() 

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
		res["message"] = "Authentication token not provided"
		return res
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {

		message := ""

		if err != nil {
			message = err.Error()
		}

		if message != "Token is expired" {
			message = "Invalid token"
		}

		res["success"] = false
		res["message"] = message
		return res
	}

	/*claims, ok := token.Claims.(jwt.MapClaims)
	expirationTime := claims["sub"]*/

	res["success"] = true
	res["message"] = "ok"

    return res
}