package person

import (
	"persons/model"
)

type PersonGormInt interface {
	FetchEveryoneFromDB() []model.PersonResponse
	FetchPersonFromdb(model.PersonRequest) model.PersonResponse
	InsertPersonIntodb(model.Persons)
	UpdatePersonInDB(model.PersonRequest, model.Persons)
	DeletePersonFromDB(model.PersonRequest)
}
