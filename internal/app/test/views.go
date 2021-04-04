package test

import (
	"github.com/gin-gonic/gin"
	"go-starter-gin/internal/app/apollo"
	"go-starter-gin/internal/app/mysql"
	"log"
)


func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"RedisSentinels": apollo.Config.RedisSentinels,
		"RedisCluster": apollo.Config.RedisCluster,
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