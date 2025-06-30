package handler

import (
	"net/http"

	"gc3-p2-gym-app-JerSbs/dto"
	"gc3-p2-gym-app-JerSbs/service"

	"github.com/labstack/echo/v4"
)

// LoginUserHandler godoc
// @Summary Login user
// @Description Authenticate user with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "User login payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/users/login [post]
func LoginUserHandler(c echo.Context) error {
	var input dto.LoginRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request body"})
	}

	token, err := service.LoginUser(input)
	if err != nil {
		switch err {
		case service.ErrInvalidInput:
			return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		case service.ErrUserNotFound, service.ErrWrongPassword:
			return c.JSON(http.StatusUnauthorized, echo.Map{"message": err.Error()})
		default:
			return c.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "login successful",
		"token":   token,
	})
}
