package conversation

import (
	"github.com/mitchellh/mapstructure"
	"github.com/wsqigo/easy-chat/apps/im/ws/internal/svc"
	"github.com/wsqigo/easy-chat/apps/im/ws/websocket"
	"github.com/wsqigo/easy-chat/apps/im/ws/ws"
)

func Chat(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		// todo: 私聊
		var data ws.Chat
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err), conn)
			return
		}

	}
}
