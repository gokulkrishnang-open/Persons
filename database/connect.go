package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"persons/model"
)

var DB *gorm.DB

func Connection() {
	fmt.Println("connecting to database")
	dsn := "host=localhost user=postgres password='' dbname=persons_api port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	DB.AutoMigrate(&model.Persons{})
}
