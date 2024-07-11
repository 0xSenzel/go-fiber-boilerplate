package middlewares

import (
	"log"
	"time"

	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/tables"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuth() fiber.Handler {
	return func(c fiber.Ctx) error {
		// authHeader := c.Get("Authorization")
		// if authHeader == "" {
		// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		// 		"message": "Missing or malformed JWT",
		// 	})
		// }

		// tokenString := authHeader[len("Bearer ")]
		token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{})
		signedToken, err := token.SignedString(token)
		if err!= nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map {
					"error": err.Error(),
					"message": "Failed to sign access token",
				})
		}

		return c.JSON(fiber.Map{
			"access_token": signedToken,
		})
	}
}

func GenerateToken(user *tables.User) (string, error) {
	claims := jwt.MapClaims{
		"email": user.Email,
		"name": user.Name,
		"id": user.ID,
		"exp": time.Now().Add(time.Minute * 30).Unix(),
	}
	
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))
	if err != nil {
		log.Panic(err)
		return "", err
	}

	return token, nil
}