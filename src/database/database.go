package database

import (
	"fmt"

	"github.com/guisantosalves/go-api-fiber/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbURL := "postgres://postgres:123456@localhost:5432/book"

	// (dialector, gormConfigs)
	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected Successfully to the database")

	// run auto migration for given models
	database.AutoMigrate(&models.Book{})

	// pass for DB
	DB = database
}
