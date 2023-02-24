package person

import (
	// "github.com/gin-gonic/gin"
	"persons/model"
)

type PersonsInterface interface {
	FindEveryone(string, string) ([]model.Persons, error)
	FindPerson(string) (model.Persons, error)
	CreatePerson(model.PersonRequest) error
	UpdatePerson(string, model.PersonRequest) error
	DeletePerson(string) error
	FetchPersonCreds(string) (model.PersonAuthReq, error)
}
