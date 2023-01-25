package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"persons/config"
	"persons/model"
)

var DB *gorm.DB

func Connection() {
	fmt.Println("connecting to database")
	// dsn := "host=localhost user=postgres password='' dbname=persons_api port=5432 sslmode=disable"
	dsn := "host=" + config.Db.Host +
		" user=" + config.Db.Username +
		" password=''" + config.Db.Password +
		" dbname=" + config.Db.Database +
		" port=" + config.Db.Port +
		" sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	DB.AutoMigrate(&model.Persons{})
}
