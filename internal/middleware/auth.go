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
		if !found || !lib.VerifyToken(token) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			return
		}
		ctx.Next()
	}
}
