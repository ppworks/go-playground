package database

import (
	"os"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

func init() {
	redisURL := os.Getenv("REDIS_URL")
	opt, _ := redis.ParseURL(redisURL)
	client = redis.NewClient(opt)
}

// Redis return cofigured *redis.Client
func Redis() *redis.Client {
	return client
}
