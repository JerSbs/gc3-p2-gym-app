package routes

import (
	"gc3-p2-gym-app-JerSbs/handler"
	"gc3-p2-gym-app-JerSbs/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterLogRoutes(e *echo.Echo) {
	log := e.Group("/api/logs")
	log.Use(middleware.JWTMiddleware)

	log.POST("", handler.CreateLogHandler)
	log.GET("", handler.GetUserLogsHandler)
}
