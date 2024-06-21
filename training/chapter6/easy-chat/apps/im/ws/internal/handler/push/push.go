package push

import (
	"github.com/mitchellh/mapstructure"
	"github.com/wsqigo/easy-chat/apps/im/ws/internal/svc"
	"github.com/wsqigo/easy-chat/apps/im/ws/websocket"
	"github.com/wsqigo/easy-chat/apps/im/ws/ws"
	"github.com/wsqigo/easy-chat/pkg/constants"
)

func Push(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		var data ws.Push
		if err := mapstructure.Decode(msg.Data, &data); err != nil {
			srv.Send(websocket.NewErrMessage(err))
			return
		}

		// 发送的目标
		switch data.ChatType {
		case constants.SingleChatType:
			single(srv, &data, data.RecvId)
		case constants.GroupChatType:
			group(srv, &data)
		}

	}
}

func group(server *websocket.Server, w *ws.Push) {

}

func single(srv *websocket.Server, w *ws.Push, id string) {

}
