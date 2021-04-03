package server

import (
	"github.com/gin-gonic/gin"
	"go-starter-gin/internal/app/apollo"
	"go-starter-gin/internal/app/test"
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
	apollo.ReadRemoteConfig()
	if nil != err {
		log.Fatal(err)
	}
}

// 加载gin 路由配置
func (s *server) InitRouter() *gin.Engine{
	test.Url(s.App)
	return s.App
}