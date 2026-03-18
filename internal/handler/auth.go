package handler

import (
	"backend/internal/lib"
	"backend/internal/models"
	"backend/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var data models.LoginRequest

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Validation error " + fmt.Sprintf("%v", err.Error()),
		})
		return
	}
	user, err := h.service.Login(&models.User{
		Email:    data.Email,
		Password: data.Password,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Bad request: " + fmt.Sprintf("%v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Login success",
		"results": gin.H{
			"token": lib.GenerateToken(user.Id),
			"user":  user,
		},
	})
}
