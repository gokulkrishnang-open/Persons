package person

import (
	"persons/model"
)

type PersonGormInt interface {
	FetchEveryoneFromDB() ([]model.PersonResponse, error)
	FetchPersonFromdb(model.PersonRequest) (model.PersonResponse, error)
	InsertPersonIntodb(model.Persons) error
	UpdatePersonInDB(model.PersonRequest, model.Persons) error
	DeletePersonFromDB(model.PersonRequest) error
}
