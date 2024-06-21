package ws

import "github.com/wsqigo/easy-chat/pkg/constants"

type Msg struct {
	MType constants.MType `mapstructure:"mType"`

	Content string `mapstructure:"content"`
}

type Chat struct {
	ConversationId string             `mapstructure:"conversationId"`
	ChatType       constants.ChatType `mapstructure:"chatType"`
	SendId         string             `mapstructure:"sendId"` // 客户也要知道是谁发送的
	RecvId         string             `mapstructure:"receId"`
	SendTime       int64              `mapstructure:"sendTime"`
	MSg            Msg                `mapstructure:"msg"`
}

type Push struct {
	ConversationId string             `mapstructure:"conversationId"`
	ChatType       constants.ChatType `mapstructure:"chatType"`
	SendId         string             `mapstructure:"sendId"`
	RecvId         string             `mapstructure:"receId"`
	SendTime       int64              `mapstructure:"sendTime"`

	MType   constants.MType `mapstructure:"mType"`
	Content string          `mapstructure:"content"`
}

type MarkRead struct {
	ChatType constants.ChatType `mapstructure:"chatType"`

	RecvId         string   `mapstructure:"receId"`
	ConversationId string   `mapstructure:"conversationId"`
	MsgIds         []string `mapstructure:"msgIds"`
}
