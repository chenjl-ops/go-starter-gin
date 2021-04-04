package server

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-playground/validator"
)

type server struct {
	//Config   *apollo.Specification
	App      *gin.Engine
	//Validate *validator.Validate
}

func NewServer() (*server, error){
	return &server{
		//Config: globalConfig,
		App: gin.New(),
	}, nil
}


func StartServer() error {
	initApolloConfig()
	initMysql()

	server, err1 := NewServer()
	if err1 != nil {
		return err1
	}
	err := server.InitRouter().Run()
	if err != nil {
		return err
	}
	return nil
}

//启动服务
func (s *server) Run() error {
	return s.App.Run(":8080")
}
