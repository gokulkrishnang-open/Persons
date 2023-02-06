package controller

import (
	"encoding/base64"
	"persons/app/middleware/jwt"
	"persons/model"
	"persons/services/person"
	"strings"

	"github.com/gin-gonic/gin"
)

var Person_ person.PersonsInterface = person.PersonsImplementation{}

func LoginHandler(ctx *gin.Context) {
	encodedThings := strings.SplitN(ctx.GetHeader("Authorization"), " ", 2)[1]
	b, _ := base64.StdEncoding.DecodeString(encodedThings)
	credentials := strings.SplitN(string(b), ":", 2)
	Person := model.PersonAuthReq{credentials[0], credentials[1]}
	ActualPerson, err := Person_.FetchPersonCreds(Person.UserName)

	if Person == ActualPerson {
		token, token_err := jwt.CreateToken(Person)
		if err != nil || token_err != nil {
			ctx.JSON(500, gin.H{"error": "Unable to login"})
			ctx.Abort()
			return
		}
		ctx.JSON(200, gin.H{"token": token})
	} else {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		ctx.Abort()
	}
}
