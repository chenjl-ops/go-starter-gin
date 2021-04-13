package rdssentinels

import "github.com/go-redis/redis/v8"

//type Redis struct {
//	config      *apollo.Specification
//	redisClient *redis.Client
//}

type Redis struct {
	RedisMasterName      string
	RedisSentinelAddress []string
	RedisPasswd          string
	RedisClient          *redis.Client
}
