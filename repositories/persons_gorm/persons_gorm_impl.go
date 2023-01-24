package person

import (
	// "fmt"
	"gorm.io/gorm"
	"persons/database"
	"persons/model"
)

type PersonGormImp struct{}

func (p PersonGormImp) FetchEveryoneFromDB() ([]model.PersonResponse, error) {
	AllPersons := []model.Persons{}
	AllPersonToReturn := []model.PersonResponse{}
	err := db.DB.Session(&gorm.Session{}).
		Model(&AllPersons).
		Where("deleted_at IS NULL").
		Find(&AllPersonToReturn).Error
	return AllPersonToReturn, err
}

func (p PersonGormImp) FetchPersonFromdb(Person_to_find model.PersonRequest) (model.PersonResponse, error) {
	user_name := Person_to_find.UserName
	SomeOne := model.Persons{}
	SomeOneToReturn := model.PersonResponse{}
	err := db.DB.Session(&gorm.Session{}).
		Model(&SomeOne).
		Where("user_name = ? and deleted_at IS NULL", user_name).
		First(&SomeOneToReturn).Error
	return SomeOneToReturn, err
}

func (p PersonGormImp) InsertPersonIntodb(person model.Persons) error {
	err := db.DB.Session(&gorm.Session{}).
		Create(&person).Error
	return err
}

func (p PersonGormImp) UpdatePersonInDB(Person_to_change model.PersonRequest, person model.Persons) error {
	user_name := Person_to_change.UserName
	SomeOne := model.Persons{}
	err := db.DB.Session(&gorm.Session{}).
		Model(&SomeOne).
		Where("user_name = ?", user_name).
		Updates(person).Error
	return err
}

func (p PersonGormImp) DeletePersonFromDB(Person_to_delete model.PersonRequest) error {
	user_name := Person_to_delete.UserName
	person := model.Persons{}
	err := db.DB.Where("user_name = ?", user_name).
		Delete(&person).Error
	return err
}
