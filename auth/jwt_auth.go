package auth

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func GenerateToken(c *fiber.Ctx, id string, email string, role string, duration float64, typetoken ...string) string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	secretKey := os.Getenv("SECRET_KEY")

	if len(typetoken) > 0 && typetoken[0] == "refresh" {
		secretKey = os.Getenv("SECRET_KEY_REFRESH")
	}

	durationInHours := time.Duration(duration * float64(time.Hour))
	userAgent := c.Get("User-Agent")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = id
	claims["email"] = email
	claims["role"] = role
	claims["aud"] = userAgent
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(durationInHours).Unix()

	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return err.Error()
	}

	return tokenString
}

func ValidateToken(c *fiber.Ctx, tokenString string, typetoken ...string) map[string]interface{} {
	res := make(map[string]interface{})

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	secretKey := os.Getenv("SECRET_KEY")

	if len(typetoken) > 0 && typetoken[0] == "refresh" {
		secretKey = os.Getenv("SECRET_KEY_REFRESH")
	}

	userAgent := c.Get("User-Agent")

	if tokenString == "" {
		res["success"] = false
		res["message"] = "Authentication token not provided"
		res["errorno"] = 0
		res["claims"] = nil
		return res
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {

		message := ""
		errorno := 1

		if err != nil {
			message = err.Error()
		}

		if message != "Token is expired" {
			message = "Invalid token"
			errorno = 2
		}

		res["success"] = false
		res["message"] = message
		res["errorno"] = errorno
		res["claims"] = nil
		return res
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		log.Fatal(err.Error())
	}

	aud := claims["aud"].(string)

	if aud != userAgent {
		res["success"] = false
		res["message"] = "Invalid token"
		res["errorno"] = 3
		res["claims"] = nil
		return res
	}

	claims_map := make(map[string]interface{})
	claims_map["sub"] = claims["sub"].(string)
	claims_map["email"] = claims["email"].(string)
	claims_map["role"] = claims["role"].(string)
	claims_map["iat"] = claims["iat"].(float64)

	res["success"] = true
	res["message"] = "ok"
	res["errorno"] = nil
	res["claims"] = claims_map

	return res
}
