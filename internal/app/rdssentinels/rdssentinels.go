package rdssentinels

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-starter-gin/internal/app/apollo"
	"strings"
	"time"
)

var RedisClient *redis.Client
var RedisConfig *Redis

func NewRedis(cfg *Redis) *Redis {
	// 格式化成 数组
	rdsAddress := strings.Split(apollo.Config.RedisSentinelAddress, ",")
	// 每个数组元素 去除首尾空格
	for i, v := range rdsAddress {
		rdsAddress[i] = strings.Trim(v, " ")
	}

	rdb := &Redis{
		RedisSentinelAddress: rdsAddress,
		RedisPasswd:          apollo.Config.RedisPasswd,
		RedisMasterName:      apollo.Config.RedisMasterName,
	}

	// 参数覆盖初始化配置
	if cfg != nil {
		if cfg.RedisMasterName != "" {
			rdb.RedisMasterName = cfg.RedisMasterName
		}
		if cfg.RedisPasswd != "" {
			rdb.RedisPasswd = cfg.RedisPasswd
		}
		if len(cfg.RedisSentinelAddress) != 0 {
			rdb.RedisSentinelAddress = cfg.RedisSentinelAddress
		}
	}
	rdb.RedisClient = rdb.initClient()
	rdb.Ping()
	RedisConfig = rdb
	return rdb
}

func (r *Redis) initClient() *redis.Client {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    r.RedisMasterName,
		SentinelAddrs: r.RedisSentinelAddress,
		Password:      r.RedisPasswd,
	})
	RedisClient = rdb
	return RedisClient
}

func (r *Redis) GetKey(name string) *redis.StringCmd {
	get := r.RedisClient.Get(context.Background(), name)
	return get
}

func (r *Redis) SetKey(name string, v interface{}, h time.Duration) *redis.StatusCmd {
	set := r.RedisClient.Set(context.Background(), name, v, h)
	return set
}

func (r *Redis) XAdd(a *redis.XAddArgs) *redis.StringCmd {
	add := r.RedisClient.XAdd(context.Background(), a)
	return add
}

func (r *Redis) GetStreamsData(streamName string, groupName string, ss string) *redis.XStreamSliceCmd {
	sb := r.RedisClient.XReadStreams(context.Background(), ss)

	// TODO 需要获取所有id ACK一下
	r.RedisClient.XAck(context.Background(), streamName, groupName, ss)
	return sb
}

//测试联通性
func (r *Redis) Ping() {
	_, err := r.RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}
