package test

import (
	"github.com/gin-gonic/gin"
	"go-starter-gin/internal/app/apollo"
)


func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"RedisSentinels": apollo.Config.RedisSentinels,
		"RedisCluster": apollo.Config.RedisCluster,
	})
}
