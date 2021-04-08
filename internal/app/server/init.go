package server

import (
	"github.com/gin-gonic/gin"
	"go-starter-gin/internal/app/apollo"
	"go-starter-gin/internal/app/mysql"
	"go-starter-gin/internal/app/rdssentinels"
	"go-starter-gin/internal/app/test"
	"go-starter-gin/internal/app/middleware/logger"
	"log"
)

/*
TODO
1、日志初始化
2、通用数据库mysql封装
3、通用Redis封装
4、Event跨实例间通讯
*/

// 初始化 apollo config
func initApolloConfig() {
	var err error
	err = apollo.ReadRemoteConfig()
	if nil != err {
		log.Fatal(err)
	}
}

// 初始化mysql
func initMysql() {
	err := mysql.NewMysql()
	if nil != err {
		log.Fatal(err)
	}
}

// 初始化redis
func initRedis() {
	rdssentinels.NewRedis(nil)
}

// 初始化log配置
func (s *server) initLog() *gin.Engine {
	logs := logger.LogMiddleware()
	s.App.Use(logs)
	return s.App
}

// 加载gin 路由配置
func (s *server) InitRouter() *gin.Engine{
	test.Url(s.App)
	return s.App
}