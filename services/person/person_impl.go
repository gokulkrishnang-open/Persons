package person

import (
	"persons/model"
	"persons/repositories/persons_gorm"
	transformer "persons/transformers"
)

type PersonsImplementation struct{}

var PersonGorm person.PersonGormInt = person.PersonGormImp{}

func (p PersonsImplementation) FindEveryone() ([]model.PersonResponse, error) {
	allpersons, err := PersonGorm.FetchEveryoneFromDB()
	allpersonsresp := transformer.GetPeople(allpersons)
	if err != nil {
		return allpersonsresp, err
	}
	return allpersonsresp, nil
}

func (p PersonsImplementation) FindPerson(user_name string) (model.PersonResponse, error) {
	someone, err := PersonGorm.FetchPersonFromdb(user_name)
	someoneresp := transformer.GetPerson(someone)
	if err != nil {
		return someoneresp, err
	}
	return someoneresp, nil
}

func (p PersonsImplementation) CreatePerson(person model.Persons) error {
	err := PersonGorm.InsertPersonIntodb(person)
	if err != nil {
		return err
	}
	return nil
}

func (p PersonsImplementation) ChangePerson(user_name string, person model.Persons) error {
	err := PersonGorm.UpdatePersonInDB(user_name, person)
	if err != nil {
		return err
	}
	return nil
}

func (p PersonsImplementation) DeletePerson(user_name string) error {
	err := PersonGorm.DeletePersonFromDB(user_name)
	if err != nil {
		return err
	}
	return nil
}
