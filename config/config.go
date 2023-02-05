package config

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func LoadConfigs() {
	// curr_path, _ := os.Getwd()
	_, b, _, _ := runtime.Caller(0)

	// Root folder of this project
	ProjectRootPath := filepath.Join(filepath.Dir(b), "../")
	err := godotenv.Load(ProjectRootPath + "/.env")
	if err != nil {
		fmt.Println(err.Error())
	}

	loadDBConfig()
	loadJwtConfig()
	loadAuthConfig()
	loadAppConfig()
}
