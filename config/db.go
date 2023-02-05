package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	Connection string `json:"DB_CONNECTION"`
	Host       string `json:"DB_HOST"`
	Port       string `json:"DB_PORT"`
	Database   string `json:"DB_DATABASE"`
	Username   string `json:"DB_USERNAME"`
	Password   string `json:"DB_PASSWORD"`
}

var Db *DBConfig

func loadDBConfig() {
	Db = &DBConfig{}
	err := envconfig.Process("db", Db)
	if err != nil {
		fmt.Println(err.Error())
	}
}
