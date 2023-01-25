package middleware

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"persons/config"
	"persons/model"
	"strings"
)

func AuthenticateUser(ctx *gin.Context) {
	encoded_things := strings.SplitN(ctx.GetHeader("Authorization"), " ", 2)[1]
	b, _ := base64.StdEncoding.DecodeString(encoded_things)
	credentials := strings.SplitN(string(b), ":", 2)
	Person := model.PersonRequest{credentials[0], credentials[1]}

	if Person == *config.AuthorizedPerson {
		ctx.Next()
	} else {
		ctx.JSON(403, gin.H{"error": "Unauthorized"})
		ctx.Abort()
	}
}
