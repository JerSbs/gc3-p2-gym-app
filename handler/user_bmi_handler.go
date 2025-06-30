package handler

import (
	"fmt"
	"net/http"

	"gc3-p2-gym-app-JerSbs/service"

	"github.com/labstack/echo/v4"
)

// GetUserBMIHandler godoc
// @Summary      Get BMI from 3rd party API
// @Description  Calculate BMI using user's weight and height from DB, call external API
// @Tags         Users
// @Security     BearerAuth
// @Success      200 {object} dto.BMIResponse
// @Failure      401 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /api/users/bmi [get]
func GetUserBMIHandler(c echo.Context) error {
	userIDInterface := c.Get("user_id") // ✅ Match the key used in middleware
	fmt.Println("📥 Handler Received user_id:", userIDInterface)

	userID, ok := userIDInterface.(uint)
	if !ok {
		fmt.Println("❌ user_id type assertion failed:", userIDInterface)
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Unauthorized or invalid token",
		})
	}

	bmi, err := service.GetUserBMIService(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, bmi)
}
