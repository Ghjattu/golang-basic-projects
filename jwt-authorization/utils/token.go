package utils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

var key string

func init() {
	key := os.Getenv("SECRET_KEY")
	if key == "" {
		key = "secret-key"
	}
}

func GenerateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(key))
}

func ValidateToken(c *fiber.Ctx) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(key), nil
	})

	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *fiber.Ctx) string {
	// from query string parameter
	token := c.Query("token")
	if token != "" {
		return token
	}

	// from Header Authorization field
	bearerToken := c.GetReqHeaders()["Authorization"]
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
