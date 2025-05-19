package utils

import (
	"fmt"
	"log"
	"os"

	"coupon_system/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DB_URI")

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting to Database", err)
	}

	if err := DB.AutoMigrate(&models.Users{}); err != nil {
		log.Fatalln("Failed to migrate user model to the db", err)
	}

	if err := DB.AutoMigrate(&models.Medicine{}); err != nil {
		log.Fatalln("Failed to migrate medicine model to the db", err)
	}

	if err := DB.AutoMigrate(&models.Coupon{}); err != nil {
		log.Fatalln("Failed to migrate coupon model to the db", err)
	}

	fmt.Println("DB Connected Successfully😎")
}
