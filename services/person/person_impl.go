package person

import (
	"github.com/gin-gonic/gin"
	"persons/model"
	"persons/repositories/persons_gorm"
)

type PersonsImplementation struct{}

var PersonGorm person.PersonGormInt = person.PersonGormImp{}

func (p PersonsImplementation) FindEveryone(ctx *gin.Context) {
	allpersons := PersonGorm.FetchEveryoneFromDB()
	ctx.JSON(200, allpersons)
}

func (p PersonsImplementation) FindPerson(user_name string, ctx *gin.Context) {
	PersonReq := model.PersonRequest{user_name}
	someone := PersonGorm.FetchPersonFromdb(PersonReq)
	ctx.JSON(200, someone)
}

func (p PersonsImplementation) CreatePerson(ctx *gin.Context) {
	person := model.Persons{}
	ctx.BindJSON(&person)
	PersonGorm.InsertPersonIntodb(person)
}

func (p PersonsImplementation) ChangePerson(user_name string, ctx *gin.Context) {
	PersonReq := model.PersonRequest{user_name}
	person := model.Persons{}
	ctx.BindJSON(&person)
	PersonGorm.UpdatePersonInDB(PersonReq, person)
}

func (p PersonsImplementation) DeletePerson(user_name string) {
	PersonReq := model.PersonRequest{user_name}
	PersonGorm.DeletePersonFromDB(PersonReq)
}
