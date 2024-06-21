package mq

import "github.com/wsqigo/easy-chat/pkg/constants"

type MsgChatTransfer struct {
	ConversationId string             `json:"conversationId"`
	ChatType       constants.ChatType `json:"chatType"`
	SendId         string             `json:"sendId"`
	RecvId         string             `json:"recvId"`
	SendTime       int64              `json:"sendTime"`

	MType   constants.MType `json:"MType"`
	Content string          `json:"content"`
}

type MsgMarkRead struct {
	ChatType       constants.ChatType `json:"chatType"`
	ConversationId string             `json:"conversationId"`
	SendId         string             `json:"sendId"`
	RecvId         string             `json:"recvId"`
	MsgIds         []string           `json:"msgIds"`
}
