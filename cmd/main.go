package main

import (
	"backend/internal/di"
	"backend/internal/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	c := di.NewContainer()

	userHandler := c.UserHandler()
	authHandler := c.AuthHandler()
	profileHandler := c.ProfileHandler()

	main := r.Group("/")
	{
		auth := main.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
		}
		profile := main.Group("/profile")
		profile.Use(middleware.Auth(""))
		{
			profile.GET("", profileHandler.GetProfile)
			profile.PATCH("", profileHandler.UpdateProfile)
		}
	}

	admin := r.Group("/admin")
	{
		users := admin.Group("/users")
		{
			users.GET("", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUserById)
			users.POST("", userHandler.CreateUser)
			users.PATCH("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	port := os.Getenv("PORT")
	r.Run(fmt.Sprintf(":%v", port))
}
