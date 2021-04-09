package test

import (
	"github.com/gin-gonic/gin"
	"go-starter-gin/internal/app/apollo"
	"go-starter-gin/internal/app/mysql"
	"go-starter-gin/internal/app/rdssentinels"
	"log"
	"time"
)

// @Tag Test API
// @Summary List apollo some config
// @Description get apollo config
// @Accept  json
// @Produce  json
// @Success 200 {array} Response
// @Header 200 {string} Response
// @Failure 400,404 {object} string "Bad Request"
// @Router /v1/test1 [get]
func test(c *gin.Context) {
	c.JSON(200, Response{
		Code: 200,
		Data: map[string]string{
			"RedisSentinels": apollo.Config.RedisSentinelAddress,
			"RedisCluster":   apollo.Config.RedisMasterName,
		},
		Msg: "success",
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
	//rds := rdssentinels.NewRedis(nil)
	//rds.SetKey("testKey", "this is test demo", 3600 * time.Second)
	//result := rds.GetKey("testKey")

	rdssentinels.RedisConfig.SetKey("testKey", "this is test demo", 3600 * time.Second)
	result := rdssentinels.RedisConfig.GetKey("testKey")

	c.JSON(200, gin.H{
		"redis_testKey": result.Val(),
	})
}