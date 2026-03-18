package middleware

import (
	"backend/internal/lib"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authStr := ctx.GetHeader("Authorization")
		token, found := strings.CutPrefix(authStr, "Bearer ")
		valid, payload := lib.VerifyToken(token)
		if !found {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			return
		}

		if !valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			return
		}

		userId := payload.(lib.CustomClaims).UserId
		ctx.Set("userId", userId)
		ctx.Next()
	}
}
