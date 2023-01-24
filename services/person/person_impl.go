package person

import (
	// "fmt"
	// "github.com/gin-gonic/gin"
	"persons/model"
	"persons/repositories/persons_gorm"
)

type PersonsImplementation struct{}

var PersonGorm person.PersonGormInt = person.PersonGormImp{}

func (p PersonsImplementation) FindEveryone() ([]model.PersonResponse, error) {
	allpersons, err := PersonGorm.FetchEveryoneFromDB()
	if err != nil {
		return allpersons, err
	}
	return allpersons, nil
}

func (p PersonsImplementation) FindPerson(user_name string) (model.PersonResponse, error) {
	PersonReq := model.PersonRequest{user_name}
	someone, err := PersonGorm.FetchPersonFromdb(PersonReq)
	if err != nil {
		return someone, err
	}
	return someone, nil
}

func (p PersonsImplementation) CreatePerson(person model.Persons) error {
	// person := model.Persons{}
	// ctx.BindJSON(&person)
	err := PersonGorm.InsertPersonIntodb(person)
	if err != nil {
		return err
	}
	return nil
}

func (p PersonsImplementation) ChangePerson(user_name string, person model.Persons) error {
	PersonReq := model.PersonRequest{user_name}
	// person := model.Persons{}
	// ctx.BindJSON(&person)
	err := PersonGorm.UpdatePersonInDB(PersonReq, person)
	if err != nil {
		return err
	}
	return nil
}

func (p PersonsImplementation) DeletePerson(user_name string) error {
	PersonReq := model.PersonRequest{user_name}
	err := PersonGorm.DeletePersonFromDB(PersonReq)
	if err != nil {
		return err
	}
	return nil
}
