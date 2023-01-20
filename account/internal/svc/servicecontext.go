package svc

import (
	"fmt"
	"github.com/reation/micro_service_account_service/account/internal/config"
	"github.com/reation/micro_service_account_service/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config       config.Config
	AccountModel model.UserAccountModel
	Conn         sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	dataSourceUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		c.Mysql.AccountTable.User,
		c.Mysql.AccountTable.Passwd,
		c.Mysql.AccountTable.Addr,
		c.Mysql.AccountTable.Port,
		c.Mysql.AccountTable.DBName,
	)
	conn := sqlx.NewMysql(dataSourceUrl)
	AccountConn := model.NewUserAccountModel(conn)
	return &ServiceContext{
		Config:       c,
		AccountModel: AccountConn,
		Conn:         conn,
	}
}
