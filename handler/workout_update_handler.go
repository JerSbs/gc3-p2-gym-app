package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"gc3-p2-gym-app-JerSbs/dto"
	"gc3-p2-gym-app-JerSbs/service"
)

// UpdateWorkoutHandler godoc
// @Summary Update workout
// @Description Update workout name and description (ownership required)
// @Tags workouts
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Workout ID"
// @Param request body dto.WorkoutUpdateRequest true "Workout update payload"
// @Success 200 {object} dto.WorkoutResponse
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/workouts/{id} [put]
func UpdateWorkoutHandler(c echo.Context) error {
	workoutIDParam := c.Param("id")
	workoutID, err := strconv.Atoi(workoutIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid workout ID"})
	}

	var payload dto.WorkoutUpdateRequest
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request body"})
	}

	userID := c.Get("user_id").(uint)

	result, err := service.UpdateWorkoutService(uint(workoutID), userID, payload)
	if err != nil {
		return service.HandleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Workout updated successfully",
		"data":    result,
	})
}
