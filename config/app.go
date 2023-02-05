package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type AppConfigStruct struct {
	Name        string `json:"APP_NAME"`
	Port        string `json:"APP_PORT"`
	Env         string `json:"APP_ENV"`
	Debug       bool   `json:"APP_DEBUG"`
	Maintenance bool   `json:"APP_MAINTENANCE`
}

var AppConfig *AppConfigStruct

func loadAppConfig() {
	AppConfig = &AppConfigStruct{}
	err := envconfig.Process("app", AppConfig)
	if err != nil {
		fmt.Println(err.Error())
	}
}
