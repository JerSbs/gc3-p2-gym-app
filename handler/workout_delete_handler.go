package handler

import (
	"net/http"
	"strconv"

	"p2-graded-challenge-3-JerSbs/service"

	"github.com/labstack/echo/v4"
)

// DeleteWorkoutHandler godoc
// @Summary Delete workout
// @Description Delete a workout (and all its exercises) if owned by user
// @Tags workouts
// @Security BearerAuth
// @Produce json
// @Param id path int true "Workout ID"
// @Success 200 {object} map[string]interface{}
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/workouts/{id} [delete]
func DeleteWorkoutHandler(c echo.Context) error {
	workoutIDParam := c.Param("id")
	workoutID, err := strconv.Atoi(workoutIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid workout ID"})
	}

	userID := c.Get("user_id").(uint)

	result, err := service.DeleteWorkoutService(uint(workoutID), userID)
	if err != nil {
		return service.HandleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "proses penghapusan data workout berhasil",
		"data":    result,
	})
}
