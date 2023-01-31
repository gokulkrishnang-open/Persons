package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"persons/model"
)

var AuthorizedPerson *model.PersonRequest

func loadAuthConfig() {
	AuthorizedPerson = &model.PersonRequest{}
	err := envconfig.Process("default", AuthorizedPerson)
	if err != nil {
		fmt.Println(err.Error())
	}
}
