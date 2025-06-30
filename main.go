package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"p2-graded-challenge-3-JerSbs/config"
	"p2-graded-challenge-3-JerSbs/routes"
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
	log.Println("Server started at http://localhost:" + port)

	if err := e.Start(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
