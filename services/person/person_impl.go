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
	if err != nil {
		return allpersons, err
	}
	return allpersons, nil
}

func (p PersonsImplementation) FindPerson(userName string) (model.Persons, error) {
	someone, err := PersonGorm.FetchPersonFromDB(userName)
	if err != nil {
		return someone, err
	}
	return someone, nil
}

func (p PersonsImplementation) CreatePerson(personReq model.PersonRequest) error {
	personEntity := transformer.GetPersonEntity(personReq)
	err := PersonGorm.InsertPersonIntoDB(personEntity)
	if err != nil {
		return err
	}
	return nil
}

func (p PersonsImplementation) ChangePerson(userName string, personReq model.PersonRequest) error {
	personEntity := transformer.GetPersonEntity(personReq)
	err := PersonGorm.UpdatePersonInDB(userName, personEntity)
	if err != nil {
		return err
	}
	return nil
}

func (p PersonsImplementation) DeletePerson(userName string) error {
	err := PersonGorm.DeletePersonFromDB(userName)
	if err != nil {
		return err
	}
	return nil
}

func (p PersonsImplementation) FetchPersonCreds(userName string) (model.PersonAuthReq, error) {
	user, err := PersonGorm.FetchPersonFromDB(userName)
	userCreds := transformer.GetPersonAuthReq(user)
	if err != nil {
		return userCreds, err
	}
	return userCreds, nil
}
