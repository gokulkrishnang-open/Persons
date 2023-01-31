package person

import (
	// "github.com/gin-gonic/gin"
	"persons/model"
)

type PersonsInterface interface {
	FindEveryone() ([]model.PersonResponse, error)
	FindPerson(string) (model.PersonResponse, error)
	CreatePerson(model.Persons) error
	ChangePerson(string, model.Persons) error
	DeletePerson(string) error
	FetchPersonCreds(string) (model.PersonAuthReq, error)
}
