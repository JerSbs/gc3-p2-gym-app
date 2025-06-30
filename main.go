package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"

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
	// Load .env
	config.LoadEnv()

	// Init DB
	config.InitDB()

	// Init Echo
	e := echo.New()

	// Register all routes
	routes.RegisterUserRoutes(e)
	routes.RegisterWorkoutRoutes(e)
	routes.RegisterExerciseRoutes(e)
	routes.RegisterLogRoutes(e)

	// Use Heroku's dynamic PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local
	}

	log.Println("Starting server on port:", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatal(err)
	}
}
