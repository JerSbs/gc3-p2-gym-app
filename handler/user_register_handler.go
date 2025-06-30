package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gc3-p2-gym-app-JerSbs/dto"
	"gc3-p2-gym-app-JerSbs/service"
)

// RegisterUserHandler godoc
// @Summary Register a new user
// @Description Create an account with full name, email, password, age
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "User registration payload"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /api/users/register [post]
func RegisterUserHandler(c echo.Context) error {
	var input dto.RegisterRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "invalid request body"})
	}

	user, err := service.RegisterUser(input)
	if err != nil {
		if err == service.ErrEmailAlreadyExists {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "internal server error"})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "user registered successfully",
		"data":    user,
	})
}
