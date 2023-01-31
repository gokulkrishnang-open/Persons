package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"persons/services/person"
	transformer "persons/transformers"
	"persons/validators"
)

var Person person.PersonsInterface = person.PersonsImplementation{}

func FindEveryoneHandler(ctx *gin.Context) {
	//validator

	//service
	resp, err := Person.FindEveryone()
	if err != nil {
		validator.PersonsError(ctx, err)
		return
	}

	//transformer
	validator.ReturnAllPersons(ctx, resp)
}

func FindPersonHandler(ctx *gin.Context) {
	user_name, er := validator.ValidatePersonRequest(ctx)
	if er != nil {
		return
	}
	resp, err := Person.FindPerson(user_name)
	if err != nil {
		validator.PersonsError(ctx, err)
		return
	}
	validator.ReturnSomePerson(ctx, resp)
}

func CreatePersonHandler(ctx *gin.Context) {

	person_req, err := validator.ValidateCreatePersonRequest(ctx)
	if err != nil {
		return
	}

	person := transformer.GetPersonEntity(person_req)

	err = Person.CreatePerson(person)
	if err != nil {
		validator.PersonsError(ctx, err)
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
	person := transformer.GetPersonEntity(person_req)

	err = Person.ChangePerson(user_name, person)
	if err != nil {
		validator.PersonsError(ctx, errors.New("Unable to update, something happened!!"))
		return
	}
	if err != nil {
		validator.PersonsError(ctx, err)
	}
}

func DeletePersonHandler(ctx *gin.Context) {
	user_name, er := validator.ValidatePersonRequest(ctx)
	if er != nil {
		return
	}
	err := Person.DeletePerson(user_name)
	if err != nil {
		validator.PersonsError(ctx, err)
	}
}
