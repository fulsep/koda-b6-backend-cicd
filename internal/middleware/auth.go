package middleware

import (
	"backend/internal/lib"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authStr := ctx.GetHeader("Authorization")
		token, found := strings.CutPrefix(authStr, "Bearer ")
		valid, payload := lib.VerifyToken(token)
		if !found || !valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			return
		}
		fmt.Println(payload)
		ctx.Set("userId", 1)
		ctx.Next()
	}
}
