package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"gc3-p2-gym-app-JerSbs/config"
	"gc3-p2-gym-app-JerSbs/routes"
)

// @title Workout API
// @version 1.0
// @description RESTful API for managing users, workouts, exercises, and logs.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Load ENV (this already calls godotenv.Load())
	config.LoadEnv()

	// Init DB after .env loaded
	config.InitDB()

	// Echo instance
	e := echo.New()

	// Swagger docs
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Route groups
	routes.RegisterUserRoutes(e)
	routes.RegisterWorkoutRoutes(e)
	routes.ExerciseRoutes(e)
	routes.LogRoutes(e)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := e.Start(":" + port); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
