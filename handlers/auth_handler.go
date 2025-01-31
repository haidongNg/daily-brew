package handlers

import (
	"daily-brew/config"
	"daily-brew/model"
	"daily-brew/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

func Register(c *gin.Context) {
	var reqMember RegisterRequest
	if err := c.ShouldBindJSON(&reqMember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	member := model.Member{
		Email:    reqMember.Email,
		Password: reqMember.Password,
		Name:     reqMember.Name,
		Phone:    reqMember.Phone,
	}
	if err := member.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := member.Save(config.DB); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// find by email
	member, err := model.GetMemberByEmail(config.DB, request.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	if !member.CheckPassword(request.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	accessToken, _ := utils.GenerateAccessToken(int(member.ID))
	refreshToken, _ := utils.GenerateRefreshToken(int(member.ID))
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func RefreshToken(c *gin.Context) {

}
