package models

import (
	"daily-brew/config"
	"errors"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"time"
)

// Lưu refresh token vào Redis
func SaveRefreshTokenToRedis(memberId uint, refreshTokenId string) error {
	// Set refresh token với thời gian sống (TTL)
	err := config.RedisClient.Set(config.Ctx, GetKey(memberId), refreshTokenId, 7*24*time.Hour).Err()
	if err != nil {
		log.Printf("Error saving refresh token to Redis: %v", err)
		return err
	}
	return nil
}

// Lấy refresh token từ Redis
func GetRefreshTokenFromRedis(memberId uint) (string, error) {
	token, err := config.RedisClient.Get(config.Ctx, GetKey(memberId)).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil // Không có token
		}
		log.Printf("Error getting refresh token from Redis: %v", err)
		return "", err
	}
	return token, nil
}

// Xóa refresh token khỏi Redis khi logout hoặc hết hạn
func DeleteRefreshTokenFromRedis(memberId uint) error {
	err := config.RedisClient.Del(config.Ctx, GetKey(memberId)).Err()
	if err != nil {
		log.Printf("Error deleting refresh token from Redis: %v", err)
		return err
	}
	return nil
}

func Validate(memberId uint, refreshTokenId string) bool {
	storedId := config.RedisClient.Get(config.Ctx, GetKey(memberId)).Val()
	if storedId != refreshTokenId {
		return false
	}
	return true
}

func GetKey(userId uint) string {
	return "user-" + strconv.Itoa(int(userId))
}
