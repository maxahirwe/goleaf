package initializer

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DATABASE *gorm.DB

// Loads .env environment variables
func LoadEnv() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// Loads the database based on the DB_NAME set in .env
func LoadDB() {
	dbName := os.Getenv("DB_NAME")
	log.Println("DB_NAME:", dbName)
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed connecting to DB_NAME:", dbName)
		panic("Failed to connect database")
	}
	DATABASE = db
}
