package utils

import (
	"daily-brew/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	MemberID       uint   `json:"memberId"`
	Role           string `json:"role"`
	RefreshTokenId string `json:"refreshTokenId"`
	jwt.RegisteredClaims
}

type ClaimsRefresh struct {
	MemberID       uint      `json:"memberId"`
	RefreshTokenId uuid.UUID `json:"refreshTokenId"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(MemberID uint) (string, error) {
	claims := &Claims{
		MemberID: MemberID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Token 1 ngày
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

func GenerateRefreshToken(MemberID uint, refreshTokenId uuid.UUID) (string, error) {
	claims := &ClaimsRefresh{
		MemberID:       MemberID,
		RefreshTokenId: refreshTokenId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)), // Token 7 ngày
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

func VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
