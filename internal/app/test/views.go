package test

import (
	"github.com/gin-gonic/gin"
	"go-starter-gin/internal/app/apollo"
	"go-starter-gin/internal/app/mysql"
	"go-starter-gin/internal/app/rdssentinels"
	"log"
	"time"
)


func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"RedisSentinels": apollo.Config.RedisSentinelAddress,
		"RedisCluster": apollo.Config.RedisMasterName,
	})
}

func testSql(c *gin.Context) {
	ticketData := make([]RobotTicket, 0)

	if err := mysql.Engine.Find(&ticketData); err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"ticket_data": ticketData,
	})
}

func testRedis(c *gin.Context) {
	rds := rdssentinels.NewRedis(nil)
	rds.SetKey("testKey", "this is test demo", 3600 * time.Second)
	result := rds.GetKey("testKey")

	c.JSON(200, gin.H{
		"redis_testKey": result.Val(),
	})
}