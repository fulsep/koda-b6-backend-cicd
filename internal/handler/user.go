package handler

import (
	"backend/internal/models"
	"backend/internal/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var data models.CreateUserRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": "Failed to validate",
			},
		)
	}

	user, err := h.service.CreateUser(&models.User{
		Email:    data.Email,
		Password: data.Password,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to create user: " + fmt.Sprintf("%v", err.Error()),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Create user success",
		"results": user,
	})

}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": "ID format not vaild",
			},
		)
	}

	var data models.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": "Failed to validate",
			},
		)
	}

	user, err := h.service.UpdateUser(&models.User{
		Id:       idInt,
		Email:    data.Email,
		Password: data.Password,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to update user: " + fmt.Sprintf("%v", err.Error()),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Update user success",
		"results": user,
	})

}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": "ID format not valid",
			},
		)
	}

	user, err := h.service.DeleteUser(&models.User{
		Id: idInt,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to delete user: " + fmt.Sprintf("%v", err.Error()),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Delete user success",
		"results": user,
	})

}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {

	user, err := h.service.GetAllUsers()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to get all users: " + fmt.Sprintf("%v", err.Error()),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "List all users",
		"results": user,
	})

}

func (h *UserHandler) GetUserById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"success": false,
				"message": "ID format not valid",
			},
		)
	}

	user, err := h.service.GetUserById(&models.User{
		Id: idInt,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to get user: " + fmt.Sprintf("%v", err.Error()),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Detail user",
		"results": user,
	})

}
