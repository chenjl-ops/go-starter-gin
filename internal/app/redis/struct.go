package redis

import (
	"github.com/go-redis/redis/v8"
	"go-starter-gin/internal/app/apollo"
)

type Redis struct {
	config      *apollo.Specification
	redisClient *redis.Client
}
