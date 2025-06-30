package handler

import (
	"net/http"

	"gc3-p2-gym-app-JerSbs/service"

	"github.com/labstack/echo/v4"
)

// GetUserBMIHandler godoc
// @Summary Get BMI using 3rd party API
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dto.UserBMIResponse
// @Failure 500 {object} map[string]string
// @Router /api/users/bmi [get]
func GetUserBMIHandler(c echo.Context) error {
	userIDStr := c.Get("user_id")
	userID, ok := userIDStr.(uint)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": "unauthorized"})
	}

	data, err := service.GetUserBMIService(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}
