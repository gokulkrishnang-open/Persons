package controller

import (
	"persons/services/person"
	transform "persons/transformers"
	validator "persons/validators"

	"github.com/gin-gonic/gin"
)

var Person person.PersonsInterface = person.PersonsImplementation{}

func FindEveryoneHandler(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")
	resp, err := Person.FindEveryone(page, limit)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	people := transform.GetPeople(resp)
	ctx.JSON(200, people)
}

func FindPersonHandler(ctx *gin.Context) {
	user_name, er := validator.ValidatePersonRequest(ctx)
	if er != nil {
		return
	}
	resp, err := Person.FindPerson(user_name)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	p := transform.GetPerson(resp)
	ctx.JSON(200, p)
}

func CreatePersonHandler(ctx *gin.Context) {
	person_req, err := validator.ValidateCreatePersonRequest(ctx)
	if err != nil {
		return
	}
	err = Person.CreatePerson(person_req)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
}

func ChangePersonHandler(ctx *gin.Context) {
	user_name, er := validator.ValidatePersonRequest(ctx)
	if er != nil {
		return
	}
	person_req, err := validator.ValidateUpdatePersonRequest(ctx)
	if err != nil {
		return
	}
	err = Person.ChangePerson(user_name, person_req)
	if err != nil {
		ctx.JSON(400, err)
		return
	}
}

func DeletePersonHandler(ctx *gin.Context) {
	user_name, er := validator.ValidatePersonRequest(ctx)
	if er != nil {
		return
	}
	err := Person.DeletePerson(user_name)
	if err != nil {
		ctx.JSON(500, err)
	}
}
