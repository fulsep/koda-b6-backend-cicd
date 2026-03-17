package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users []User

func corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "content-type,authorization")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusOK)
		} else {
			ctx.Next()
		}
	}
}

func main() {
	r := gin.Default()

	r.Use(corsMiddleware())

	r.POST("/auth/login", func(ctx *gin.Context) {
		var data User
		ctx.ShouldBindJSON(&data)

		if data.Email == "admin@mail.com" &&
			data.Password == "1234" {
			ctx.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "Login success",
			})
		} else {

			ctx.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "Wrong username or password",
			})
		}
	})

	r.GET("/profile", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Backend is running well",
		})
	})

	r.GET("/users", func(ctx *gin.Context) {
		users := Users{
			{
				Id:       1,
				Email:    "admin@mail.com",
				Password: "1234",
			},
			{
				Id:       2,
				Email:    "guest@mail.com",
				Password: "1234",
			},
		}
		ctx.JSON(http.StatusOK, users)
	})

	r.Run(":8888")
}
