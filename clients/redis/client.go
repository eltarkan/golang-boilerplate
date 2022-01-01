package redis

import "github.com/go-redis/redis/v8"

func CreateRedisClient(client *redis.Client, address string) *redis.Client {
	client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}
