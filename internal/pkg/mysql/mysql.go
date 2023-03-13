package mysql

import _ "github.com/go-sql-driver/mysql"

import (
	"fmt"
	"go-starter-gin/internal/pkg/apollo"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func NewMysql() error {

	result := &DB{
		UserName: apollo.Config.MysqlUserName,
		Password: apollo.Config.MysqlPassword,
		Host:     apollo.Config.MysqlHost,
		Port:     apollo.Config.MysqlPort,
		DBName:   apollo.Config.MysqlDBName,
	}
	err := result.initMysql()
	if err != nil {
		return err
	}
	return nil
}

// 初始化mysql engine/client
func (db *DB) initMysql() error {
	mysqlString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", db.UserName, db.Password, db.Host, db.Port, db.DBName)
	mysqlEngine, err := xorm.NewEngine("mysql", mysqlString)
	if err != nil {
		return err
	}
	Engine = mysqlEngine
	return nil
}
