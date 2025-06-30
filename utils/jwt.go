package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint) (string, error) {
	// Get secret from .env
	secret := os.Getenv("JWT_SECRET")

	// Set claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // token berlaku 24 jam
	}

	// Buat token baru
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign dan return string token
	return token.SignedString([]byte(secret))
}
