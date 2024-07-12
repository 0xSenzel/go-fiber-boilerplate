package middlewares

import (
	"log"
	"strings"
	"time"

	"github.com/0xsenzel/go-fiber-boilerplate/internal/services/user/tables"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuth(c fiber.Ctx) error {
	clientToken := c.Get("Authorization")
	if clientToken == "" {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "No Authorization header provided",
		})
	}
	
	// remove "Bearer" prefix if exists
	splitToken := strings.Split(clientToken, "Bearer ")
	if len(splitToken) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token format",
		})
	}
	clientToken = splitToken[1]
	
	token, err := jwt.Parse(clientToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &fiber.Error{ Message: "Invalid authorization token"}
		}
		return []byte("secret"), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// store claim in local context
		c.Locals("userClaims", claims)
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid token claims",
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