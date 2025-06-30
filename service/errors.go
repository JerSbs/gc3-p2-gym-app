package service

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrInvalidInput       = errors.New("invalid input")
	ErrEmailAlreadyExists = errors.New("email already registered")
	ErrUserNotFound       = errors.New("user not found")
	ErrWrongPassword      = errors.New("wrong password")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden: you do not have access")
	ErrNotFound           = errors.New("resource not found")
	ErrInternal           = errors.New("internal server error")
)

func HandleServiceError(c echo.Context, err error) error {
	switch err {
	case ErrInvalidInput:
		return c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	case ErrUnauthorized:
		return c.JSON(http.StatusUnauthorized, echo.Map{"message": err.Error()})
	case ErrForbidden:
		return c.JSON(http.StatusForbidden, echo.Map{"message": err.Error()})
	case ErrNotFound:
		return c.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
	default:
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": ErrInternal.Error()})
	}
}
