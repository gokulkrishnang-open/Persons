package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type jwtConfig struct {
	Secret string `json:"JWT_SECRETKEY"`
}

var Jwt *jwtConfig

func loadJwtConfig() {
	Jwt = &jwtConfig{}
	err := envconfig.Process("jwt", Jwt)
	if err != nil {
		fmt.Println(err.Error())
	}
}
