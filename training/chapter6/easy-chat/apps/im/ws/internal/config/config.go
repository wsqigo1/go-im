package config

import (
	"github.com/zeromicro/go-zero/core/service"
)

type Config struct {
	service.ServiceConf

	ListenOn string

	JwtAuth struct {
		AccessSecret string
	}

	Mongo           MongoConfig
	MsgChatTransfer MsgChatTransferConfig
	MsgReadTransfer MsgReadTransferConfig
}

type MongoConfig struct {
	Url string
	Db  string
}

type MsgChatTransferConfig struct {
	Topic string
	Addrs []string
}

type MsgReadTransferConfig struct {
	Topic string
	Addrs []string
}
