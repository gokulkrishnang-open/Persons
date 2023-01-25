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

func (p PersonGormImp) FetchPersonFromdb(user_name string) (model.PersonResponse, error) {
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

func (p PersonGormImp) UpdatePersonInDB(user_name string, person model.Persons) error {
	SomeOne := model.Persons{}
	err := db.DB.Session(&gorm.Session{}).
		Model(&SomeOne).
		Where("user_name = ?", user_name).
		Updates(person).Error
	return err
}

func (p PersonGormImp) DeletePersonFromDB(user_name string) error {
	person := model.Persons{}
	err := db.DB.Where("user_name = ?", user_name).
		Delete(&person).Error
	return err
}
