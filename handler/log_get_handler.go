package handler

import (
	"net/http"

	"gc3-p2-gym-app-JerSbs/service"

	"github.com/labstack/echo/v4"
)

// GetUserLogsHandler godoc
// @Summary Get user logs
// @Description Get all logs created by the authenticated user
// @Tags logs
// @Security BearerAuth
// @Produce json
// @Success 200 {array} dto.LogWithExerciseResponse
// @Failure 401 {object} map[string]string
// @Router /api/logs [get]
func GetUserLogsHandler(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	logs, err := service.GetLogsByUserService(userID)
	if err != nil {
		return service.HandleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "logs fetched successfully",
		"data":    logs,
	})
}
