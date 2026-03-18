package main

import (
	"backend/internal/di"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	c := di.NewContainer()

	userHandler := c.UserHandler()

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
