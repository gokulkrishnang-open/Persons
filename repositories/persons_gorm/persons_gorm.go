package person

import (
	db "persons/database"
	"persons/model"
	"strconv"

	"gorm.io/gorm"
)

type PersonGormImp struct{}

func (p PersonGormImp) FetchEveryoneFromDB(pag model.Pagination) ([]model.Persons, error) {
	limit, _ := strconv.Atoi(pag.ResultsPerPage)
	pageNoo, _ := strconv.Atoi(pag.Page)
	offset := (pageNoo - 1) * limit

	AllPersons := []model.Persons{}
	err := db.DB.Session(&gorm.Session{}).
		Model(&AllPersons).
		Where("deleted_at IS NULL").
		Limit(limit).
		Offset(offset).
		Find(&AllPersons).Error
	return AllPersons, err
}

func (p PersonGormImp) FetchPersonFromDB(user_name string) (model.Persons, error) {
	SomeOne := model.Persons{}
	err := db.DB.Session(&gorm.Session{}).
		Model(&SomeOne).
		Where("user_name = ? and deleted_at IS NULL", user_name).
		First(&SomeOne).Error
	return SomeOne, err
}

func (p PersonGormImp) InsertPersonIntoDB(person model.Persons) error {
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
	err := db.DB.Session(&gorm.Session{}).
		Where("user_name = ?", user_name).
		Delete(&person).Error
	return err
}
