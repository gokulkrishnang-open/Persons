package validator

import (
	"github.com/gin-gonic/gin"
	"persons/model"
)

func ReturnAllPersons(ctx *gin.Context, persons []model.PersonResponse) {
	ctx.JSON(200, persons)
}

func ReturnSomePerson(ctx *gin.Context, person model.PersonResponse) {
	ctx.JSON(200, person)
}

func PersonsError(ctx *gin.Context, err error) {
	ctx.JSON(500, gin.H{"error": err.Error()})
}
