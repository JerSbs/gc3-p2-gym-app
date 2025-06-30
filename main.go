package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	"gc3-p2-gym-app-JerSbs/config"
	"gc3-p2-gym-app-JerSbs/routes"

	_ "gc3-p2-gym-app-JerSbs/docs" // 👉 penting untuk swagger
)

// @title GC3 - Gym App API
// @version 1.0
// @description RESTful API for Gym workout tracking
// @host gc3-p2-gym-app-8a1fe5dad844.herokuapp.com
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

	// Routes
	routes.RegisterUserRoutes(e)
	routes.RegisterWorkoutRoutes(e)
	routes.RegisterExerciseRoutes(e)
	routes.RegisterLogRoutes(e)

	// Swagger endpoint
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Starting server on port:", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatal(err)
	}
}
