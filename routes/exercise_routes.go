package routes

import (
	"p2-graded-challenge-3-JerSbs/handler"
	"p2-graded-challenge-3-JerSbs/middleware"

	"github.com/labstack/echo/v4"
)

func ExerciseRoutes(e *echo.Echo) {
	ex := e.Group("/api/exercises")
	ex.Use(middleware.JWTMiddleware)

	ex.POST("", handler.CreateExerciseHandler)
	ex.DELETE("/:id", handler.DeleteExerciseHandler)
}
