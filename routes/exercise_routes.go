package routes

import (
	"gc3-p2-gym-app-JerSbs/handler"
	"gc3-p2-gym-app-JerSbs/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterExerciseRoutes(e *echo.Echo) {
	ex := e.Group("/api/exercises")
	ex.Use(middleware.JWTMiddleware)

	ex.POST("", handler.CreateExerciseHandler)
	ex.DELETE("/:id", handler.DeleteExerciseHandler)
}
