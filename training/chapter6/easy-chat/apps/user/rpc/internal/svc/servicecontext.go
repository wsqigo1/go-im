package svc

import (
	"github.com/wsqigo/easy-chat/apps/user/models"
	"github.com/wsqigo/easy-chat/apps/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	models.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.My)
	return &ServiceContext{
		Config: c,

		UsersModel: models.NewUsersModel(sqlC),
	}
}