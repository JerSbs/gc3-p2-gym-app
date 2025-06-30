package routes

import (
	"gc3-p2-gym-app-JerSbs/handler"
	"gc3-p2-gym-app-JerSbs/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/api/users")

	userGroup.POST("/register", handler.RegisterUserHandler)
	userGroup.POST("/login", handler.LoginUserHandler)

	userGroup.GET("/profile", handler.GetUserProfile, middleware.JWTMiddleware)
	userGroup.GET("/bmi", handler.GetUserBMIHandler, middleware.JWTMiddleware)

}
