package controllers

import (
	"daily-brew/models"
	"daily-brew/service/authentication_service"
	"daily-brew/service/member_service"
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
	var form RegisterRequest
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	memberService := member_service.Member{
		Email:    form.Email,
		Password: form.Password,
		Name:     form.Name,
		Phone:    form.Phone,
	}

	if err := memberService.Register(); err != nil {
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

	memberService := member_service.Member{
		Email: request.Email,
	}
	// find by email
	member, err := memberService.GetMemberByEmail()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}

	if !utils.BcryptCheck(request.Password, member.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
		return
	}
	var (
		accessToken, refreshToken string
	)
	accessToken, refreshToken, err = authentication_service.GenerateTokens(member.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid credentials"})
	}
	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

func RefreshToken(c *gin.Context) {
	var request RefreshTokenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Xác thực refresh token
	claims, err := utils.VerifyToken(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	memberService := member_service.Member{
		ID: uint(claims.MemberID),
	}
	member, err := memberService.GetMemberByID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid credentials"})
		return
	}

	isValid := models.Validate(member.ID, claims.RefreshTokenId)
	if !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token is invalid"})
		return
	}

	err = models.DeleteRefreshTokenFromRedis(member.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid credentials"})
		return
	}

	var (
		accessToken, refreshToken string
	)
	accessToken, refreshToken, err = authentication_service.GenerateTokens(member.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid credentials"})
	}
	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}
