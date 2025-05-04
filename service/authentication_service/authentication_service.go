package authentication_service

import (
	"daily-brew/models"
	"daily-brew/utils"
	"fmt"
	"github.com/google/uuid"
)

func GenerateTokens(memberId uint, role string) (accessToken string, refreshToken string, err error) {
	// Tạo UUID cho refresh token
	refreshTokenId, err := uuid.NewRandom()
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token ID: %w", err)
	}
	// Tạo access token
	accessToken, err = utils.GenerateAccessToken(memberId, refreshTokenId, role)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %w", err)
	}

	// Tạo refresh token
	refreshToken, err = utils.GenerateRefreshToken(memberId, refreshTokenId)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	// Lưu refreshTokenId vào Redis
	err = models.SaveRefreshTokenToRedis(memberId, refreshTokenId.String())
	if err != nil {
		return "", "", fmt.Errorf("failed to save refresh token to Redis: %w", err)
	}

	return accessToken, refreshToken, nil
}
