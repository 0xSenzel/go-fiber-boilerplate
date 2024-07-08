package middlewares

import (
	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/models"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth(user *models.UserRequestDto) fiber.Handler {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		user,
	})
	signedToken, err := token.SignedString(token)
	if err!= nil {
		return fiber.
	}
	return signedToken
}