package utils

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
  	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// var host = os.Getenv("DB_HOST")
	// var user = os.Getenv("DB_USER")
	// var name = os.Getenv("DB_NAME")
	// var password = os.Getenv("DB_PASSWORD")
	// var port = os.Getenv("DB_PORT")

	var err error
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, name, port)
	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect with DB!")
	} else {
		log.Println("Connected to DB!")
	}
}