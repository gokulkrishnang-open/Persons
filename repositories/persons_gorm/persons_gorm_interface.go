package person

import "persons/model"

type PersonGormInt interface {
	FetchEveryoneFromDB(model.Pagination) ([]model.Persons, error)
	FetchPersonFromDB(string) (model.Persons, error)
	InsertPersonIntoDB(model.Persons) error
	UpdatePersonInDB(string, model.Persons) error
	DeletePersonFromDB(string) error
}
