package main

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"log"
	rdc "main/clients/redis"
	"main/config"
	"main/database"
	"os"
)

func main() {
	var err error
	if config.LoadConfig(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var DB *gorm.DB
	config.DB, err = database.ConnectDB(DB, config.Config.DBName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	
	var redisClient *redis.Client
	config.RedisClient = rdc.CreateRedisClient(redisClient, config.Config.RedisAddr)

	InitRouting()
}
