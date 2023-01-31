package middleware

import (
	// "encoding/base64"
	"github.com/gin-gonic/gin"
	"persons/app/middleware/jwt"
	// "persons/model"
	// "persons/services/person"
	"fmt"
	"strings"
)

// var Person_ person.PersonsInterface = person.PersonsImplementation{}

// func AuthenticateUser(ctx *gin.Context) {
// 	encoded_things := strings.SplitN(ctx.GetHeader("Authorization"), " ", 2)[1]
// 	b, _ := base64.StdEncoding.DecodeString(encoded_things)
// 	credentials := strings.SplitN(string(b), ":", 2)
// 	Person := model.PersonAuthReq{credentials[0], credentials[1]}
// 	ActualPerson, err := Person_.FetchPersonCreds(Person.UserName)

// 	if Person == ActualPerson {
// 		token, token_err := jwt.CreateToken(Person)
// 		if err != nil || token_err != nil {
// 			ctx.JSON(500, gin.H{"error": "Unable to login"})
// 			ctx.Abort()
// 			return
// 		}
// 		ctx.JSON(200, gin.H{"token": token})
// 	} else {
// 		ctx.JSON(401, gin.H{"error": "Unauthorized"})
// 		ctx.Abort()
// 	}
// }

func AuthorizeUser(ctx *gin.Context) {
	auth_header := ctx.GetHeader("Authorization")
	token := strings.TrimPrefix(auth_header, "Bearer ")
	err := jwt.VerifyJwtAuthToken(token)
	if err != nil {
		ctx.JSON(403, gin.H{"error": "Invalid Token"})
	}
	fmt.Println(token)
}
