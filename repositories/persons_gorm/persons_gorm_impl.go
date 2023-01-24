package person

import (
	"persons/database"
	"persons/model"
)

type PersonGormImp struct{}

func (p PersonGormImp) FetchEveryoneFromDB() []model.PersonResponse {
	allpersons := []model.Persons{}
	allpersonstoreturn := []model.PersonResponse{}
	db.DB.Model(&allpersons).Find(&allpersonstoreturn)

	return allpersonstoreturn
}

func (p PersonGormImp) FetchPersonFromdb(Person_to_find model.PersonRequest) model.PersonResponse {
	user_name := Person_to_find.UserName
	someone := model.Persons{}
	someonetoreturn := model.PersonResponse{}
	db.DB.Model(&someone).Where("user_name = ?", user_name).Find(&someonetoreturn)

	return someonetoreturn
}

func (p PersonGormImp) InsertPersonIntodb(person model.Persons) {
	db.DB.Create(&person)
}

func (p PersonGormImp) UpdatePersonInDB(Person_to_change model.PersonRequest, person model.Persons) {
	user_name := Person_to_change.UserName
	someone := model.Persons{}
	db.DB.Model(&someone).Where("user_name = ?", user_name).Updates(person)
}

func (p PersonGormImp) DeletePersonFromDB(Person_to_delete model.PersonRequest) {
	user_name := Person_to_delete.UserName
	person := model.Persons{}
	db.DB.Where("user_name = ?", user_name).Delete(&person)
}
