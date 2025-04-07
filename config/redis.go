package config

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", AppConfig.RedisHost, AppConfig.RedisPort),
		Password: AppConfig.RedisPassword, // Nếu có password
		DB:       0,                       // Sử dụng DB mặc định
	})

	// Kiểm tra kết nối
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Không thể kết nối Redis: %v", err)
	}

	log.Println("Đã kết nối Redis thành công!")
}
