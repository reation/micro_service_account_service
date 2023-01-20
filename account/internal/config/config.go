package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		AccountTable struct {
			User   string
			Passwd string
			Addr   string
			DBName string
			Port   string
		}
	}
}
