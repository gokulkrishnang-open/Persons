package person

import (
	"persons/model"
	person "persons/repositories/persons_gorm"
	transformer "persons/transformers"
)

type PersonsImplementation struct{}

var PersonGorm person.PersonGormInt = person.PersonGormImp{}

func (p PersonsImplementation) FindEveryone(pageNum string, limit string) ([]model.Persons, error) {
	paginator := model.Pagination{
		ResultsPerPage: limit,
		Page:           pageNum,
	}
	allpersons, err := PersonGorm.FetchEveryoneFromDB(paginator)
	// allpersonsresp := transformer.GetPeople(allpersons)
	if err != nil {
		return allpersons, err
	}
	return allpersons, nil
}

func (p PersonsImplementation) FindPerson(user_name string) (model.Persons, error) {
	someone, err := PersonGorm.FetchPersonFromDB(user_name)
	// someoneresp := transformer.GetPerson(someone)
	if err != nil {
		return someone, err
	}
	return someone, nil
}

func (p PersonsImplementation) CreatePerson(person_req model.PersonRequest) error {
	person_ := transformer.GetPersonEntity(person_req)
	err := PersonGorm.InsertPersonIntoDB(person_)
	if err != nil {
		return err
	}
	return nil
}

func (p PersonsImplementation) ChangePerson(user_name string, person_req model.PersonRequest) error {
	person_ := transformer.GetPersonEntity(person_req)
	err := PersonGorm.UpdatePersonInDB(user_name, person_)
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

func (p PersonsImplementation) FetchPersonCreds(user_name string) (model.PersonAuthReq, error) {
	user, err := PersonGorm.FetchPersonFromDB(user_name)
	userCreds := transformer.GetPersonAuthReq(user)
	if err != nil {
		return userCreds, err
	}
	return userCreds, nil
}
