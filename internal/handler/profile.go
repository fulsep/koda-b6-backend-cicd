package handler

import (
	"backend/internal/models"
	"backend/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService *services.ProfileService
}

func NewProfileHandler(profileService *services.ProfileService) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
	}
}

func (h *ProfileHandler) GetProfile(ctx *gin.Context) {
	Id := ctx.GetInt("userId")
	user, err := h.profileService.GetProfile(&models.User{
		Id: Id,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Cannot get profile: " + fmt.Sprintf("%v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Profile user",
		"results": user,
	})

}

func (h *ProfileHandler) UpdateProfile(ctx *gin.Context) {
	data := models.UpdateProfileRequest{}

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Validation error",
		})
		return
	}

	user, err := h.profileService.UpdateProfile(&models.User{
		Email:    data.Email,
		Password: data.Password,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Failed to update user: " + fmt.Sprintf("%v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Update profile success",
		"results": user,
	})
}
