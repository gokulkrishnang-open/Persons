package person

import (
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
	// PersonReq := model.PersonRequest{user_name, ""}
	someone, err := PersonGorm.FetchPersonFromdb(user_name)
	if err != nil {
		return someone, err
	}
	return someone, nil
}

func (p PersonsImplementation) CreatePerson(person model.Persons) error {
	err := PersonGorm.InsertPersonIntodb(person)
	if err != nil {
		return err
	}
	return nil
}

func (p PersonsImplementation) ChangePerson(user_name string, person model.Persons) error {
	// PersonReq := model.PersonRequest{user_name, ""}
	err := PersonGorm.UpdatePersonInDB(user_name, person)
	if err != nil {
		return err
	}
	return nil
}

func (p PersonsImplementation) DeletePerson(user_name string) error {
	// PersonReq := model.PersonRequest{user_name, ""}
	err := PersonGorm.DeletePersonFromDB(user_name)
	if err != nil {
		return err
	}
	return nil
}
