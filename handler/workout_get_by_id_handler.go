package handler

import (
	"net/http"
	"strconv"

	"gc3-p2-gym-app-JerSbs/service"

	"github.com/labstack/echo/v4"
)

// GetWorkoutByIDHandler godoc
// @Summary Get workout detail
// @Description Get workout and all its exercises (ownership required)
// @Tags workouts
// @Security BearerAuth
// @Produce json
// @Param id path int true "Workout ID"
// @Success 200 {object} dto.WorkoutDetailResponse
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/workouts/{id} [get]
func GetWorkoutByIDHandler(c echo.Context) error {
	workoutIDParam := c.Param("id")
	workoutID, err := strconv.Atoi(workoutIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid workout ID"})
	}

	userID := c.Get("user_id").(uint)

	workout, err := service.GetWorkoutDetailService(uint(workoutID), userID)
	if err != nil {
		switch err {
		case service.ErrNotFound:
			return c.JSON(http.StatusNotFound, echo.Map{"message": "resource not found"})
		case service.ErrForbidden:
			return c.JSON(http.StatusForbidden, echo.Map{"message": "you do not own this workout"})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "workout detail retrieved",
		"data":    workout,
	})
}
