package database

import (
	"fmt"

	"github.com/guisantosalves/go-api-fiber/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDB() {

	// change the user and password in another project
	// host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
	dbURL := "host=databasePost user=postgres password=123456 dbname=book port=5432 sslmode=disable"

	// "postgres://postgres:123456@localhost:5432/book"

	// (dialector, gormConfigs)
	database, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connected Successfully to the database")

	// run auto migration for given models
	database.AutoMigrate(&models.Book{})

	// pass for DB
	// DB receive a DbInstance
	DB = DbInstance{
		Db: database,
	}
}
