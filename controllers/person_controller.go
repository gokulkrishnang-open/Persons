package controller

import (
	"github.com/gin-gonic/gin"
	"persons/model"
	"persons/services/person"
)

var Person person.PersonsInterface = person.PersonsImplementation{}

func FindEveryoneHandler(ctx *gin.Context) {
	resp, err := Person.FindEveryone()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, resp)
}

func FindPersonHandler(ctx *gin.Context) {
	user_name := ctx.Params.ByName("username")
	resp, err := Person.FindPerson(user_name)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, resp)
}

func CreatePersonHandler(ctx *gin.Context) {
	person := model.Persons{}
	ctx.BindJSON(&person)
	err := Person.CreatePerson(person)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
}

func ChangePersonHandler(ctx *gin.Context) {
	user_name := ctx.Params.ByName("username")
	person := model.Persons{}
	ctx.BindJSON(&person)
	err := Person.ChangePerson(user_name, person)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
}

func DeletePersonHandler(ctx *gin.Context) {
	user_name := ctx.Params.ByName("username")
	err := Person.DeletePerson(user_name)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
}
