package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		fmt.Println("🔐 Raw Authorization Header:", authHeader)

		if authHeader == "" {
			fmt.Println("❌ Missing Authorization header")
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "missing authorization header",
			})
		}

		// Split "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			fmt.Println("❌ Invalid Authorization format:", parts)
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "invalid authorization format",
			})
		}

		tokenString := parts[1]
		fmt.Println("🔑 Token extracted:", tokenString)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				fmt.Println("❌ Unexpected signing method")
				return nil, echo.ErrUnauthorized
			}
			secret := os.Getenv("JWT_SECRET")
			fmt.Println("🛡️ JWT_SECRET from env:", secret)
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			fmt.Println("❌ Token parse error or invalid:", err)
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "invalid or expired token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Println("❌ Failed to parse claims")
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "invalid token claims",
			})
		}

		fmt.Printf("📄 Token Claims: %+v\n", claims)

		// Read the "user_id" field
		rawUserID, ok := claims["user_id"].(float64)
		if !ok {
			fmt.Println("❌ user_id missing or invalid type:", claims["user_id"])
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "user_id not found or invalid in token",
			})
		}

		userID := uint(rawUserID)
		fmt.Println("✅ Authenticated user_id:", userID)

		c.Set("user_id", userID)
		return next(c)
	}
}
