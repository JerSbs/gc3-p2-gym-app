package routes

import (
	"p2-graded-challenge-3-JerSbs/handler"
	"p2-graded-challenge-3-JerSbs/middleware"

	"github.com/labstack/echo/v4"
)

func LogRoutes(e *echo.Echo) {
	log := e.Group("/api/logs")
	log.Use(middleware.JWTMiddleware)

	log.POST("", handler.CreateLogHandler)
	log.GET("", handler.GetUserLogsHandler)
}
