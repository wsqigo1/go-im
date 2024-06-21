package immodels

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ ChatLogModel = (*customChatLogModel)(nil)

// ChatLogModel is an interface to be customized, add more methods here,
// and implement the added methods in customChatLogModel.
type ChatLogModel interface {
	chatLogModel
}

type customChatLogModel struct {
	*defaultChatLogModel
}

// NewChatLogModel returns a model from the mongo.
func NewChatLogModel(url, db, collection string) ChatLogModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customChatLogModel{
		defaultChatLogModel: newDefaultChatLogModel(conn),
	}
}

func MustChatLogModel(url, db string) ChatLogModel {
	return NewChatLogModel(url, db, "chat_log")
}
