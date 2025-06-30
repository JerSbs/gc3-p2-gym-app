package handler

import (
	"net/http"

	"gc3-p2-gym-app-JerSbs/service"

	"github.com/labstack/echo/v4"
)

// GetUserProfileHandler godoc
// @Summary Get user profile
// @Description Retrieve profile of the authenticated user
// @Tags users
// @Security BearerAuth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /api/users [get]
func GetUserProfile(c echo.Context) error {
	userID := c.Get("user_id").(uint)

	user, err := service.GetUserProfile(userID)
	if err != nil {
		if err == service.ErrNotFound {
			return c.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
	}

	return c.JSON(http.StatusOK, user)
}
