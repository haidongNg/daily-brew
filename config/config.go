package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
	ServerPort string `mapstructure:"SERVER_PORT"`
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv() // Auto get value

	if err := viper.ReadInConfig(); err != nil {
		log.Println("File config.yaml could not be loaded")
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
	}

	log.Println("Successfully loaded config")
}
