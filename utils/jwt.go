package utils

import (
	"daily-brew/config"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type Claims struct {
	MemberID       uint      `json:"memberId"`
	Role           string    `json:"role"`
	RefreshTokenId uuid.UUID `json:"refreshTokenId"`
	jwt.RegisteredClaims
}

type ClaimsRefresh struct {
	MemberID       uint      `json:"memberId"`
	RefreshTokenId uuid.UUID `json:"refreshTokenId"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(MemberID uint, refreshTokenId uuid.UUID, role string) (string, error) {
	claims := &Claims{
		MemberID:       MemberID,
		Role:           role,
		RefreshTokenId: refreshTokenId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Token 1 ngày
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    config.AppConfig.JWTIssuer,
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
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    config.AppConfig.JWTIssuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

// VerifyAccessToken kiểm tra và parse Access Token
func VerifyAccessToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("access token parse error: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid access token")
	}

	if claims.Issuer != config.AppConfig.JWTIssuer {
		return nil, errors.New("invalid token issuer")
	}

	return claims, nil
}

// VerifyRefreshToken kiểm tra và parse Refresh Token
func VerifyRefreshToken(tokenString string) (*ClaimsRefresh, error) {
	token, err := jwt.ParseWithClaims(tokenString, &ClaimsRefresh{}, keyFunc)
	if err != nil {
		return nil, fmt.Errorf("refresh token parse error: %w", err)
	}

	claims, ok := token.Claims.(*ClaimsRefresh)
	if !ok || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	if claims.Issuer != config.AppConfig.JWTIssuer {
		return nil, errors.New("invalid token issuer")
	}

	return claims, nil
}

// keyFunc dùng chung để kiểm tra thuật toán và trả về key
func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(config.AppConfig.JWTSecret), nil
}
