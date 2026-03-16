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

func main() {
	r := gin.Default()

	r.GET("/auth/login", func(ctx *gin.Context) {
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
