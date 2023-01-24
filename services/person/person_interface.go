package person

import (
	"github.com/gin-gonic/gin"
)

type PersonsInterface interface {
	FindEveryone(*gin.Context)
	FindPerson(string, *gin.Context)
	CreatePerson(*gin.Context)
	ChangePerson(string, *gin.Context)
	DeletePerson(string)
}
