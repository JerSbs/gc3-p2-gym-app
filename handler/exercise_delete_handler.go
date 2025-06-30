package handler

import (
	"net/http"
	"strconv"

	"gc3-p2-gym-app-JerSbs/service"

	"github.com/labstack/echo/v4"
)

// DeleteExerciseHandler godoc
// @Summary Delete exercise
// @Description Delete an exercise and its logs (only if user owns it)
// @Tags exercises
// @Security BearerAuth
// @Produce json
// @Param id path int true "Exercise ID"
// @Success 200 {object} map[string]interface{}
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/exercises/{id} [delete]
func DeleteExerciseHandler(c echo.Context) error {
	exerciseIDParam := c.Param("id")
	exerciseID, err := strconv.Atoi(exerciseIDParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid exercise ID"})
	}

	userID := c.Get("user_id").(uint)

	result, err := service.DeleteExerciseService(userID, uint(exerciseID))
	if err != nil {
		return service.HandleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "proses penghapusan data exercise berhasil",
		"data":    result,
	})
}
