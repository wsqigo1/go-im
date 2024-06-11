package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql MysqlConfig
	Cache cache.CacheConf
	Jwt   Jwt
}

type MysqlConfig struct {
	DataSource string
}

type Jwt struct {
	AccessSecret string
	AccessExpire int64
}
