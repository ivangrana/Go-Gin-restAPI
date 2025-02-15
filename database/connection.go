package database

import (
	"Phinance/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbURL := "postgres://" + dbUser + ":" + dbPassword + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + dbName + "?sslmode=disable"

	db, _ := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	// Migration of tables to database

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Budget{})
	db.AutoMigrate(&models.Goals{})
	db.AutoMigrate(&models.Categories{})
	db.AutoMigrate(&models.Transactions{})
	db.AutoMigrate(&models.MarketProduct{})
	return db
}
