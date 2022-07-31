package model

// Import
import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

// ConnectToDb init connection
func ConnectToDb(){
	// Load .env file with configs
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Set connection variables
	DbDriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	// Init connection
	DB, err = gorm.Open(DbDriver, DbUrl)

	// Test connection and show error or success
	if err != nil {
		fmt.Println("Cannot connect to database ", DbDriver)
		log.Fatal("Connection error:", err)
	} else {
		fmt.Println("Established database connection to ", DbDriver)
	}

	// Run migrations
	DB.AutoMigrate(&User{})
}