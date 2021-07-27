package svc

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/zhuxinlei/micro_user/user/cmd/api/internal/config"
	"github.com/zhuxinlei/micro_user/user/model"
	"github.com/zhuxinlei/micro_user/user/model/table"
)

type ServiceContext struct {
	Config config.Config
	User   model.User
}

func NewServiceContext(c config.Config) *ServiceContext {

	DB, err := gorm.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/micro_user?charset=utf8&parseTime=true&loc=UTC")

	if err != nil {
		logx.Errorf("连接数据库出错%s", err.Error())
	}
	user := table.NewUserModel(DB)
	return &ServiceContext{
		Config: c,
		User:   user,
	}
}
