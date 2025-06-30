package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gc3-p2-gym-app-JerSbs/dto"
	"gc3-p2-gym-app-JerSbs/service"
)

// CreateLogHandler godoc
// @Summary Create exercise log
// @Description Create a new log entry for an exercise (must own the exercise)
// @Tags logs
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.LogCreateRequest true "Log create payload"
// @Success 201 {object} dto.LogResponse
// @Failure 400 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /api/logs [post]
func CreateLogHandler(c echo.Context) error {
	var input dto.LogCreateRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid input"})
	}

	userID := c.Get("user_id").(uint)

	result, err := service.CreateLogService(userID, input)
	if err != nil {
		return service.HandleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "log created successfully",
		"data":    result,
	})
}
