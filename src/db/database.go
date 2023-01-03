package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	dsn := "host=localhost user=postgres password=root dbname=pet-project"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	}

	err = DB.AutoMigrate(&Pet{})

	if err != nil {
		panic("Failed to migrate Pet table")
	}

	fmt.Println("Connection opened to db")
}
