package routes

import (
	"gc3-p2-gym-app-JerSbs/handler"
	"gc3-p2-gym-app-JerSbs/middleware"

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
