package controller

import (
	"github.com/gin-gonic/gin"
	"persons/services/person"
)

var Person person.PersonsInterface = person.PersonsImplementation{}

func FindEveryoneHandler(ctx *gin.Context) {
	Person.FindEveryone(ctx)
}

func FindPersonHandler(ctx *gin.Context) {
	user_name := ctx.Params.ByName("username")
	Person.FindPerson(user_name, ctx)
}

func CreatePersonHandler(ctx *gin.Context) {
	Person.CreatePerson(ctx)
}

func ChangePersonHandler(ctx *gin.Context) {
	user_name := ctx.Params.ByName("username")
	Person.ChangePerson(user_name, ctx)
}

func DeletePersonHandler(ctx *gin.Context) {
	user_name := ctx.Params.ByName("username")
	Person.DeletePerson(user_name)
}
