package ws

import "github.com/wsqigo/easy-chat/pkg/constants"

type Msg struct {
	constants.HandlerResult `mapstructure:"mType"`
	Content                 string `mapstructure:"content"`
}

type Chat struct {
	ConversationId     string `mapstructure:"conversationId"`
	constants.ChatType `mapstructure:"chatType"`
	SendId             string `mapstructure:"sendId"`
	RecvId             string `mapstructure:"receId"`
	SendTime           int64  `mapstructure:"sendTime"`
	Msg                `mapstructure:"msg"`
}

type Push struct {
	ConversationId     string `mapstructure:"conversationId"`
	constants.ChatType `mapstructure:"chatType"`
	SendId             string `mapstructure:"sendId"`
	RecvId             string `mapstructure:"receId"`
	SendTime           int64  `mapstructure:"sendTime"`

	constants.MType `mapstructure:"mType"`
	Content         string `mapstructure:"content"`
}
