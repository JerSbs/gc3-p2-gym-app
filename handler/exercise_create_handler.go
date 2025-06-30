package handler

import (
	"net/http"

	"p2-graded-challenge-3-JerSbs/dto"
	"p2-graded-challenge-3-JerSbs/service"

	"github.com/labstack/echo/v4"
)

// CreateExerciseHandler godoc
// @Summary Create exercise
// @Description Create a new exercise under a workout (must own the workout)
// @Tags exercises
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.ExerciseCreateRequest true "Exercise create payload"
// @Success 201 {object} dto.ExerciseResponse
// @Failure 400 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/exercises [post]
func CreateExerciseHandler(c echo.Context) error {
	var input dto.ExerciseCreateRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request"})
	}

	userID := c.Get("user_id").(uint)

	result, err := service.CreateExerciseService(userID, input)
	if err != nil {
		return service.HandleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "exercise created successfully",
		"data":    result,
	})
}
