package server

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator"
)

type server struct {
	//Config   *apollo.Specification
	App *gin.Engine
	//Validate *validator.Validate
}

func NewServer() (*server, error) {
	return &server{
		//Config: globalConfig,
		App: gin.New(),
	}, nil
}

func StartServer() error {
	initApolloConfig()
	initMysql()
	initRedis()

	server, err1 := NewServer()
	// server.App.Use(Cors())
	if err1 != nil {
		return err1
	}
	// 初始化日志
	server.initLog()
	// 初始化swagger
	server.InitSwagger()
	// 初始化路由
	server.InitRouter()

	//启动服务
	err := server.Run()
	if err != nil {
		return err
	}
	return nil
}

//启动服务
func (s *server) Run() error {
	return s.App.Run(":8080")
}

//func Cors() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.Header("Access-Control-Allow-Origin", "*")
//		c.Next()
//	}
//}
