package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadConfigs() {
	curr_path, _ := os.Getwd()
	err := godotenv.Load(curr_path + "/.env")
	if err != nil {
		fmt.Println(err.Error())
	}

	loadDBConfig()
	loadAuthConfig()
}
