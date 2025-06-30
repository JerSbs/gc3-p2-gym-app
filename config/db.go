package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gc3-p2-gym-app-JerSbs/models"
)

var DB *gorm.DB

func InitDB() {
	// Load .env (optional in production)
	_ = godotenv.Load()

	// Get Heroku-style ENV
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	// Use sslmode=require for Heroku
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta",
		host, user, password, dbname, port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	log.Println("Database connection established")

	err = db.AutoMigrate(
		&models.User{},
		&models.Workout{},
		&models.Exercise{},
		&models.Log{},
	)
	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}
}

func GetDB() *gorm.DB {
	return DB
}
