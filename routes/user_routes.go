package routes

import (
	"p2-graded-challenge-3-JerSbs/handler"
	"p2-graded-challenge-3-JerSbs/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/api/users")

	userGroup.POST("/register", handler.RegisterUserHandler)
	userGroup.POST("/login", handler.LoginUserHandler)
	userGroup.GET("", handler.GetUserProfile, middleware.JWTMiddleware)
	userGroup.GET("/bmi", handler.GetUserBMIHandler)

}
