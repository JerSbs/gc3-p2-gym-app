package routes

import (
	"p2-graded-challenge-3-JerSbs/handler"
	"p2-graded-challenge-3-JerSbs/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterWorkoutRoutes(e *echo.Echo) {
	workoutGroup := e.Group("/api/workouts")

	// Apply JWT middleware to all /workouts routes
	workoutGroup.Use(middleware.JWTMiddleware)

	workoutGroup.GET("", handler.GetAllWorkoutsHandler)
	workoutGroup.GET("/:id", handler.GetWorkoutByIDHandler)
	workoutGroup.POST("", handler.CreateWorkoutHandler)
	workoutGroup.PUT("/:id", handler.UpdateWorkoutHandler)
	workoutGroup.DELETE("/:id", handler.DeleteWorkoutHandler)

}
