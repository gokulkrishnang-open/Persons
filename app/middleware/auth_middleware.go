package middleware

import (
	"persons/app/middleware/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthorizeUser(ctx *gin.Context) {
	auth_header := ctx.GetHeader("Authorization")
	token := strings.TrimPrefix(auth_header, "Bearer ")
	err := jwt.ValidateJwtAuthToken(token)
	if err != nil {
		ctx.JSON(401, gin.H{"error": "Unauthorized"})
		ctx.Abort()
	}
	ctx.Next()
}
