package person

import (
	"persons/model"
)

type PersonGormInt interface {
	FetchEveryoneFromDB() ([]model.PersonResponse, error)
	FetchPersonFromdb(string) (model.PersonResponse, error)
	InsertPersonIntodb(model.Persons) error
	UpdatePersonInDB(string, model.Persons) error
	DeletePersonFromDB(string) error
}
