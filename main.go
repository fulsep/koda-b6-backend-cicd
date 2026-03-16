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
