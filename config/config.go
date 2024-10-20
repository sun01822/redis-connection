package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type RedisConfig struct {
	Host string 
	Port string 
	Password string 
	DB int 
}

type Config struct{
	Redis *RedisConfig
}

var config Config

func Redis() *RedisConfig{
	return config.Redis
}

func LoadConfig(){
	err := godotenv.Load()
	if err!= nil {
        log.Printf("Error loading .env file: %v", err)
    }

	viper.AutomaticEnv()

	config.Redis = &RedisConfig{
		Host: viper.GetString("HOST"),
        Port: viper.GetString("PORT"),
        Password: viper.GetString("PASSWORD"),
        DB: viper.GetInt("DB"),
	}
}