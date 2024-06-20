package svc

import "github.com/wsqigo/easy-chat/apps/im/ws/internal/config"

type ServiceContext struct {
	Config config.Config

	//immodels.ChatLogModel
	//mqclient.MsgChatTransferClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//MsgChatTransferClient: mqclient.New
	}
}
