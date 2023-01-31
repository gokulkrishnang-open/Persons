package person

import (
	"persons/model"
)

type PersonGormInt interface {
	FetchEveryoneFromDB() ([]model.Persons, error)
	FetchPersonFromdb(string) (model.Persons, error)
	InsertPersonIntodb(model.Persons) error
	UpdatePersonInDB(string, model.Persons) error
	DeletePersonFromDB(string) error
}
