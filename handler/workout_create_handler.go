package handler

import (
	"net/http"

	"gc3-p2-gym-app-JerSbs/dto"
	"gc3-p2-gym-app-JerSbs/service"

	"github.com/labstack/echo/v4"
)

// CreateWorkoutHandler godoc
// @Summary Create a new workout
// @Description Create a workout for the authenticated user
// @Tags workouts
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.WorkoutCreateRequest true "Workout create payload"
// @Success 201 {object} dto.WorkoutResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/workouts [post]
func CreateWorkoutHandler(c echo.Context) error {
	var input dto.WorkoutCreateRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request body"})
	}

	userID := c.Get("user_id").(uint)

	workout, err := service.CreateWorkoutService(userID, input)
	if err != nil {
		switch err {
		case service.ErrInvalidInput:
			return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "workout created successfully",
		"data":    workout,
	})
}
