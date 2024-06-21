package immodels

import (
	"github.com/wsqigo/easy-chat/pkg/constants"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var DefaultChatLogLimit int64 = 100

// goctl model mongo --type chatLog --dir ./apps/im/immodels
type ChatLog struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`

	ConversationId string             `bson:"conversationId"` // sendId + recvId
	SendId         string             `bson:"sendId"`
	RecvId         string             `bson:"recvId"`
	MsgFrom        int                `bson:"msgFrom"`
	ChatType       constants.ChatType `bson:"chatType"` // 群聊/私聊
	MsgType        constants.MType    `bson:"msgType"`  // 消息类型
	MsgContent     string             `bson:"msgContent"`
	SendTime       int64              `bson:"sendTime"`
	Status         int                `bson:"status"`

	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
