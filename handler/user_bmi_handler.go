package handler

import (
	"net/http"

	"gc3-p2-gym-app-JerSbs/service"

	"github.com/labstack/echo/v4"
)

// GetUserBMIHandler godoc
// @Summary Get authenticated user's BMI
// @Description Retrieves user profile and calculates BMI using 3rd party API
// @Tags Users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} service.UserBMIResult
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/users [get]
func GetUserBMIHandler(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	result, err := service.GetUserBMIService(userID)
	if err != nil {
		if err == service.ErrNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch BMI"})
	}

	return c.JSON(http.StatusOK, result)
}
