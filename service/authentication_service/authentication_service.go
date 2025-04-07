package authentication_service

import (
	"daily-brew/models"
	"daily-brew/utils"
	"github.com/google/uuid"
)

func GenerateTokens(memberId uint) (string, string, error) {
	refreshTokenId, _ := uuid.NewRandom()
	accessToken, _ := utils.GenerateAccessToken(memberId)
	refreshToken, _ := utils.GenerateRefreshToken(memberId, refreshTokenId)

	err := models.SaveRefreshTokenToRedis(memberId, refreshTokenId.String())

	return accessToken, refreshToken, err
}
