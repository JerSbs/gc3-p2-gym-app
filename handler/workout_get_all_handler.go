package handler

import (
	"net/http"

	"gc3-p2-gym-app-JerSbs/service"

	"github.com/labstack/echo/v4"
)

// GetAllWorkoutsHandler godoc
// @Summary Get all workouts
// @Description Get all workouts owned by the authenticated user
// @Tags workouts
// @Security BearerAuth
// @Produce json
// @Success 200 {array} dto.WorkoutResponse
// @Failure 401 {object} map[string]string
// @Router /api/workouts [get]
func GetAllWorkoutsHandler(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	workouts, err := service.GetAllWorkoutsService(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "workouts retrieved",
		"data":    workouts,
	})
}
