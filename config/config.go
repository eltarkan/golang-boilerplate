package config

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"os"
)

var (
	basePath, _ = os.Getwd()
	Config      = Environment{}
	DB          *gorm.DB
	RedisClient *redis.Client
)

func LoadConfig() error {
	viper.AddConfigPath(basePath)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		return err
	}
	return nil
}
