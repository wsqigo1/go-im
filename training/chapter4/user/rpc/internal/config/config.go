package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Cache cache.CacheConf
	Mysql MysqlConfig
}

type MysqlConfig struct {
	DataSource string
}
